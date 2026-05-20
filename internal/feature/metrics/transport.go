package feature_metrics

import core_logger "MetricService/internal/core/logger"

type MetricCase struct {
	log *core_logger.Logger
}

func NewMetricCase(
	log *core_logger.Logger,
) *MetricCase {
	return &MetricCase{
		log: log,
	}
}
