package http

import (
	"fmt"
	"ikurotime/ideasai/internal/application/query"
	pkg "ikurotime/ideasai/pkg/query"
	"net/http"

	"go.uber.org/zap"
)

type HelloHandler struct {
	queryBus pkg.QueryBus
	log      *zap.Logger
}

func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{log: log}
}
func (*HelloHandler) Pattern() string {
	return "/project/{id}"
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var q query.FindProjectQuery
	id := r.PathValue("id")
	h.log.Info("HelloHandler.ServeHTTP", zap.String("id", id))

	q.ID = id

	response, err := h.queryBus.Execute(&q)

	if err != nil {
		h.log.Error("Failed to execute query", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	if _, err := fmt.Fprintf(w, "%s\n", response); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
