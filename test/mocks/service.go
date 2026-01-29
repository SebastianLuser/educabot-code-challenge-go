package mocks

import (
	"context"

	"educabot.com/bookshop/models"
)

type MockMetricsService struct {
	MeanUnitsSold uint
	CheapestBook  models.Book
	BooksCount    uint
	Err           error
}

func NewMockMetricsService() *MockMetricsService {
	return &MockMetricsService{}
}

func (m *MockMetricsService) WithMeanUnitsSold(mean uint) *MockMetricsService {
	m.MeanUnitsSold = mean
	return m
}

func (m *MockMetricsService) WithCheapestBook(book models.Book) *MockMetricsService {
	m.CheapestBook = book
	return m
}

func (m *MockMetricsService) WithBooksCount(count uint) *MockMetricsService {
	m.BooksCount = count
	return m
}

func (m *MockMetricsService) WithError(err error) *MockMetricsService {
	m.Err = err
	return m
}

func (m *MockMetricsService) GetMeanUnitsSold(_ context.Context) (uint, error) {
	return m.MeanUnitsSold, m.Err
}

func (m *MockMetricsService) GetCheapestBook(_ context.Context) (models.Book, error) {
	return m.CheapestBook, m.Err
}

func (m *MockMetricsService) GetBooksCountByAuthor(_ context.Context, _ string) (uint, error) {
	return m.BooksCount, m.Err
}
