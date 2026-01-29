package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()
	router.SetTrustedProxies(nil)

	bookRepo := newBookRepository()
	metricsSvc := newMetricsService(bookRepo)
	metricsHandler := newMetricsHandler(metricsSvc)

	setupRoutes(router, metricsHandler)
	router.Run(":3000")
}
