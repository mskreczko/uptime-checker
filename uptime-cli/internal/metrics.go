package internal

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TargetHealthMetric = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "target_health",
			Help: "Current state of target (0 = down, 1 = up)",
		},
		[]string{})
)
