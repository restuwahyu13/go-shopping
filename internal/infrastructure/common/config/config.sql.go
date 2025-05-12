package cfg

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"
)

func Database(env *cdto.Environtment) (*bun.DB, error) {
	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(env.POSTGRES.URL)))

	if err := db.Ping(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	if db != nil {
		logrus.Info("Database connection success")

		db.SetConnMaxIdleTime(time.Duration(time.Second * time.Duration(30)))
		db.SetConnMaxLifetime(time.Duration(time.Second * time.Duration(30)))
	}

	bundb := bun.NewDB(db, pgdialect.New())

	if env.APP.ENV != cons.PROD {
		bundb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithEnabled(true), bundebug.WithVerbose(true), bundebug.FromEnv("BUNDEBUG")))
	}

	return bundb, nil
}
