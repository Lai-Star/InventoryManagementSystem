package utils

import (
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SMTP(username string, email string, otp string) {

	// Loading the dotenv file
	err := godotenv.Load("../../config/.env");
	CheckError(err);

	// Retrieve env variables
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpAddress := os.Getenv("SMTP_ADDRESS")

    to := email
    subject := "IMS One-Time password"
    body := "Dear " + username + ",\n\n" + "\tThis is an auto-generated email. Please do not reply to this email. \n\n\tThis is your OTP: " + 
				otp + "\n\nRegards,\nThe IMS Team"

    msg := "From: no-reply@IMS.com\n" + 
            "To: " + to + "\n" + 
            "Subject: " + subject + "\n\n" + 
            body

    err = smtp.SendMail(
        smtpAddress, // mailtrap.io SMTP server and port
        smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost), // authentication
        "no-reply@IMS.com", // sender's email address
        []string{to}, // recipient's email address
        []byte(msg), // email message
    )

    CheckError(err)
}