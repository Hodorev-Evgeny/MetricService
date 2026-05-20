package main

import (
	core_logger "MetricService/internal/core/logger"
	core_server "MetricService/internal/core/transport/server"
	"context"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	ctx, close := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer close()

	conf_log := core_logger.NewLoggerConfigMust()
	log, err := core_logger.NewLogger(conf_log)
	if err != nil {
		panic(err)
	}

	serv_conf := core_server.GetServerConfigMust()
	server := core_server.NewServer(
		serv_conf,
		log,
	)

	log.Info("Starting metricservice")
	if err := server.StartServer(ctx); err != nil {
		log.Error("failed to start server", zap.Error(err))
	}
}
