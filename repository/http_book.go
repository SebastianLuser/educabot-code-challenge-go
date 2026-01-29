package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"educabot.com/bookshop/models"
)

var booksAPIURL = "https://6781684b85151f714b0aa5db.mockapi.io/api/v1/books"

type HTTPBookRepository struct {
	client *http.Client
}

func NewHTTPBookRepository(client *http.Client) *HTTPBookRepository {
	return &HTTPBookRepository{client: client}
}

func (r *HTTPBookRepository) GetBooks(ctx context.Context) ([]models.Book, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, booksAPIURL, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCreatingRequest, err)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrExecutingRequest, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %d", ErrUnexpectedStatus, resp.StatusCode)
	}

	var books []models.Book
	if err := json.NewDecoder(resp.Body).Decode(&books); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodingResponse, err)
	}

	return books, nil
}
