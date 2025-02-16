package internal

import (
	"github.com/mskreczko/uptime-checker/pkg"
	"gopkg.in/yaml.v3"
)

type NotificationChannel string

const (
	SMS   NotificationChannel = "SMS"
	EMAIL                     = "EMAIL"
)

type NotificationSettings struct {
	SettingEntries []NotificationSettingEntry `yaml:"notification_settings"`
}

type NotificationSettingEntry struct {
	Channel   NotificationChannel `yaml:"channel"`
	Receivers []string            `yaml:"receivers"`
}

func (v *NotificationChannel) UnmarshalYAML(value *yaml.Node) error {
	var strValue string
	if err := value.Decode(&strValue); err != nil {
		return err
	}

	switch strValue {
	case "EMAIL":
		*v = EMAIL
	case "SMS":
		*v = SMS
	}
	return nil
}

type NotificationService struct {
	smtpClient pkg.SMTPClient
}

func NewNotificationService(smtpClient pkg.SMTPClient) *NotificationService {
	return &NotificationService{
		smtpClient: smtpClient,
	}
}

func (s *NotificationService) SendNotifications(request pkg.EmailRequest) {
	s.smtpClient.SendEmail(request)
}
