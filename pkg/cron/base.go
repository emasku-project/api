package cron

import (
	"fmt"

	"api/internal/injector"
	"github.com/go-co-op/gocron/v2"
)

func New() {
	fmt.Println("cron started")
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println(err)
	}

	if _, err = s.NewJob(
		gocron.CronJob("*/10 * * * *", false),
		gocron.NewTask(
			func() {
				fmt.Println("running gold cron")
				injector.InitGoldCron().Start()
			},
		),
	); err != nil {
		fmt.Println(err)
	}

	if _, err = s.NewJob(
		gocron.CronJob("*/10 * * * *", false),
		gocron.NewTask(
			func() {
				fmt.Println("running dollar cron")
				injector.InitDollarCron().Start()
			},
		),
	); err != nil {
		fmt.Println(err)
	}

	s.Start()
}
