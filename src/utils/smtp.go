package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

type Email struct {
    From    string
    To      string
    Subject string
    Body    string
}

func SMTP(username string, email string, otp string) {

    var e Email

	// Retrieve env variables
	smtpAddress := os.Getenv("SMTP_ADDRESS")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")

    auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

    e.From = "no-reply@IMS.com"
    e.To = email
    e.Subject = "IMS One-Time password"
    e.Body = `
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

    err := smtp.SendMail(
        smtpAddress, // mailtrap.io SMTP server and port
        auth,
        "no-reply@IMS.com", // sender's email address
        []string{e.To}, // recipient's email address
        []byte(e.Message()), // email message
    )

    CheckError(err)
}

func (e Email) Message() string {
    return fmt.Sprintf(`MIME-version: 1.0;
Content-Type: text/html; charset="UTF-8";
From: %s
To: %s
Subject: %s

%s
`, e.From, e.To, e.Subject, e.Body)
}