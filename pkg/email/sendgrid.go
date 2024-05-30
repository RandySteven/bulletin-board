package email

import (
	"bytes"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type SendGrid struct {
	Client   *sendgrid.Client
	To       string
	Subject  string
	Metadata map[string]interface{}
}

func NewSendGrid(to, subject string, metadata map[string]interface{}) *SendGrid {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	log.Println("Using SendGrid API Key:", apiKey)
	client := sendgrid.NewSendClient(apiKey)
	return &SendGrid{
		Client:   client,
		To:       to,
		Subject:  subject,
		Metadata: metadata,
	}
}

func (e *SendGrid) sendEmail(templatePath string) (*rest.Response, error) {
	htmlContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		log.Fatalf("Failed to read HTML file: %v", err)
		return nil, err
	}

	tmpl, err := template.New("emailTemplate").Parse(string(htmlContent))
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
		return nil, err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, e.Metadata); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
		return nil, err
	}

	from := mail.NewEmail("From Sender", os.Getenv("CONFIG_AUTH_EMAIL"))
	to := mail.NewEmail("To Receiver", e.To)
	message := mail.NewSingleEmail(from, e.Subject, to, "", body.String())

	response, err := e.Client.Send(message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
		return nil, err
	}

	if response.StatusCode != 200 && response.StatusCode != 202 {
		log.Printf("Email not sent. Status code: %d, response body: %s\n", response.StatusCode, response.Body)
	} else {
		log.Println("Email sent successfully!")
	}
	return response, nil
}

func (e *SendGrid) SendEmailRegister() (*rest.Response, error) {
	return e.sendEmail("./scripts/email/register.html")
}

func (e *SendGrid) SendEmailTest() (*rest.Response, error) {
	return e.sendEmail("./scripts/email/index.html")
}
