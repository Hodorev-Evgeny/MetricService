package feature_metrics

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"go.uber.org/zap"
)

type CPUMetrics struct {
	AllWorkload float64
}

func (c *MetricCase) GetCPUMetrics() (CPUMetrics, error) {
	c.log.Info("Starting GetCPUMetrics")
	precent, err := cpu.Percent(100*time.Millisecond, false)
	if err != nil {
		c.log.Warn("Failed to get cpu metrics", zap.Error(err))
		return CPUMetrics{}, err
	}

	c.log.Info("successfully get cpu metrics", zap.Any("precent", precent))
	return CPUMetrics{
		AllWorkload: precent[0],
	}, err
}
