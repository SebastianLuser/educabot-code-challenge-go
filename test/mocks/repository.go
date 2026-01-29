package mocks

import (
	"context"

	"educabot.com/bookshop/models"
)

type MockBookRepository struct {
	Books []models.Book
	Err   error
}

func NewMockBookRepository() *MockBookRepository {
	return &MockBookRepository{}
}

func (m *MockBookRepository) WithBooks(books []models.Book) *MockBookRepository {
	m.Books = books
	return m
}

func (m *MockBookRepository) WithError(err error) *MockBookRepository {
	m.Err = err
	return m
}

func (m *MockBookRepository) GetBooks(_ context.Context) ([]models.Book, error) {
	return m.Books, m.Err
}
