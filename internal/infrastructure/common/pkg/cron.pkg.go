package pkg

import (
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"

	"github.com/go-co-op/gocron/v2"
)

type cron struct{}

func NewCron() pinf.ICron {
	return cron{}
}

func (p cron) Handler(name, crontime string, task func()) (gocron.Scheduler, gocron.Job, error) {
	scheduler, err := gocron.NewScheduler()

	if err != nil {
		return nil, nil, err
	}

	job, err := scheduler.NewJob(gocron.CronJob(crontime, true), gocron.NewTask(task), gocron.WithName(name))
	if err != nil {
		return nil, nil, err
	}

	return scheduler, job, nil
}
