package main

import (
	"context"
	"encoding/json"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/hillside-labs/userservice-go-sdk/pkg/userapi"
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

	prefsJson, _ := json.Marshal(prefs)

	prefsStruct := &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"uiPreferences": {
				Kind: &structpb.Value_StringValue{
					StringValue: string(prefsJson),
				},
			},
		},
	}

	resp, err := client.Update(ctx, &userapi.UserRequest{Id: userid, Attributes: prefsStruct})
	return resp, err
}
