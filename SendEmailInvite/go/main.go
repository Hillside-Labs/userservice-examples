package main

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"

	userup "github.com/hillside-labs/userservice-go-sdk/go-client"
)

func main() {
	addr := os.Getenv("USERSERVICE_URI")
	smtpPassword := os.Getenv("USERSERVICE_SMTP_PASSWORD")
	smtpUsername := os.Getenv("USERSERVICE_SMTP_SENDEREMAIL")
	email := os.Getenv("USERSERVICE_EMAIL")

	client, err := userup.NewClient(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	logger := userup.NewLogger(userup.NewLoggerConfig("https://github.com/hillside-labs/userservice-examples/SendEmailInviteClient/go", client))

	user, err := client.AddUser(context.Background(), &userup.User{
		Username: email,
		Attributes: map[string]interface{}{
			"status": "pending",
		},
	})
	if err != nil {
		log.Panicf("Error creating user: %v", err)
	}

	logger.LogEvent(context.Background(),
		user.Id,
		"InviteUserEvent",
		"userup.io/example/invite",
		strconv.FormatUint(user.Id, 10),
		user)

	SendEmail(email, smtpUsername, smtpPassword, "Invite User", "You are invited to userup.io")
}

func SendEmail(toEmail, username, password, subject, body string) error {
	return smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", username, password, "smtp.gmail.com"),
		username,
		[]string{toEmail},
		[]byte(fmt.Sprintf("Subject: %s\n\n%s", subject, body)),
	)
}
