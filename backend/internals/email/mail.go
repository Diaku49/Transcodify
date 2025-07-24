package email

import (
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/Diaku49/FoodOrderSystem/backend/internals/constants"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/model"
	"github.com/Diaku49/FoodOrderSystem/backend/utilities"
	"gopkg.in/gomail.v2"
)

type MailClient struct {
	Dialer                *gomail.Dialer
	From                  string
	ResetPasswordTemplate *template.Template
}

var MailC *MailClient

func InitMail() error {
	host := os.Getenv("MAIL_HOST")

	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		return err
	}

	user := os.Getenv("MAIL_USER")
	password := os.Getenv("MAIL_PASSWORD")
	from := os.Getenv("MAIL_FROM")

	// Set up the SMTP client
	dialer := gomail.NewDialer(host, port, user, password)
	dialer.SSL = true // force SSL

	// Load the reset password template
	// Maybe change this to a map of templates
	resetPasswordTemplate, err := LoadTemplate(constants.ResetPasswordTemplate)
	if err != nil {
		return err
	}

	MailC = &MailClient{
		Dialer:                dialer,
		From:                  from,
		ResetPasswordTemplate: resetPasswordTemplate,
	}

	return nil
}

func LoadTemplate(templateName string) (*template.Template, error) {
	filePath := fmt.Sprintf("templates/%s.html", templateName)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New(templateName).Parse(string(content))
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func (mc *MailClient) SendResetPasswordEmail(to string, subject string, data model.ResetPasswordMailData) error {
	body, err := utilities.RenderTemplate(mc.ResetPasswordTemplate, data)
	if err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", mc.From)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)

	return mc.Dialer.DialAndSend(msg)
}
