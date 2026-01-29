package repository

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	validBooksJSON = `[{"id":1,"name":"The Fellowship of the Ring","author":"J.R.R. Tolkien","units_sold":50000000,"price":20}]`
	invalidJSON    = `{"invalid`

	expectedBookName   = "The Fellowship of the Ring"
	expectedBookAuthor = "J.R.R. Tolkien"
	expectedUnitsSold  = uint(50000000)
	expectedPrice      = uint(20)
)

func TestGetBooks_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(validBooksJSON))
	}))
	defer server.Close()
	repo := &HTTPBookRepository{client: server.Client()}
	originalURL := booksAPIURL
	defer func() { booksAPIURL = originalURL }()
	booksAPIURL = server.URL

	books, err := repo.GetBooks(context.Background())

	require.NoError(t, err)
	require.Len(t, books, 1)
	require.Equal(t, expectedBookName, books[0].Name)
	require.Equal(t, expectedBookAuthor, books[0].Author)
	require.Equal(t, expectedUnitsSold, books[0].UnitsSold)
	require.Equal(t, expectedPrice, books[0].Price)
}

func TestGetBooks_UnexpectedStatusCode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	repo := &HTTPBookRepository{client: server.Client()}
	originalURL := booksAPIURL
	defer func() { booksAPIURL = originalURL }()
	booksAPIURL = server.URL

	_, err := repo.GetBooks(context.Background())

	require.ErrorIs(t, err, ErrUnexpectedStatus)
}

func TestGetBooks_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(invalidJSON))
	}))
	defer server.Close()
	repo := &HTTPBookRepository{client: server.Client()}
	originalURL := booksAPIURL
	defer func() { booksAPIURL = originalURL }()
	booksAPIURL = server.URL

	_, err := repo.GetBooks(context.Background())

	require.ErrorIs(t, err, ErrDecodingResponse)
}

func TestGetBooks_RequestError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	server.Close()
	repo := &HTTPBookRepository{client: server.Client()}
	originalURL := booksAPIURL
	defer func() { booksAPIURL = originalURL }()
	booksAPIURL = server.URL

	_, err := repo.GetBooks(context.Background())

	require.ErrorIs(t, err, ErrExecutingRequest)
}

func TestGetBooks_ContextCanceled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(validBooksJSON))
	}))
	defer server.Close()
	repo := &HTTPBookRepository{client: server.Client()}
	originalURL := booksAPIURL
	defer func() { booksAPIURL = originalURL }()
	booksAPIURL = server.URL
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := repo.GetBooks(ctx)

	require.ErrorIs(t, err, ErrExecutingRequest)
}

func TestGetBooks_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[]"))
	}))
	defer server.Close()
	repo := &HTTPBookRepository{client: server.Client()}
	originalURL := booksAPIURL
	defer func() { booksAPIURL = originalURL }()
	booksAPIURL = server.URL

	books, err := repo.GetBooks(context.Background())

	require.NoError(t, err)
	require.Empty(t, books)
}
