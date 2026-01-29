package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"educabot.com/bookshop/models"
	"educabot.com/bookshop/service"
	"educabot.com/bookshop/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

const (
	testAuthorTolkien  = "Tolkien"
	testAuthorUnknown  = "Unknown"
	testBookLion       = "The Lion, the Witch and the Wardrobe"
	testMeanUnitsSold  = uint(53750000)
	testBooksCount     = uint(3)
	testCheapestPrice  = uint(15)

	pathMeanUnitsSold = "/books/mean-units-sold"
	pathCheapest      = "/books/cheapest"
	pathCountByAuthor = "/books/count-by-author/"

	keyMeanUnitsSold = "mean_units_sold"
	keyCount         = "count"
	keyName          = "name"
	keyPrice         = "price"
)

func setupRouter(h MetricsHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/books/mean-units-sold", h.GetMeanUnitsSold)
	r.GET("/books/cheapest", h.GetCheapestBook)
	r.GET("/books/count-by-author/:author", h.GetBooksCountByAuthor)
	return r
}

func TestGetMeanUnitsSold_Success(t *testing.T) {
	mockSvc := mocks.NewMockMetricsService().WithMeanUnitsSold(testMeanUnitsSold)
	handler := NewMetricsHandler(mockSvc)
	router := setupRouter(handler)
	req := httptest.NewRequest(http.MethodGet, pathMeanUnitsSold, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Equal(t, float64(testMeanUnitsSold), response[keyMeanUnitsSold])
}

func TestGetMeanUnitsSold_NoBooksFound(t *testing.T) {
	mockSvc := mocks.NewMockMetricsService().WithError(service.ErrNoBooksFound)
	handler := NewMetricsHandler(mockSvc)
	router := setupRouter(handler)
	req := httptest.NewRequest(http.MethodGet, pathMeanUnitsSold, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetMeanUnitsSold_FetchingError(t *testing.T) {
	mockSvc := mocks.NewMockMetricsService().WithError(service.ErrFetchingBooks)
	handler := NewMetricsHandler(mockSvc)
	router := setupRouter(handler)
	req := httptest.NewRequest(http.MethodGet, pathMeanUnitsSold, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadGateway, rec.Code)
}

func TestGetCheapestBook_Success(t *testing.T) {
	mockSvc := mocks.NewMockMetricsService().WithCheapestBook(models.Book{Name: testBookLion, Price: testCheapestPrice})
	handler := NewMetricsHandler(mockSvc)
	router := setupRouter(handler)
	req := httptest.NewRequest(http.MethodGet, pathCheapest, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Equal(t, testBookLion, response[keyName])
	require.Equal(t, float64(testCheapestPrice), response[keyPrice])
}

func TestGetCheapestBook_NoBooksFound(t *testing.T) {
	mockSvc := mocks.NewMockMetricsService().WithError(service.ErrNoBooksFound)
	handler := NewMetricsHandler(mockSvc)
	router := setupRouter(handler)
	req := httptest.NewRequest(http.MethodGet, pathCheapest, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetCheapestBook_FetchingError(t *testing.T) {
	mockSvc := mocks.NewMockMetricsService().WithError(service.ErrFetchingBooks)
	handler := NewMetricsHandler(mockSvc)
	router := setupRouter(handler)
	req := httptest.NewRequest(http.MethodGet, pathCheapest, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadGateway, rec.Code)
}

func TestGetBooksCountByAuthor_Success(t *testing.T) {
	mockSvc := mocks.NewMockMetricsService().WithBooksCount(testBooksCount)
	handler := NewMetricsHandler(mockSvc)
	router := setupRouter(handler)
	req := httptest.NewRequest(http.MethodGet, pathCountByAuthor+testAuthorTolkien, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Equal(t, float64(testBooksCount), response[keyCount])
}

func TestGetBooksCountByAuthor_AuthorNotFound(t *testing.T) {
	mockSvc := mocks.NewMockMetricsService().WithError(service.ErrAuthorNotFound)
	handler := NewMetricsHandler(mockSvc)
	router := setupRouter(handler)
	req := httptest.NewRequest(http.MethodGet, pathCountByAuthor+testAuthorUnknown, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetBooksCountByAuthor_NoBooksFound(t *testing.T) {
	mockSvc := mocks.NewMockMetricsService().WithError(service.ErrNoBooksFound)
	handler := NewMetricsHandler(mockSvc)
	router := setupRouter(handler)
	req := httptest.NewRequest(http.MethodGet, pathCountByAuthor+testAuthorTolkien, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetBooksCountByAuthor_FetchingError(t *testing.T) {
	mockSvc := mocks.NewMockMetricsService().WithError(service.ErrFetchingBooks)
	handler := NewMetricsHandler(mockSvc)
	router := setupRouter(handler)
	req := httptest.NewRequest(http.MethodGet, pathCountByAuthor+testAuthorTolkien, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadGateway, rec.Code)
}
