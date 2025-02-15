package main

import (
	"github.com/mskreczko/uptime-checker/internal"
	"github.com/mskreczko/uptime-checker/pkg"
	"sync"
)

func main() {
	config := internal.ReadConfig("./config.yaml")

	var wg sync.WaitGroup

	for _, application := range config.Applications {
		for _, targetGroup := range application.TargetGroups {
			wg.Add(1)
			go func() {
				job := pkg.CreateHealthCheckJob(targetGroup.Targets, targetGroup.HealthcheckInterval)
				defer wg.Done()
				job.Run()
			}()
		}
	}

	wg.Wait()
}
