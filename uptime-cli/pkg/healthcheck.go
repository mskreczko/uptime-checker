package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type TargetGroupHealthcheckJob struct {
	Healthchecks []Healthcheck
	Interval     int
}

type Healthcheck struct {
	url    url.URL
	lastUp time.Time
}

func CreateHealthCheckJob(urls []url.URL, interval int) TargetGroupHealthcheckJob {
	var healthchecks []Healthcheck

	for _, _url := range urls {
		healthchecks = append(healthchecks, Healthcheck{_url, time.Now()})
	}

	return TargetGroupHealthcheckJob{healthchecks, interval}
}

func (job TargetGroupHealthcheckJob) Run() {
	for range time.Tick(time.Duration(job.Interval) * time.Second) {
		for _, healthcheck := range job.Healthchecks {
			if makeRequest(healthcheck.url) {
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
	body, err := ioutil.ReadAll(res.Body)
	if string(body) == "{\"status\": \"UP\"" {
		return true
	}
	return false
}
