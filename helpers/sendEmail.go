package helpers

import (
	emailServer "grpc-sample/data"
	"net/smtp"
	"os"
)

func SendEmail(data *emailServer.EmailData) *emailServer.Message {
	response := &emailServer.Message{}
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	to := []string{data.Email}
	message := []byte("Subject:" + data.Subject + "\r\n\r\n" + data.Body)
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		response.ResponseText = "There was some error while sending email. Please try again later."
		response.Status = -1
		return response
	}
	response.ResponseText = "Email Sent Successfully!"
	response.Status = 1
	return response
}
func SendEmailDefault(email string, body string, subject string) (map[string]string, int) {
	resp := make(map[string]string)
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	to := []string{email}
	message := []byte("Subject:" + subject + "\r\n\r\n" + body)
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		resp["Message"] = "There is some error in sending email"
		return resp, 500
	}
	resp["Message"] = "Email Sent Successfully!"
	return resp, 200
}
