package internal

const NotificationChannel (
	SMS
	EMAIL
)

type NotificationSettings struct {
	SettingEntries []NotificationSettingEntry `yaml:"notification_settings"`
}

type NotificationSettingEntry struct {
	Channel NotificationChannel `yaml:"channel"`
	Receivers []string `yaml:"receivers`
}

func validateEmail(email string) bool {

}

func validatePhoneNumber(phoneNumber string) bool {

}