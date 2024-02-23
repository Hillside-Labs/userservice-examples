package main

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hillside-labs/userservice-go-sdk/pkg/userapi"
)

type EventLoggerConfig struct {
	Source      string
	SpecVersion string
}

type EventLogger struct {
	config EventLoggerConfig
}

type InviteUserEvent struct {
	Email string
	From  string
}

func (e EventLogger) LogEvent(ctx context.Context, client userapi.UsersClient, userId uint64, dataType string, schema string, subject string, data interface{}) (*userapi.EventResponse, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	event := &userapi.Event{
		UserId:          userId,
		Source:          e.config.Source,
		Type:            dataType,
		Data:            dataBytes,
		Specversion:     e.config.SpecVersion,
		Timestamp:       timestamppb.Now(),
		Id:              uuid.New().String(),
		Datacontenttype: "application/json",
		Subject:         subject,
		Dataschema:      schema,
	}
	return client.LogEvent(context.Background(), &userapi.EventRequest{
		Event: event,
	})
}
