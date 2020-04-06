package grpcserver

import (
	"context"

	"github.com/mpuzanov/otus-go/calendar/internal/calendar"
	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/mpuzanov/otus-go/calendar/pkg/calendar/api"
	"go.uber.org/zap"
)

//GRPCServer структура сервера
type GRPCServer struct {
	cfg          *config.Config
	logger       *zap.Logger
	eventService *calendar.Calendar
}

//AddEvent Добавить событие
func (s *GRPCServer) AddEvent(ctx context.Context, req *api.Event) (*api.AddResponseResult, error) {
	event, err := EventProtoMsgToEvent(req)
	if err != nil {
		s.logger.Error("GRPCServer AddEvent", zap.String("EventProtoMsgToEvent", err.Error()))
		return &api.AddResponseResult{Status: false, Error: err.Error()}, err
	}
	id, err := s.eventService.AddEvent(event)
	if err != nil {
		s.logger.Error("GRPCServer AddEvent", zap.String("eventService.AddEvent", err.Error()))
		return &api.AddResponseResult{Status: false, Id: "", Error: err.Error()}, err
	}
	s.logger.Info("GRPCServer AddEvent", zap.Bool("status", true), zap.String("Id", event.ID.String()))
	return &api.AddResponseResult{Status: true, Id: id, Error: ""}, nil

}

//UpdateEvent Обновить событие
func (s *GRPCServer) UpdateEvent(ctx context.Context, req *api.Event) (*api.UpdateResponseResult, error) {
	event, err := EventProtoMsgToEvent(req)
	if err != nil {
		s.logger.Error("GRPCServer UpdateEvent", zap.String("EventProtoMsgToEvent", err.Error()))
		return &api.UpdateResponseResult{Status: false, Error: err.Error()}, err
	}
	res, err := s.eventService.UpdateEvent(event)
	if err != nil {
		s.logger.Error("GRPCServer UpdateEvent", zap.String("eventService.UpdateEvent", err.Error()))
		return &api.UpdateResponseResult{Status: res, Id: "", Error: err.Error()}, err
	}
	s.logger.Info("GRPCServer UpdateEvent", zap.Bool("status", true), zap.String("Id", req.GetId()))
	return &api.UpdateResponseResult{Status: res, Id: event.ID.String(), Error: ""}, nil
}

//DelEvent Удалить (ID события)
func (s *GRPCServer) DelEvent(ctx context.Context, req *api.EventID) (*api.DelResponseResult, error) {
	id := req.GetId()
	res, err := s.eventService.DelEvent(id)
	if err != nil {
		s.logger.Error("GRPCServer DelEvent", zap.String("Id", id), zap.String("error", err.Error()))
		return &api.DelResponseResult{Status: res, Error: err.Error()}, err
	}
	s.logger.Info("GRPCServer DelEvent", zap.Bool("status", true), zap.String("Id", id))
	return &api.DelResponseResult{Status: res, Error: ""}, nil
}

//FindEventByID Поиск события по коду
func (s *GRPCServer) FindEventByID(ctx context.Context, req *api.EventID) (*api.EventResponse, error) {
	id := req.GetId()
	event, err := s.eventService.FindEventByID(id)
	if err != nil {
		s.logger.Error("GRPCServer FindEventByID", zap.String("Id", id), zap.String("error", err.Error()))
		return &api.EventResponse{Status: false, Event: nil, Error: err.Error()}, err
	}
	eventMsg, err := EventToEventProtoMsg(event)
	if err != nil {
		s.logger.Error("GRPCServer FindEventByID", zap.String("Id", id), zap.String("EventToEventProtoMsg", err.Error()))
		return &api.EventResponse{Status: false, Event: nil, Error: err.Error()}, err
	}
	s.logger.Info("GRPCServer FindEventByID", zap.Bool("status", true), zap.String("Id", id))
	return &api.EventResponse{Status: true, Event: eventMsg, Error: ""}, nil
}

//GetUserEvents Получаем события пользователя
func (s *GRPCServer) GetUserEvents(ctx context.Context, req *api.RequestUser) (*api.EventsResponse, error) {
	user := req.GetUser()
	events, err := s.eventService.GetUserEvents(user)
	if err != nil {
		s.logger.Error("GRPCServer GetUserEvents", zap.String("user", user), zap.String("error", err.Error()))
		return nil, err
	}
	protoEvents := &api.EventsResponse{}
	for _, v := range events {
		protoEvent, err := EventToEventProtoMsg(&v)
		if err != nil {
			s.logger.Error("GRPCServer GetUserEvents", zap.String("Id", v.ID.String()), zap.String("EventToEventProtoMsg", err.Error()))
			return &api.EventsResponse{Events: nil, Error: err.Error()}, err
		}
		protoEvents.Events = append(protoEvents.Events, protoEvent)
	}
	s.logger.Info("GRPCServer GetUserEvents", zap.Bool("status", true), zap.String("user", user))
	protoEvents.Status = true
	return protoEvents, nil
}
