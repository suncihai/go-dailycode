package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

func SendEmail() {
      smtpHost := "smtp.gmail.com"
	  smtpPort := "587"

	  auth := smtp.PlainAuth("", os.Getenv("fromEmail"), os.Getenv("fromEmailPassword"), smtpHost)
	  email, _ := template.ParseFiles("template.html")

	  to := []string{os.Getenv("toEmail")}

	  var body bytes.Buffer

	  mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	  body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	  email.Execute(&body, struct {
		Name    string
		Message string
	  }{
		Name:    "Puneet Singh",
		Message: "This is a test message in a HTML template",
	  })

	  fmt.Println("email service run")

      err := smtp.SendMail(smtpHost+":"+smtpPort, auth, os.Getenv("fromEmail"), to, body.Bytes())	
	  if err != nil {
		  fmt.Println(err)
		  return
	  }
	  fmt.Println("Email sent!")  
}