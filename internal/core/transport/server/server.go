package core_server

import (
	core_logger "MetricService/internal/core/logger"
	feature_metrics "MetricService/internal/feature/metrics"
	"context"
	"fmt"
	"net"
	"time"

	metric "github.com/Hodorev-Evgeny/ShareSystemMonitoring/api/metric-service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	metric.UnimplementedMetricsServer
	config  ServerConfig
	log     *core_logger.Logger
	metcase *feature_metrics.MetricCase
}

func NewServer(
	config ServerConfig,
	logger *core_logger.Logger,
	metcase *feature_metrics.MetricCase,
) *Server {
	return &Server{
		config:  config,
		log:     logger,
		metcase: metcase,
	}
}

func (s *Server) StartServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.config.Addr))
	if err != nil {
		s.log.Error("Failed to listen", zap.Error(err))
		return err
	}
	defer lis.Close()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	metric.RegisterMetricsServer(grpcServer, s)

	reflection.Register(grpcServer)

	ch := make(chan error)
	go func() {
		defer close(ch)

		err := grpcServer.Serve(lis)

		if err != nil {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		return fmt.Errorf("server error: %w", err)

	case <-ctx.Done():
		s.log.Info("Start shutting down server")

		stop := make(chan struct{})
		go func() {
			grpcServer.GracefulStop()
			close(stop)
		}()

		select {
		case <-stop:
			s.log.Info("Shutting down server")
		case <-time.After(1 * s.config.Timeout):
			s.log.Info("Timeout exceeded")
			grpcServer.Stop()
		}
	}

	return nil
}
