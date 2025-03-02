package internal

import (
	"time"
)

type TargetGroupHealthcheckJob struct {
	Healthchecks        []Healthcheck
	Interval            int
	HealthCheckStrategy HealthCheckStrategy
}

type Healthcheck struct {
	Url    YamlURL
	LastUp time.Time
}

type HealthcheckService struct {
	HealthcheckJobs []TargetGroupHealthcheckJob
}

func (hs *HealthcheckService) CreateHealthCheckJob(urls []YamlURL, interval int, healthCheckStrategy string) TargetGroupHealthcheckJob {
	var healthchecks []Healthcheck

	for _, _url := range urls {
		healthchecks = append(healthchecks, Healthcheck{_url, time.Now()})
	}

	job := TargetGroupHealthcheckJob{healthchecks, interval, ResolveHealthCheckStrategy(healthCheckStrategy)}
	hs.HealthcheckJobs = append(hs.HealthcheckJobs, job)
	return job
}

func (job *TargetGroupHealthcheckJob) Run() {
	ticker := time.NewTicker(time.Duration(job.Interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for _, healthcheck := range job.Healthchecks {
			if job.HealthCheckStrategy.healthcheck(*healthcheck.Url.URL) {
				healthcheck.LastUp = time.Now()
				TargetHealthMetric.WithLabelValues(healthcheck.Url.URL.String()).Set(1.0)
			} else {
				TargetHealthMetric.WithLabelValues(healthcheck.Url.URL.String()).Set(0.0)
			}
		}
	}
}
