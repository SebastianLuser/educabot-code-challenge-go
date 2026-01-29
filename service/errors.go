package service

import "errors"

var (
	ErrNoBooksFound    = errors.New("no books found")
	ErrAuthorNotFound  = errors.New("author not found")
	ErrFetchingBooks   = errors.New("fetching books")
)
