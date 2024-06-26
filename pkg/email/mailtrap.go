package email

import (
	"bytes"
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type Mailtrap struct {
	Dialer   *gomail.Dialer
	To       string
	Subject  string
	Metadata map[string]interface{}
}

func NewMailtrap(to, subject string, metadata map[string]interface{}) *Mailtrap {
	host := "sandbox.smtp.mailtrap.io"
	port := 587
	username := "4a87f6082d6054"
	password := "a711952511e0bf"

	log.Println("SMTP Host:", host)
	log.Println("SMTP Port:", port)
	log.Println("SMTP Username:", username)

	dialer := gomail.NewDialer(host, port, username, password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &Mailtrap{
		Dialer:   dialer,
		To:       to,
		Subject:  subject,
		Metadata: metadata,
	}
}

func (e *Mailtrap) sendEmail(templatePath string) (err error) {
	htmlContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		log.Fatalf("Failed to read HTML file: %v", err)
		return err
	}

	tmpl, err := template.New("emailTemplate").Parse(string(htmlContent))
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, e.Metadata); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("CONFIG_AUTH_EMAIL"))
	mailer.SetHeader("To", e.To)
	mailer.SetHeader("Subject", e.Subject)
	mailer.SetBody("text/html", body.String())

	log.Println("Sending email to:", e.To)
	if err := e.Dialer.DialAndSend(mailer); err != nil {
		log.Fatalf("Failed to send email: %v", err)
		return err
	}
	log.Println("Email sent successfully!")
	return nil
}

func (e *Mailtrap) SendEmailRegister() error {
	return e.sendEmail("./scripts/email/register.html")
}

func (e *Mailtrap) SendEmailTest() error {
	return e.sendEmail("./scripts/email/index.html")
}
