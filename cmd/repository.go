package main

import (
	"net/http"

	"educabot.com/bookshop/repository"
)

func newBookRepository() repository.BookRepository {
	return repository.NewHTTPBookRepository(&http.Client{})
}
