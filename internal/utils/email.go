package utils

import (
	"bytes"
	"crypto/tls"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

type SMTP struct {
	EmailFrom string
	User      string
	Pass      string
	Host      string
	Port      int
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func (s SMTP) SendEmail(email string, data *EmailData, emailTemp string) error {
	// Sender data.
	from := s.EmailFrom
	smtpPass := s.Pass
	smtpUser := s.User
	to := email
	smtpHost := s.Host
	smtpPort := s.Port

	var body bytes.Buffer

	template, err := ParseTemplateDir("internal/templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template.ExecuteTemplate(&body, emailTemp, &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	return d.DialAndSend(m)
}
