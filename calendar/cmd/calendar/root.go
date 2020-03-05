package main

import (
	"github.com/mpuzanov/otus-go/calendar/cmd/calendar/grpc"
	"github.com/mpuzanov/otus-go/calendar/cmd/calendar/web"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Calendar is a calendar microservice demo",
}

func init() {
	rootCmd.AddCommand(grpc.ServerCmd, grpc.GrpcClientCmd, web.ServerCmd)
}
