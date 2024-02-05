package main

import (
	"context"

	"github.com/hillside-labs/userservice-go-sdk/pkg/userapi"
	"google.golang.org/protobuf/types/known/structpb"
)

type Prefs struct {
	// Visual Preferences
	Theme       string `json:"theme"`
	FontSize    int    `json:"fontSize"`
	ColorScheme string `json:"colorScheme"`

	// Behavior Preferences
	Language          string `json:"language"`
	AutoSave          bool   `json:"autoSave"`
	ShowNotifications bool   `json:"showNotifications"`

	// Feature-specific Preferences
	CachePages CacheConfig `json:"cacheConfig"`
}

type CacheConfig struct {
	Enabled   bool `json:"enabled"`
	Threshold int  `json:"threshold"`
}

func SavePreferences(userid uint64, client userapi.UsersClient, prefs Prefs) (*userapi.UserResponse, error) {
	ctx := context.Background()
	
    resp, err := client.Get(ctx, &userapi.UserRequest{
		Id: userid,
	})
	if err != nil {
		return nil, err
	}

	attrMap := resp.Attributes.AsMap()
	attrMap["ui_preferences"] = prefs
	attrs, _ := structpb.NewStruct(attrMap)
	
    resp, err = client.Update(ctx, &userapi.UserRequest{Attributes: attrs})
	return resp, err
}
