package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Print("Got a GET, serve /")
			var t = template.Must(template.ParseFiles("template.html"))
			t.Execute(w, nil)
		case "POST":
			log.Print("Got a POST, send a mail and redirect")
			go sendMail()
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		default:
			errString := fmt.Sprintf("Method %v not allowed", r.Method)
			log.Print(errString)
			http.Error(w, errString, http.StatusMethodNotAllowed)
		}
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

var (
	smtpLogin    = os.Getenv("MAILGUN_SMTP_LOGIN")
	smtpPassword = os.Getenv("MAILGUN_SMTP_PASSWORD")
	smtpServer   = os.Getenv("MAILGUN_SMTP_SERVER")
	smtpPort     = os.Getenv("MAILGUN_SMTP_PORT")
)

func sendMail() {
	auth := smtp.PlainAuth("", smtpLogin, smtpPassword, smtpServer)
	to := []string{"yannsalaun1@gmail.com"}
	msg := []byte("Subject: This is the subject\nThis is the body")
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, "sender@example.com", to, msg)
	if err != nil {
		log.Print(err)
	}
}
