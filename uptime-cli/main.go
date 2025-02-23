package main

import (
	"github.com/mskreczko/uptime-checker/internal"
	"sync"
)

func main() {
	config := internal.ReadConfig("./config.yaml")
	notificationService := internal.NewNotificationService(config.SMTPSettings)

	var wg sync.WaitGroup

	for _, application := range config.Applications {
		for _, targetGroup := range application.TargetGroups {
			wg.Add(1)
			go func() {
				job := internal.CreateHealthCheckJob(targetGroup.Targets, targetGroup.HealthcheckInterval)
				defer wg.Done()
				job.Run()
			}()
		}
	}

	wg.Wait()
}
