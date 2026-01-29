package handler

import (
	"errors"
	"net/http"

	"educabot.com/bookshop/service"
)

func mapErrorToHTTPStatus(err error) int {
	switch {
	case errors.Is(err, service.ErrNoBooksFound):
		return http.StatusNotFound
	case errors.Is(err, service.ErrAuthorNotFound):
		return http.StatusNotFound
	case errors.Is(err, service.ErrFetchingBooks):
		return http.StatusBadGateway
	default:
		return http.StatusInternalServerError
	}
}
