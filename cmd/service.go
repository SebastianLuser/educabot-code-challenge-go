package main

import (
	"educabot.com/bookshop/repository"
	"educabot.com/bookshop/service"
)

func newMetricsService(bookRepo repository.BookRepository) service.MetricsService {
	return service.NewMetricsService(bookRepo)
}
