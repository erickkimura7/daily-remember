package main

import (
	"bytes"
	"text/template"

	nt "github.com/erickkimura7/daily-remember/notificationEvent"
	"gopkg.in/gomail.v2"
)

func main() {

	from := "teste"
	password := "teste"

	toEmailAddress := "teste"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := 587

	subject := "Subject: This is the subject of the mail\n"

	body, err := generateEmailBody()
	if err != nil {
		panic(err)
	}
	// body := "This is the body of the mailHello <b>Bob</b>!"

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, from, password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func generateEmailBody() (string, error) {
	t, err := template.ParseFiles("./emailsender/main.html")
	if err != nil {
		return "", err
	}

	templateData := struct {
		Name   string
		URL    string
		Events []nt.Event
	}{
		Name: "Dhanush",
		URL:  "http://geektrust.in",
		Events: []nt.Event{{
			Title:       "teste title",
			Description: "teste description",
		}},
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, templateData); err != nil {
		return "", err
	}

	return buf.String(), nil
}
