package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SMTP(username string, email string, otp string) {

	// Retrieve env variables
	smtpAddress := os.Getenv("SMTP_ADDRESS")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")

    auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

    to := email
    subject := "IMS One-Time password"
    body := `
        <html>
            <head>
                <style>
                /* Add some styling for the email */
                body {
                    font-family: Arial, sans-serif;
                }
                .container {
                    max-width: 600px;
                    margin: 0 auto;
                    text-align: center;
                }
                h1 {
                    margin-top: 20px;
                    color: #333;
                }
                p {
                    color: #555;
                    text-align:left;
                    padding-left:20px;
                }
                </style>
            </head>
            <body>
                <div class="container">
                <h1>IMS One-Time Password</h1>
                <p>Hi <b>` + username + `</b>,</p>
                <p>Your one-time password is:</p>
                <h2>` + otp + `</h2>
                <p>This password is valid for only 10 minutes. Please do not share this with anyone.</p>
                <p>Thanks, <br>
                    IMS Team
                </p>
                </div>
            </body>
        </html>
    `

    // Create new message
    message := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
    message += fmt.Sprintf("From: %s\r\n", "no-reply@IMS.com")
    message += fmt.Sprintf("To: %s\r\n", email)
    message += fmt.Sprintf("Subject: %s\r\n", subject)
    message += fmt.Sprintf("\r\n%s\r\n", body)

    err := smtp.SendMail(
        smtpAddress, // mailtrap.io SMTP server and port
        auth,
        "no-reply@IMS.com", // sender's email address
        []string{to}, // recipient's email address
        []byte(message), // email message
    )

    CheckError(err)
}