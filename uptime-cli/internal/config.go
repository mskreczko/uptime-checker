package internal

type Config struct {
	Applications []Application `yaml:"applications"`	
	NotificationSettings NotificationSettings `yaml:"notifications"`
}