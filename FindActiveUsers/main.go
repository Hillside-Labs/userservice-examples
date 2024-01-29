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
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/hillside-labs/userservice/rpc/userapi"
)

// Find users that have a `last_active` attribute that is older than `days` ago.
func main() {
	inactive, _ := strconv.Atoi(os.Getenv("days"))
	dburi := os.Getenv("dburi")

	client, conn, err := GetUserClient(dburi)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	timeVal, err := structpb.NewValue(time.Now().AddDate(0, 0, -inactive))
	if err != nil {
		log.Println(err)
	}

	attrFilter := userapi.AttributeFilter{
		Name: "last_active",
		Value: timeVal,
		Operator: userapi.Operator_LESS_THAN,
	}

	if err != nil {
		log.Println(err)
	}

	q := &userapi.UserQuery{
		OrderBy:    "username",
		AttributeFilters: []*userapi.AttributeFilter{&attrFilter},
	}

	userResponse, err := client.Find(context.Background(), q)
	if err != nil {
		log.Println(err)
	}

	for _, user := range userResponse.Users {
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
