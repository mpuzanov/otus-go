package grpcserver

import (
	"log"
	"net"
	"net/http"

	promgrpc "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/mpuzanov/otus-go/calendar/internal/calendar"
	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/mpuzanov/otus-go/calendar/pkg/calendar/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Start GRPC service
func Start(conf *config.Config, logger *zap.Logger, evs *calendar.Calendar) error {

	s := &GRPCServer{
		cfg:          conf,
		logger:       logger,
		eventService: evs,
	}

	l, err := net.Listen("tcp", s.cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("Cannot listen: %s\n", err)
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(promgrpc.StreamServerInterceptor),
		grpc.UnaryInterceptor(promgrpc.UnaryServerInterceptor),
	)
	promgrpc.Register(grpcServer)
	promgrpc.EnableHandlingTimeHistogram()
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		s.logger.Info("Start http server for prometheus", zap.String("address", s.cfg.Prom.GRPCAddr))
		log.Fatal(http.ListenAndServe(s.cfg.Prom.GRPCAddr, nil))
	}()

	api.RegisterCalendarServer(grpcServer, s)

	log.Printf("Starting gRPC server %s, file log: %s\n", s.cfg.GRPCAddr, s.cfg.Log.File)
	s.logger.Info("Starting gRPC server", zap.String("address", s.cfg.GRPCAddr))

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("Cannot start gRPC server: %s\n", err)
	}

	return nil
}
