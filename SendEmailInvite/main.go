package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/hillside-labs/userservice/rpc/userapi"
)

func main() {
	dburi := os.Getenv("USERSERVICE_URI")
	smtpPassword := os.Getenv("USERSERVICE_SMTP_PASSWORD")
	smtpUsername := os.Getenv("USERSERVICE_SMTP_SENDEREMAIL")
	email := os.Getenv("USERSERVICE_EMAIL")
	el := EventLogger{
		config: EventLoggerConfig{
			Source:      "userservice/cmd/democlient/invite",
			SpecVersion: "1.0.0",
		},
	}

	client, conn, err := GetUserClient(dburi)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	userResponse, err := CreateInviteUser(client, email, smtpUsername)
	if err != nil {
		log.Fatal(err)
	}

	payload := &InviteUserEvent{
		Email: email,
		From:  smtpUsername,
	}

	payloadBytes, _ := json.Marshal(payload)

	el.LogEvent(context.Background(),
		client,
		userResponse.Id,
		"userup.io/example/invite",
		"application/json",
		strconv.FormatUint(userResponse.Id, 10),
		"InviteUserEvent",
		payloadBytes)

	SendEmail(email, smtpUsername, smtpPassword, "Invite User", "You are invided to userup.io")
}

func CreateInviteUser(client userapi.UsersClient, email string, from string) (*userapi.UserResponse, error) {
	attr, err := structpb.NewStruct(
		map[string]interface{}{
			"status": "pending",
		},
	)
	if err != nil {
		return nil, err
	}

	userResponse, err := client.Create(context.Background(), &userapi.NewUser{
		Username:   email,
		Attributes: attr})
	if err != nil {
		return nil, err
	}

	return userResponse, nil
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

func GetUserClient(dburi string) (userapi.UsersClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(dburi, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	client := userapi.NewUsersClient(conn)
	return client, conn, nil
}
