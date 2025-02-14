package internal

import (
	"net/url"
)

type Application struct {
	Name string `yaml:"name"`
	TargetGroups []TargetGroup `yaml:"targetGroups"`
}

type TargetGroup struct {
	Name string `yaml:"name"`
	Targets Target[] `yaml:"targets"`
}

type Target struct {
	Name string `yaml:"name"`
	HealtcheckUrl url
}

