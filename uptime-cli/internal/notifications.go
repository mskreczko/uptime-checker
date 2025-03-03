package internal

import (
	"fmt"
	"github.com/mskreczko/uptime-checker/pkg"
	"gopkg.in/yaml.v3"
	"log"
)

type NotificationChannel string

const (
	SMS     NotificationChannel = "SMS"
	EMAIL                       = "EMAIL"
	WEBHOOK                     = "WEBHOOK"
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
	case "WEBHOOK":
		*v = WEBHOOK
	}
	return nil
}

type NotificationService struct {
	smtpClient pkg.AWSSesClient
	config     NotificationSettings
}

func NewNotificationService(smtpClient pkg.AWSSesClient) *NotificationService {
	return &NotificationService{
		smtpClient: smtpClient,
	}
}

func (s *NotificationService) SendNotifications(request pkg.EmailRequest) {
	s.smtpClient.SendEmail(request)
}

func (s *NotificationService) SendServicesDownNotification(failedHealthcheck Healthcheck) {
	for _, setting := range s.config.SettingEntries {
		switch setting.Channel {
		case EMAIL:
			s.handleEmailNotifications(failedHealthcheck, setting)
		case SMS:
			s.handleSMSNotifications(failedHealthcheck, setting)
		case WEBHOOK:
			s.handleWebhookNotifications(failedHealthcheck, setting)
		}
	}
}

func (s *NotificationService) handleEmailNotifications(failedHealthcheck Healthcheck, settings NotificationSettingEntry) {
	for _, receiver := range settings.Receivers {
		s.smtpClient.SendEmail(pkg.EmailRequest{
			To:      receiver,
			Subject: fmt.Sprintf("%s is down", failedHealthcheck.Url),
			Body:    fmt.Sprintf("%s is down, last successfull healthcheck: %s", failedHealthcheck.Url, failedHealthcheck.LastUp.String()),
		})
	}
}

func (s *NotificationService) handleSMSNotifications(failedHealthcheck Healthcheck, settings NotificationSettingEntry) {
	log.Println("SMS notifications are not yet implemented.")
}

func (s *NotificationService) handleWebhookNotifications(failedHealthcheck Healthcheck, settings NotificationSettingEntry) {
	for _, receiver := range settings.Receivers {
		pkg.SendNotificationToWebhook(receiver, fmt.Sprintf("%s is down, last successfull healthcheck: %s", failedHealthcheck.Url, failedHealthcheck.LastUp.String()))
	}
}
