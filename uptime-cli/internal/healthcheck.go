package internal

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type TargetGroupHealthcheckJob struct {
	Healthchecks []Healthcheck
	Interval     int
}

type Healthcheck struct {
	url    YamlURL
	lastUp time.Time
}

func CreateHealthCheckJob(urls []YamlURL, interval int) TargetGroupHealthcheckJob {
	var healthchecks []Healthcheck

	for _, _url := range urls {
		healthchecks = append(healthchecks, Healthcheck{_url, time.Now()})
	}

	return TargetGroupHealthcheckJob{healthchecks, interval}
}

func (job *TargetGroupHealthcheckJob) Run() {
	ticker := time.NewTicker(time.Duration(job.Interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for _, healthcheck := range job.Healthchecks {
			if makeRequest(*healthcheck.url.URL) {
				healthcheck.lastUp = time.Now()
			}
		}
	}
}

func makeRequest(url url.URL) bool {
	res, err := http.Get(url.String())
	if err != nil {
		fmt.Printf("Error making request to %s: %s\n", url.String(), err)
		return false
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if string(body) == "{\"status\": \"UP\"}" {
		return true
	}
	return false
}
