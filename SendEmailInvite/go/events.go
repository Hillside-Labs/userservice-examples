package main

import (
	"context"

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

func (e EventLogger) LogEvent(ctx context.Context, client userapi.UsersClient, userId uint64, dataType string, schema string, subject string, dataContentType string, data []byte) (*userapi.EventResponse, error) {
	event := &userapi.Event{
		UserId:          userId,
		Source:          e.config.Source,
		Type:            dataType,
		Data:            data,
		Specversion:     e.config.SpecVersion,
		Timestamp:       timestamppb.Now(),
		Id:              uuid.New().String(),
		Datacontenttype: dataContentType,
		Subject:         subject,
		Dataschema:      schema,
	}
	return client.LogEvent(context.Background(), &userapi.EventRequest{
		Event: event,
	})
}
