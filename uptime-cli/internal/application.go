package internal

import (
	"gopkg.in/yaml.v3"
	"net/url"
)

type Application struct {
	Name         string        `yaml:"name"`
	TargetGroups []TargetGroup `yaml:"targetGroups"`
}

type TargetGroup struct {
	Name                string    `yaml:"name"`
	Targets             []YamlURL `yaml:"targets"`
	HealthcheckInterval int       `yaml:"healthCheckInterval"`
	HealthCheckStrategy string    `yaml:"healthCheckStrategy"`
}

type YamlURL struct {
	URL *url.URL
}

func (v *YamlURL) UnmarshalYAML(value *yaml.Node) error {
	var strValue string
	if err := value.Decode(&strValue); err != nil {
		return err
	}

	_url, err := url.Parse(strValue)
	v.URL = _url
	return err
}
