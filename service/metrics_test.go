package service

import (
	"context"
	"errors"
	"testing"

	"educabot.com/bookshop/models"
	"educabot.com/bookshop/test/mocks"
	"github.com/stretchr/testify/require"
)

var (
	errRepository = errors.New("repository error")

	testAuthorTolkien  = "J.R.R. Tolkien"
	testAuthorLewis    = "C.S. Lewis"
	testAuthorUnknown  = "Unknown Author"
	testBookFellowship = "The Fellowship of the Ring"
	testBookLion       = "The Lion, the Witch and the Wardrobe"
)

func newTestBooks() []models.Book {
	return []models.Book{
		{ID: 1, Name: testBookFellowship, Author: testAuthorTolkien, UnitsSold: 50000000, Price: 20},
		{ID: 2, Name: "The Two Towers", Author: testAuthorTolkien, UnitsSold: 30000000, Price: 20},
		{ID: 3, Name: "The Return of the King", Author: testAuthorTolkien, UnitsSold: 50000000, Price: 20},
		{ID: 4, Name: testBookLion, Author: testAuthorLewis, UnitsSold: 85000000, Price: 15},
	}
}

func TestGetMeanUnitsSold_Success(t *testing.T) {
	repo := mocks.NewMockBookRepository().WithBooks(newTestBooks())
	svc := NewMetricsService(repo)

	result, err := svc.GetMeanUnitsSold(context.Background())

	require.NoError(t, err)
	require.Equal(t, uint(53750000), result)
}

func TestGetMeanUnitsSold_RepositoryError(t *testing.T) {
	repo := mocks.NewMockBookRepository().WithError(errRepository)
	svc := NewMetricsService(repo)

	_, err := svc.GetMeanUnitsSold(context.Background())

	require.ErrorIs(t, err, ErrFetchingBooks)
	require.ErrorIs(t, err, errRepository)
}

func TestGetMeanUnitsSold_NoBooksFound(t *testing.T) {
	repo := mocks.NewMockBookRepository().WithBooks([]models.Book{})
	svc := NewMetricsService(repo)

	_, err := svc.GetMeanUnitsSold(context.Background())

	require.ErrorIs(t, err, ErrNoBooksFound)
}

func TestGetCheapestBook_Success(t *testing.T) {
	repo := mocks.NewMockBookRepository().WithBooks(newTestBooks())
	svc := NewMetricsService(repo)

	result, err := svc.GetCheapestBook(context.Background())

	require.NoError(t, err)
	require.Equal(t, testBookLion, result.Name)
	require.Equal(t, uint(15), result.Price)
}

func TestGetCheapestBook_RepositoryError(t *testing.T) {
	repo := mocks.NewMockBookRepository().WithError(errRepository)
	svc := NewMetricsService(repo)

	_, err := svc.GetCheapestBook(context.Background())

	require.ErrorIs(t, err, ErrFetchingBooks)
	require.ErrorIs(t, err, errRepository)
}

func TestGetCheapestBook_NoBooksFound(t *testing.T) {
	repo := mocks.NewMockBookRepository().WithBooks([]models.Book{})
	svc := NewMetricsService(repo)

	_, err := svc.GetCheapestBook(context.Background())

	require.ErrorIs(t, err, ErrNoBooksFound)
}

func TestGetBooksCountByAuthor_Success(t *testing.T) {
	repo := mocks.NewMockBookRepository().WithBooks(newTestBooks())
	svc := NewMetricsService(repo)

	result, err := svc.GetBooksCountByAuthor(context.Background(), testAuthorTolkien)

	require.NoError(t, err)
	require.Equal(t, uint(3), result)
}

func TestGetBooksCountByAuthor_SingleBook(t *testing.T) {
	repo := mocks.NewMockBookRepository().WithBooks(newTestBooks())
	svc := NewMetricsService(repo)

	result, err := svc.GetBooksCountByAuthor(context.Background(), testAuthorLewis)

	require.NoError(t, err)
	require.Equal(t, uint(1), result)
}

func TestGetBooksCountByAuthor_RepositoryError(t *testing.T) {
	repo := mocks.NewMockBookRepository().WithError(errRepository)
	svc := NewMetricsService(repo)

	_, err := svc.GetBooksCountByAuthor(context.Background(), testAuthorTolkien)

	require.ErrorIs(t, err, ErrFetchingBooks)
	require.ErrorIs(t, err, errRepository)
}

func TestGetBooksCountByAuthor_NoBooksFound(t *testing.T) {
	repo := mocks.NewMockBookRepository().WithBooks([]models.Book{})
	svc := NewMetricsService(repo)

	_, err := svc.GetBooksCountByAuthor(context.Background(), testAuthorTolkien)

	require.ErrorIs(t, err, ErrNoBooksFound)
}

func TestGetBooksCountByAuthor_AuthorNotFound(t *testing.T) {
	repo := mocks.NewMockBookRepository().WithBooks(newTestBooks())
	svc := NewMetricsService(repo)

	_, err := svc.GetBooksCountByAuthor(context.Background(), testAuthorUnknown)

	require.ErrorIs(t, err, ErrAuthorNotFound)
}
