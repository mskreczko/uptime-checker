package pkg

import (
	"github.com/wneessen/go-mail"
	"log"
)

type SMTPClient struct {
	server   string
	username string
	password string
	client   *mail.Client
}

func (c *SMTPClient) NewSMTPClient() *mail.Client {
	client, err := mail.NewClient(c.server, mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(c.username), mail.WithPassword(c.password))
	if err != nil {
		log.Fatalf("Failed to create SMTP client: %v", err)
	}
	return client
}

type EmailRequest struct {
	from    string
	to      string
	subject string
	body    string
}

func (c *SMTPClient) SendEmail(request EmailRequest) {
	message := mail.NewMsg()
	if err := message.From(request.from); err != nil {
		log.Fatalf("Failed to set From address: %s", err)
	}
	if err := message.To(request.to); err != nil {
		log.Fatalf("Failed to set To address: %s", err)
	}
	message.Subject(request.subject)
	message.SetBodyString(mail.TypeTextPlain, request.body) // TODO add html body

	if err := c.client.Send(message); err != nil {
		log.Fatalf("Failed to send email: %s", err)
	}
}
