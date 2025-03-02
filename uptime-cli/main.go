package main

import (
	"fmt"
	"github.com/mskreczko/uptime-checker/internal"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
	"time"
)

func main() {
	config := internal.ReadConfig("/config/config.yaml")
	notificationService := internal.NewNotificationService(config.SMTPSettings)

	var healthcheckService internal.HealthcheckService

	var wg sync.WaitGroup

	for _, application := range config.Applications {
		for _, targetGroup := range application.TargetGroups {
			wg.Add(1)
			go func() {
				job := healthcheckService.CreateHealthCheckJob(targetGroup.Targets, targetGroup.HealthcheckInterval, targetGroup.HealthCheckStrategy)
				defer wg.Done()
				job.Run()
			}()
		}
	}

	for _, job := range healthcheckService.HealthcheckJobs {
		for _, healthcheck := range job.Healthchecks {
			if time.Now().Sub(healthcheck.LastUp).Seconds() > float64(job.Interval) {
				notificationService.SendServicesDownNotification(healthcheck)
			}
		}
	}

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.ListeningPort), nil)
	if err != nil {
		return
	}

	wg.Wait()
}
