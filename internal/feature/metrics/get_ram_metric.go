package feature_metrics

import (
	"time"

	"github.com/shirou/gopsutil/v3/mem"
	"go.uber.org/zap"
)

type RAMMetric struct {
	TotalMemory   float64
	UsedMemory    float64
	PrecentMemory float64
}

func (c *MetricCase) GetMemory() <-chan RAMMetric {
	ch := make(chan RAMMetric, 5)
	c.log.Info("Starting GetMemory")

	go func() {
		timer := time.NewTimer(5 * time.Second)
		for {
			select {
			case <-timer.C:
				close(ch)
				c.log.Info("Stopping GetMemory")
				return

			default:
				mem, err := mem.VirtualMemory()
				if err != nil {
					c.log.Warn("Failed to get virtual memory info", zap.Error(err))
					ch <- RAMMetric{}
				}

				total := float64(mem.Total) / 1024 / 1024 / 1024
				used := float64(mem.Used) / 1024 / 1024 / 1024

				ch <- RAMMetric{
					TotalMemory:   total,
					UsedMemory:    used,
					PrecentMemory: mem.UsedPercent,
				}
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}()

	return ch
}
