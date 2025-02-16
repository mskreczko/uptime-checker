package pkg

import (
	"github.com/wneessen/go-mail"
	"log"
)

type SMTPConfig struct {
	Server   string `yaml:"server"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
}

type SMTPClient struct {
	client *mail.Client
}

func NewSMTPClient(config SMTPConfig) *SMTPClient {
	client, err := mail.NewClient(config.Server, mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(config.Username), mail.WithPassword(config.Password))
	if err != nil {
		log.Fatalf("Failed to create SMTP client: %v", err)
	}
	return &SMTPClient{client: client}
}

type EmailRequest struct {
	From    string
	To      string
	Subject string
	Body    string
}

func (c *SMTPClient) SendEmail(request EmailRequest) {
	message := mail.NewMsg()
	if err := message.From(request.From); err != nil {
		log.Fatalf("Failed to set From address: %s", err)
	}
	if err := message.To(request.To); err != nil {
		log.Fatalf("Failed to set To address: %s", err)
	}
	message.Subject(request.Subject)
	message.SetBodyString(mail.TypeTextPlain, request.Body) // TODO add html body

	if err := c.client.Send(message); err != nil {
		log.Fatalf("Failed to send email: %s", err)
	}
}
