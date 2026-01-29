package main

import (
	"educabot.com/bookshop/handler"
	"educabot.com/bookshop/service"
)

func newMetricsHandler(metricsSvc service.MetricsService) handler.MetricsHandler {
	return handler.NewMetricsHandler(metricsSvc)
}
