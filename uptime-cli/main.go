package main

import (
	"github.com/mskreczko/uptime-checker/internal"
	"github.com/mskreczko/uptime-checker/pkg"
)

func main() {
	config := internal.ReadConfig("./config.yaml")

	for _, application := range config.Applications {
		for _, targetGroup := range application.TargetGroups {
			job := pkg.CreateHealthCheckJob(targetGroup.Targets, targetGroup.HealthcheckInterval)
		}
	}
}
