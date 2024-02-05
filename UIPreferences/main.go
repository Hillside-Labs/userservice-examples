package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/hillside-labs/userservice-go-sdk/pkg/userapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// bulk setting?
func main() {
	userid, _ := strconv.ParseUint(os.Getenv("userid"), 10, 64)
	addr := os.Getenv("userservice_uri")

	client, conn, err := GetUserClient(addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	_, err = SavePreferences(userid, client, Prefs{
		Theme:             "darcula",
		FontSize:          12,
		Language:          "en-IN",
		ShowNotifications: true,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func GetUserClient(dburi string) (userapi.UsersClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(dburi, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	client := userapi.NewUsersClient(conn)
	return client, conn, nil
}
