package pkg

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type AWSSesClient struct {
	Region          string `yaml:"region"`
	Sender          string `yaml:"sender"`
	AccessKeyID     string `yaml:"access_key"`
	SecretAccessKey string `yaml:"secret_key"`
	Token           string `yaml:"token,omitempty"`
}

type EmailRequest struct {
	To      string
	Subject string
	Body    string
}

func (c *AWSSesClient) SendEmail(request EmailRequest) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(c.Region),
		Credentials: credentials.NewStaticCredentials(c.AccessKeyID, c.SecretAccessKey, ""),
	})

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(request.To),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(request.Body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(request.Subject),
			},
		},
		Source: aws.String(c.Sender),
	}

	_, err = svc.SendEmail(input)
	if err != nil {
		fmt.Println(err.Error())
	}
}
