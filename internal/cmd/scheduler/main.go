package scheduler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	cinf "restuwahyu13/shopping-cart/internal/domain/interface/common"
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/pkg"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"
	repo "restuwahyu13/shopping-cart/internal/infrastructure/repository"
	"time"

	"github.com/uptrace/bun"
)

type Scheduler struct {
	ENV *cdto.Environtment
	DB  *bun.DB
	RDS pinf.IRedis
}

func NewScheduler(options Scheduler) cinf.IScheduler {
	return Scheduler{
		ENV: options.ENV,
		DB:  options.DB,
		RDS: options.RDS,
	}
}

func (p Scheduler) ExecuteUpdateOrderStatus() {
	ctx := context.Background()
	cron := pkg.NewCron()

	key := "CALLBACK_SIMULATOR"
	now := time.Now().Format("2006-01-02 15:04:05")
	crontime := cons.Every15Seconds

	s, _, err := cron.Handler("update_order_status", crontime, func() {
		pkg.Logrus(cons.INFO, fmt.Sprintf("Cron is running %s - and execute at %s", now, crontime))

		resultCallbackSimulator, err := p.RDS.HGetAll(key)
		if err != nil {
			pkg.Logrus(cons.ERROR, err)
			return
		}

		if resultCallbackSimulator != nil {
			callbackSimulatorDTO := sdto.CallbackSimulatorPaymentDTO{}

			for k, _ := range resultCallbackSimulator {
				result := resultCallbackSimulator[k]

				if err := json.Unmarshal([]byte(result), &callbackSimulatorDTO); err != nil {
					pkg.Logrus(cons.ERROR, err)
					return
				}
			}

			// NOTE: CHECK IDEMPOTENCY KEY IS EXIST OR NOT

			paymentModel := new(model.PaymentModel)
			paymentRepository := repo.NewPaymentRepository(context.Background(), p.DB)

			err = paymentRepository.FindOne().Column("id", "request_id", "expired_at").
				Where("deleted_at IS NULL AND verified_at IS NOT NULL AND request_id = ? AND status = ?", callbackSimulatorDTO.IdempotencyKey, cons.PENDING).
				Scan(context.Background(), paymentModel)

			if err != nil && err != sql.ErrNoRows {
				pkg.Logrus(cons.ERROR, err)
				return

			} else if err == sql.ErrNoRows {
				pkg.Logrus(cons.ERROR, err)
				return
			}

			// NOTE: CHECK PAYMENT IS EXPIRED OR NOT

			now := time.Now()
			expired := paymentModel.ExpiredAt.Time.Compare(now)

			if expired == -1 {
				pkg.Logrus(cons.ERROR, fmt.Sprintf("Payment transaction expired for this payment code %s", paymentModel.RequestID))
				callbackSimulatorDTO.Status = cons.EXPIRED

				return
			}

			if paymentModel.Status != cons.PENDING && callbackSimulatorDTO.IdempotencyKey != "" {
				// NOTE: UPDATE PAYMENT STATUS TO SUCCEED | FAILED

				resultPaymentUpdate, err := p.DB.NewUpdate().Table("payment").Where("id = ?", paymentModel.ID).
					Set("status = ?", callbackSimulatorDTO.Status, "updated_at = ?", time.Now()).Exec(ctx)

				if err != nil {
					pkg.Logrus(cons.ERROR, err)
					return

				} else if rows, err := resultPaymentUpdate.RowsAffected(); rows < 1 || err != nil {
					if err == nil {
						pkg.Logrus(cons.ERROR, "Failed to update payment status")
						return
					}

					pkg.Logrus(cons.ERROR, err)
					return
				}

				// NOTE: UPDATE ORDER STATUS TO PAID

				if callbackSimulatorDTO.Status == cons.SUCCEED {
					resultOrderUpdate, err := p.DB.NewUpdate().Table("order").Column("status").
						Where("deleted_at IS NULL AND paid = false AND payment_id = ? AND status = ?", paymentModel.ID, cons.WAITING).
						Set("status = ?", cons.PAID, "paid = ?", true, "updated_at = ?", time.Now()).Exec(ctx)

					if err != nil {
						pkg.Logrus(cons.ERROR, err)
						return

					} else if rows, err := resultOrderUpdate.RowsAffected(); rows < 1 || err != nil {
						if err == nil {
							pkg.Logrus(cons.ERROR, "Failed to update order status")
							return
						}

						pkg.Logrus(cons.ERROR, err)
						return
					}

					if _, err := p.RDS.HDel(key, callbackSimulatorDTO.IdempotencyKey); err != nil {
						pkg.Logrus(cons.ERROR, err)
						return
					}

					pkg.Logrus(cons.INFO, fmt.Sprintf("Deleted payment idempotency key %s", callbackSimulatorDTO.IdempotencyKey))
					return
				}
			}
		}
	})

	if err != nil {
		pkg.Logrus(cons.ERROR, err)
	}

	s.Start()
}
