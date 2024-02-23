package main

import (
	"context"
	"log"
	"os"
	"strconv"

	userup "github.com/hillside-labs/userservice-go-sdk/go-client"
)

func main() {
	userid, _ := strconv.ParseUint(os.Getenv("userid"), 10, 64)
	addr := os.Getenv("userservice_uri")

	client, err := userup.NewClient(addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	// Method #1: Update the user with included attributes
	user := &userup.User{
		Id: userid,
	}

	prefs := Prefs{
		Theme:             "darcula",
		FontSize:          12,
		Language:          "en-IN",
		ShowNotifications: true,
	}
	user.Attributes["uiPreferences"] = prefs
	client.UpdateUser(context.Background(), user)

	// Method #2: Save the specific attribute for the user
	client.AddAttribute(context.Background(), userid, "uiPreferences", prefs)

	if err != nil {
		log.Fatal(err)
	}
}
