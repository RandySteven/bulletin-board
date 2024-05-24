package email

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type Email struct {
	Dialer   *gomail.Dialer
	To       string
	Subject  string
	Metadata map[string]interface{}
}

func NewEmail(to, subject string, metadata map[string]interface{}) *Email {
	log.Println(os.Getenv("CONFIG_SMTP_HOST"))
	return &Email{
		Dialer:   gomail.NewDialer(os.Getenv("CONFIG_SMTP_HOST"), 587, os.Getenv("CONFIG_AUTH_EMAIL"), os.Getenv("CONFIG_AUTH_PASSWORD")),
		To:       to,
		Subject:  subject,
		Metadata: metadata,
	}
}

func (e *Email) SendEmailRegister() {
	htmlFile := "./scripts/email/register.html"
	htmlContent, err := ioutil.ReadFile(htmlFile)
	if err != nil {
		log.Fatalf("Failed to read HTML file: %v", err)
		return
	}

	tmpl, err := template.New("emailTemplate").Parse(string(htmlContent))
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
		return
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, e.Metadata); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
		return
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("CONFIG_AUTH_EMAIL"))
	mailer.SetHeader("To", e.To)
	mailer.SetHeader("Subject", e.Subject)
	mailer.SetBody("text/html", body.String())

	if err := e.Dialer.DialAndSend(mailer); err != nil {
		log.Fatalf("Failed to send email: %v", err)
		return
	}
	log.Println("Email sent successfully!")
}
