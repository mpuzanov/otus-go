package grpcserver

import (
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/mpuzanov/otus-go/calendar/internal/model"
	"github.com/mpuzanov/otus-go/calendar/pkg/calendar/api"
)

//EventProtoMsgToEvent переконвертируем событие protoMsg в структуру golang
func EventProtoMsgToEvent(in *api.Event) (*model.Event, error) {
	ID := uuid.Nil
	if in.GetId() != "" {
		id, err := uuid.Parse(in.GetId())
		if err != nil {
			return nil, fmt.Errorf("uuid.parse Id, %w", err)
		}
		ID = id
	}

	startTime, err := ptypes.Timestamp(in.GetStartTime())
	if err != nil {
		return nil, fmt.Errorf("convert StartTime, %w", err)
	}
	endTime, err := ptypes.Timestamp(in.GetEndTime())
	if err != nil {
		return nil, fmt.Errorf("convert EndTime, %w", err)
	}
	reminderBefore, err := ptypes.Duration(in.GetReminderBefore())
	if err != nil {
		return nil, fmt.Errorf("convert ReminderBefore, %w", err)
	}

	event := model.Event{
		ID:             ID,
		Header:         in.GetHeader(),
		Text:           in.GetText(),
		StartTime:      startTime,
		EndTime:        endTime,
		UserName:       in.GetUser(),
		ReminderBefore: reminderBefore,
	}
	return &event, nil
}

//EventToEventProtoMsg переконвертируем событие в структуре golang в protoMsg
func EventToEventProtoMsg(in *model.Event) (*api.Event, error) {

	startTime, err := ptypes.TimestampProto(in.StartTime)
	if err != nil {
		return nil, fmt.Errorf("convert StartTime, %w", err)
	}
	endTime, err := ptypes.TimestampProto(in.EndTime)
	if err != nil {
		return nil, fmt.Errorf("convert EndTime, %w", err)
	}

	eventProto := api.Event{
		Id:             in.ID.String(),
		Header:         in.Header,
		Text:           in.Text,
		StartTime:      startTime,
		EndTime:        endTime,
		User:           in.UserName,
		ReminderBefore: ptypes.DurationProto(in.ReminderBefore),
	}
	return &eventProto, nil
}
