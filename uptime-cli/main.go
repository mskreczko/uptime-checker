package main

import (
	"github.com/mskreczko/uptime-checker/internal"
	"github.com/mskreczko/uptime-checker/pkg"
	"sync"
	"time"
)

func main() {
	config := internal.ReadConfig("./config.yaml")
	notificationService := internal.NewNotificationService(config.SMTPSettings)

	var healthcheckService internal.HealthcheckService

	var wg sync.WaitGroup

	for _, application := range config.Applications {
		for _, targetGroup := range application.TargetGroups {
			wg.Add(1)
			go func() {
				job := healthcheckService.CreateHealthCheckJob(targetGroup.Targets, targetGroup.HealthcheckInterval)
				defer wg.Done()
				job.Run()
			}()
		}
	}

	for _, job := range healthcheckService.HealthcheckJobs {
		for _, healthcheck := range job.Healthchecks {
			if time.Now().Sub(healthcheck.LastUp).Seconds() > float64(job.Interval) {
				// TODO
				// Extract it to notification service as more generic function
				notificationService.SendNotifications(pkg.EmailRequest{
					To:      config.NotificationSettings.SettingEntries[0].Receivers[0],
					Subject: "One of your services is down",
					Body:    "One of your services has not responded for healthcheck, last up: " + healthcheck.LastUp.String(),
				})
			}
		}
	}

	wg.Wait()
}
