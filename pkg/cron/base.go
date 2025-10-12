package cron

import (
	"fmt"
	"os"

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

	if os.Getenv("APP_ENV") == "development" {
		if _, err = s.NewJob(
			gocron.CronJob("*/10 * * * *", false),
			gocron.NewTask(
				func() {
					fmt.Println("running dollar debug cron")
					injector.InitDollarCron().StartDebug()
				},
			),
		); err != nil {
			fmt.Println(err)
		}
	} else if os.Getenv("APP_ENV") == "production" {
		if _, err = s.NewJob(
			gocron.CronJob("*/45 * * * *", false),
			gocron.NewTask(
				func() {
					fmt.Println("running dollar prod cron")
					injector.InitDollarCron().StartProd()
				},
			),
		); err != nil {
			fmt.Println(err)
		}
	}

	s.Start()
}
