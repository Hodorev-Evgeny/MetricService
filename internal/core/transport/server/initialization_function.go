package core_server

import (
	"context"

	metric "github.com/Hodorev-Evgeny/ShareSystemMonitoring/api/metric-service"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetCPUWorkload(ctx context.Context, req *emptypb.Empty) (*metric.WorkloadReply, error) {
	s.log.Info("Starting GetCPUWorkload")

	ans, err := s.metcase.GetCPUMetrics()
	if err != nil {
		s.log.Error("Error getting CPU metrics", zap.Error(err))
		return nil, err
	}

	return &metric.WorkloadReply{
		Allworkload: float32(ans.AllWorkload),
	}, nil
}

func (s *Server) GetMemory(req *emptypb.Empty, stream metric.Metrics_GetMemoryServer) error {
	s.log.Info("Starting GetMemory")

	ans := s.metcase.GetMemory()
	for {
		select {
		case <-stream.Context().Done():
			return nil

		case data, ok := <-ans:
			if !ok {
				return nil
			}

			err := stream.Send(&metric.MemoryReply{
				TotalMemory: float32(data.TotalMemory),
				UsedMemory:  float32(data.UsedMemory),
				Precent:     float32(data.PrecentMemory),
			})

			if err != nil {
				s.log.Error("Error sending data", zap.Error(err))
				return err
			}
		}
	}
}
