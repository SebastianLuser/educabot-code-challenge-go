package repository

import (
	"context"

	"educabot.com/bookshop/models"
)

type BookRepository interface {
	GetBooks(ctx context.Context) ([]models.Book, error)
}
