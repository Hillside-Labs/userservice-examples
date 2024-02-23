package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
