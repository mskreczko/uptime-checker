package internal

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HealthCheckStrategy interface {
	healthcheck(url url.URL) bool
}

type JSONHealthCheckStrategy struct{}

func (s *JSONHealthCheckStrategy) healthcheck(url url.URL) bool {
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

type RawStringHealthCheckStrategy struct{}

func (s *RawStringHealthCheckStrategy) healthcheck(url url.URL) bool {
	res, err := http.Get(url.String())
	if err != nil {
		fmt.Printf("Error making request to %s: %s\n", url.String(), err)
		return false
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	return strings.Contains(string(body), "UP")
}

type StatusHealthCheckStrategy struct{}

func (s *StatusHealthCheckStrategy) healthcheck(url url.URL) bool {
	res, err := http.Get(url.String())
	if err != nil {
		fmt.Printf("Error making request to %s: %s\n", url.String(), err)
		return false
	}
	return res.StatusCode == 200
}

func ResolveHealthCheckStrategy(strategy string) HealthCheckStrategy {
	switch strategy {
	case "json":
		return &JSONHealthCheckStrategy{}
	case "raw":
		return &RawStringHealthCheckStrategy{}
	case "status_code":
		return &StatusHealthCheckStrategy{}
	}
	return nil
}
