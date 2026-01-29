package repository

import "errors"

var (
	ErrCreatingRequest    = errors.New("creating request")
	ErrExecutingRequest   = errors.New("executing request")
	ErrUnexpectedStatus   = errors.New("unexpected status code")
	ErrDecodingResponse   = errors.New("decoding response")
)
