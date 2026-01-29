package service

import (
	"context"
	"fmt"
	"slices"

	"educabot.com/bookshop/models"
	"educabot.com/bookshop/repository"
)

type (
	metricsService struct {
		bookRepo repository.BookRepository
	}

	MetricsService interface {
		GetMeanUnitsSold(ctx context.Context) (uint, error)
		GetCheapestBook(ctx context.Context) (models.Book, error)
		GetBooksCountByAuthor(ctx context.Context, author string) (uint, error)
	}
)

func NewMetricsService(bookRepo repository.BookRepository) MetricsService {
	return &metricsService{bookRepo: bookRepo}
}

func (s *metricsService) GetMeanUnitsSold(ctx context.Context) (uint, error) {
	books, err := s.bookRepo.GetBooks(ctx)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrFetchingBooks, err)
	}
	if len(books) == 0 {
		return 0, ErrNoBooksFound
	}
	return meanUnitsSold(books), nil
}

func (s *metricsService) GetCheapestBook(ctx context.Context) (models.Book, error) {
	books, err := s.bookRepo.GetBooks(ctx)
	if err != nil {
		return models.Book{}, fmt.Errorf("%w: %w", ErrFetchingBooks, err)
	}
	if len(books) == 0 {
		return models.Book{}, ErrNoBooksFound
	}
	return cheapestBook(books), nil
}

func (s *metricsService) GetBooksCountByAuthor(ctx context.Context, author string) (uint, error) {
	books, err := s.bookRepo.GetBooks(ctx)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrFetchingBooks, err)
	}
	if len(books) == 0 {
		return 0, ErrNoBooksFound
	}
	count := booksCountByAuthor(books, author)
	if count == 0 {
		return 0, ErrAuthorNotFound
	}
	return count, nil
}

func meanUnitsSold(books []models.Book) uint {
	var sum uint
	for _, book := range books {
		sum += book.UnitsSold
	}
	return sum / uint(len(books))
}

func cheapestBook(books []models.Book) models.Book {
	return slices.MinFunc(books, func(a, b models.Book) int {
		return int(a.Price - b.Price)
	})
}

func booksCountByAuthor(books []models.Book, author string) uint {
	var count uint
	for _, book := range books {
		if book.Author == author {
			count++
		}
	}
	return count
}
