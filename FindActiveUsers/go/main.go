package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userup "github.com/hillside-labs/userservice-go-sdk/go-client"
	"github.com/hillside-labs/userservice-go-sdk/pkg/userapi"
)

// Find users that have a `last_active` attribute that is older than `days` ago.
func main() {
	inactive, _ := strconv.Atoi(os.Getenv("days"))
	addr := os.Getenv("userservice_uri")

	client, err := userup.NewClient(addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	begin := time.Now().AddDate(0, 0, -inactive)

	searchParams := userup.UserSearchParams{}.WithAttribute("last_active", begin, userapi.Operator_LESS_THAN_OR_EQUALS)
	users, err := client.FindUser(context.Background(), &searchParams)
	if err != nil {
		log.Println(err)
	}

	for _, user := range users {
		fmt.Println(user.Id)
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
