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
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hillside-labs/userservice-go-sdk/pkg/userapi"
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

	client, conn, err := GetUserClient(addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	resp, err := client.SearchEvents(context.Background(),
		&userapi.SearchEventsRequest{
			UserId: &userapi.UserID{
				Id: userid,
			},
			Names: []string{"userup.io/example/issue-encountered"},
			Begin: timestamppb.New(since),
		})
	if err != nil {
		log.Println(err)
	}

	for _, event := range resp.Events {
		jsonbytes, err := protojson.MarshalOptions{Multiline: true, Indent: "  ", EmitUnpopulated: true}.Marshal(event)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(string(jsonbytes))
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
