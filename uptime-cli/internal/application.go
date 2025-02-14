package internal

type Application struct {
	Name         string        `yaml:"name"`
	TargetGroups []TargetGroup `yaml:"targetGroups"`
}

type TargetGroup struct {
	Name    string   `yaml:"name"`
	Targets []string `yaml:"targets"`
}

// type Target struct {
// 	HealtcheckUrl string
// }
