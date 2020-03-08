package grpc

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/mpuzanov/otus-go/calendar/pkg/calendar/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var server string

var (
	// GrpcClientCmd .
	GrpcClientCmd = &cobra.Command{
		Use:   "grpc_client",
		Short: "Run grpc client",
		Run:   grpcClientStart,
	}
)

func init() {
	GrpcClientCmd.Flags().StringVar(&server, "server", ":50051", "host:port to connect to")
}

func grpcClientStart(cmd *cobra.Command, args []string) {

	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial : %s, %v\n", server, err)
	}
	client := api.NewCalendarClient(conn)
	ctx := context.Background()

	//Проверяем методы календаря
	//============= создаём событие для проверки ====================
	UserName := "User1"
	startTime, _ := time.Parse("2006-01-02 15:04", "2020-04-01 09:00")
	startTimeProto, err := ptypes.TimestampProto(startTime)
	if err != nil {
		log.Fatal(err.Error())
	}
	endTime, _ := time.Parse("2006-01-02 15:04", "2020-04-01 10:30")
	endTimeProto, err := ptypes.TimestampProto(endTime)
	if err != nil {
		log.Fatal(err.Error())
	}
	ReminderBefore := time.Duration(30 * time.Minute)
	eventProto := &api.Event{
		Header:         "event 1",
		Text:           "text event 1",
		StartTime:      startTimeProto,
		EndTime:        endTimeProto,
		User:           UserName,
		ReminderBefore: ptypes.DurationProto(ReminderBefore),
	}
	log.Println("Event: ", eventProto)

	//===============================================================
	// rpc AddEvent(Event) returns (ResponseResult) {}  AddEvent(event *model.Event) (string, error)
	resp, err := client.AddEvent(ctx, eventProto)
	if err != nil {
		log.Fatal(err)
	}
	if resp.GetError() != "" {
		log.Fatal(resp.GetError())
	} else {
		log.Printf("AddEvent. Status: %v. ID: %s", resp.GetStatus(), resp.GetId())
	}
	eventProto.Id = resp.GetId()

	// rpc UpdateEvent(Event) returns (ResponseResult) {} UpdateEvent(event *model.Event) (bool, error)
	respUpd, _ := client.UpdateEvent(ctx, eventProto)
	if respUpd.GetError() != "" {
		log.Fatal(respUpd.GetError())
	} else {
		log.Printf("UpdateEvent. Status: %v. ID: %s", respUpd.GetStatus(), respUpd.GetId())
	}

	// rpc FindEventByID(EventID) returns (EventResponse) {} FindEventByID(id string) (*model.Event, error)
	IDFind := eventProto.Id
	eventFind, _ := client.FindEventByID(ctx, &api.EventID{Id: IDFind})
	if eventFind.GetError() != "" {
		log.Fatal(eventFind.GetError())
	} else {
		log.Printf("FindEventByID. Status: %v. ID: %s", eventFind.GetStatus(), eventFind.Event.GetId())
	}

	// rpc GetUserEvents(RequestUser) returns (EventsResponse) {} GetUserEvents(user string) ([]model.Event, error)
	events, _ := client.GetUserEvents(ctx, &api.RequestUser{User: UserName})
	if events.GetError() != "" {
		log.Fatal(events.GetError())
	} else {
		log.Printf("GetUserEvents. Status: %v. Count: %v", events.GetStatus(), len(events.GetEvents()))
	}

	// rpc DelEvent(EventID) returns (ResponseResult) {} DelEvent(id string) (bool, error)
	IDDel := eventProto.Id
	respDel, _ := client.DelEvent(ctx, &api.EventID{Id: IDDel})
	if resp.GetError() != "" {
		log.Fatal(respDel.GetError())
	} else {
		log.Printf("DelEvent. Status: %v", respDel.GetStatus())
	}
}
