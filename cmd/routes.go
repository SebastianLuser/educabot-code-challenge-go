package main

import (
	"educabot.com/bookshop/handler"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine, metricsHandler handler.MetricsHandler) {
	books := router.Group("/books")
	{
		books.GET("/mean-units-sold", metricsHandler.GetMeanUnitsSold)
		books.GET("/cheapest", metricsHandler.GetCheapestBook)
		books.GET("/count-by-author/:author", metricsHandler.GetBooksCountByAuthor)
	}
}