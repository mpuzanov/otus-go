package integration

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cucumber/godog"
	"github.com/golang/protobuf/ptypes"
	"github.com/mpuzanov/otus-go/calendar/pkg/calendar/api"
	"google.golang.org/grpc"
)

var (
	grpcListen string = os.Getenv("GRPC_LISTEN")
)

type testCalendar struct {
	conn           *grpc.ClientConn
	client         api.CalendarClient
	eventProto     *api.Event
	addResponse    *api.AddResponseResult
	updateResponse *api.UpdateResponseResult
	deleteResponse *api.DelResponseResult
	responseErr    error
}

func init() {
	if grpcListen == "" {
		grpcListen = "0.0.0.0:50051"
	}
	//log.Println("grpcListen", grpcListen)
}

func (t *testCalendar) startSuite() {

	var err error
	//============= создаём событие для проверки ====================
	UserName := "User1"
	startTime, _ := time.Parse("2006-01-02 15:04", "2020-04-01 09:00")
	startTimeProto, err := ptypes.TimestampProto(startTime)
	if err != nil {
		log.Fatal(err)
		//return err
	}
	endTime, _ := time.Parse("2006-01-02 15:04", "2020-04-01 10:30")
	endTimeProto, err := ptypes.TimestampProto(endTime)
	if err != nil {
		log.Fatal(err)
		//return err
	}
	ReminderBefore := time.Duration(30 * time.Minute)
	t.eventProto = &api.Event{
		Header:         "event 1",
		Text:           "text event 1",
		StartTime:      startTimeProto,
		EndTime:        endTimeProto,
		User:           UserName,
		ReminderBefore: ptypes.DurationProto(ReminderBefore),
	}
	//log.Println("eventProto", t.eventProto)

	t.conn, err = grpc.Dial(grpcListen, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial : %s, %v", grpcListen, err)
	}
	t.client = api.NewCalendarClient(t.conn)
}

func (t *testCalendar) stopSuite() {
	t.conn.Close()
}

func (t *testCalendar) iCallGrpcCalendarMethodAddEvent() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	t.addResponse, t.responseErr = t.client.AddEvent(ctx, t.eventProto)
	if t.responseErr != nil {
		return t.responseErr
	}
	if t.addResponse.GetError() != "" {
		return fmt.Errorf(t.addResponse.GetError())
	}
	//log.Println("addResponse", t.addResponse)
	t.eventProto.Id = t.addResponse.GetId()
	return nil
}

func (t *testCalendar) theErrorShouldBeNil() error {
	return t.responseErr
}

func (t *testCalendar) theAddResponseSuccessShouldBeTrue() error {
	if !t.addResponse.GetStatus() {
		return fmt.Errorf("addResponse success is false")
	}
	return nil
}

func (t *testCalendar) iHaveTheAddEvent() error {
	if t.eventProto.GetId() == "" {
		return fmt.Errorf("not have AddEvent")
	}
	return nil
}

func (t *testCalendar) iCallGrpcCalendarMethodUpdateEvent() error {

	// conn, err := grpc.Dial(grpcListen, grpc.WithInsecure())
	// if err != nil {
	// 	return fmt.Errorf("fail to dial : %s, %v", grpcListen, err)
	// }
	// defer conn.Close()
	// client := api.NewCalendarClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t.updateResponse, t.responseErr = t.client.UpdateEvent(ctx, t.eventProto)
	if t.responseErr != nil {
		return t.responseErr
	}
	if t.updateResponse.GetError() != "" {
		return fmt.Errorf(t.updateResponse.GetError())
	}
	//log.Println("updateResponse", t.updateResponse)
	return nil
}

func (t *testCalendar) theUpdateResponseSuccessShouldBeTrue() error {
	if !t.updateResponse.GetStatus() {
		return fmt.Errorf("updateResponse success is false")
	}
	return nil
}

func (t *testCalendar) iHaveTheEventID() error {
	if t.eventProto.GetId() == "" {
		return fmt.Errorf("not EventID")
	}
	return nil
}

func (t *testCalendar) iCallGrpcCalendarMethodDeleteEvent() error {

	// conn, err := grpc.Dial(grpcListen, grpc.WithInsecure())
	// if err != nil {
	// 	return fmt.Errorf("fail to dial : %s, %v", grpcListen, err)
	// }
	// defer conn.Close()
	// client := api.NewCalendarClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	IDDel := t.eventProto.Id
	t.deleteResponse, t.responseErr = t.client.DelEvent(ctx, &api.EventID{Id: IDDel})
	if t.responseErr != nil {
		return t.responseErr
	}
	if t.deleteResponse.GetError() != "" {
		return fmt.Errorf(t.deleteResponse.GetError())
	}
	//log.Println("deleteResponse", t.deleteResponse)
	return nil
}

func (t *testCalendar) theDeleteResponseSuccessShouldBeTrue() error {
	if !t.deleteResponse.GetStatus() {
		return fmt.Errorf("deleteResponse success is false")
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	test := &testCalendar{}

	s.BeforeSuite(test.startSuite)

	s.Step(`^I call grpc calendar method AddEvent$`, test.iCallGrpcCalendarMethodAddEvent)
	s.Step(`^The error should be nil$`, test.theErrorShouldBeNil)
	s.Step(`^The add response success should be true$`, test.theAddResponseSuccessShouldBeTrue)

	s.Step(`^I have the Event$`, test.iHaveTheAddEvent)
	s.Step(`^I call grpc calendar method UpdateEvent$`, test.iCallGrpcCalendarMethodUpdateEvent)
	s.Step(`^The update response success should be true$`, test.theUpdateResponseSuccessShouldBeTrue)

	s.Step(`^I have the event ID$`, test.iHaveTheEventID)
	s.Step(`^I call grpc calendar method DeleteEvent$`, test.iCallGrpcCalendarMethodDeleteEvent)
	s.Step(`^The delete response success should be true$`, test.theDeleteResponseSuccessShouldBeTrue)
	s.AfterSuite(test.stopSuite)
}
