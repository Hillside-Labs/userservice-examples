package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	userup "github.com/hillside-labs/userservice-go-sdk/go-client"
)

// Query for a specific event for a given user
func main() {
	userid, _ := strconv.ParseUint(os.Getenv("userid"), 10, 64)
	addr := os.Getenv("userservice_uri")

	_, err := strconv.Atoi(os.Getenv("past_minutes"))
	if err != nil {
		log.Fatal(err)
	}

	minutes, err := strconv.Atoi(os.Getenv("past_minutes"))
	if err != nil {
		log.Fatal(err)
	}
	since := time.Now().Add(-time.Duration(minutes) * time.Minute)

	client, err := userup.NewClient(addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	if err != nil {
		log.Println(err)
	}

	events, err := client.SearchEvents(context.Background(), userid, []string{"userup.io/example/issue-encountered"}, since, time.Now())
	if err != nil {
		log.Println(err)
	}

	for _, event := range events {
		fmt.Printf("%+v\n", event)
	}
}
