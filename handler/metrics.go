package handler

import (
	"net/http"

	"educabot.com/bookshop/service"
	"github.com/gin-gonic/gin"
)

const errorKey = "error"

type (
	metricsHandler struct {
		metricsService service.MetricsService
	}

	MetricsHandler interface {
		GetMeanUnitsSold(ctx *gin.Context)
		GetCheapestBook(ctx *gin.Context)
		GetBooksCountByAuthor(ctx *gin.Context)
	}
)

func NewMetricsHandler(metricsService service.MetricsService) MetricsHandler {
	return &metricsHandler{metricsService: metricsService}
}

func (h *metricsHandler) GetMeanUnitsSold(ctx *gin.Context) {
	mean, err := h.metricsService.GetMeanUnitsSold(ctx.Request.Context())
	if err != nil {
		ctx.JSON(mapErrorToHTTPStatus(err), gin.H{errorKey: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"mean_units_sold": mean})
}

func (h *metricsHandler) GetCheapestBook(ctx *gin.Context) {
	book, err := h.metricsService.GetCheapestBook(ctx.Request.Context())
	if err != nil {
		ctx.JSON(mapErrorToHTTPStatus(err), gin.H{errorKey: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (h *metricsHandler) GetBooksCountByAuthor(ctx *gin.Context) {
	author := ctx.Param("author")

	count, err := h.metricsService.GetBooksCountByAuthor(ctx.Request.Context(), author)
	if err != nil {
		ctx.JSON(mapErrorToHTTPStatus(err), gin.H{errorKey: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"count": count})
}
