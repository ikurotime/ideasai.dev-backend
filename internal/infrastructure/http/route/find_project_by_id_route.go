package http

import (
	"ikurotime/ideasai/internal/application/query"
	pkg "ikurotime/ideasai/pkg/query"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FindProjectByIDRoute struct {
	queryBus pkg.QueryBus
	log      *zap.Logger
}

func NewFindProjectByIDRoute(log *zap.Logger) *FindProjectByIDRoute {
	return &FindProjectByIDRoute{log: log}
}

func (h *FindProjectByIDRoute) Method() string {
	return http.MethodGet
}

func (h *FindProjectByIDRoute) Pattern() string {
	return "/project/{id}"
}

func (h *FindProjectByIDRoute) Handle(ctx *gin.Context) {
	var q query.FindProjectQuery
	id := ctx.Param("id")
	h.log.Info("FindProjectByIDRoute.ServeHTTP", zap.String("id", id))

	q.ID = id

	response, err := h.queryBus.Execute(&q)

	if err != nil {
		h.log.Error("Failed to execute query", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, response)
}
