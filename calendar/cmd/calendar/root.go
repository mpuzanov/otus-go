package main

import (
	"github.com/mpuzanov/otus-go/calendar/cmd/calendar/grpc"
	"github.com/mpuzanov/otus-go/calendar/cmd/calendar/scheduler"
	"github.com/mpuzanov/otus-go/calendar/cmd/calendar/sender"
	"github.com/mpuzanov/otus-go/calendar/cmd/calendar/web"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Calendar is a calendar microservice demo",
}

func init() {
	rootCmd.AddCommand(grpc.ServerCmd, grpc.GrpcClientCmd, web.ServerCmd, scheduler.ServerCmd, sender.ServerCmd)
}
