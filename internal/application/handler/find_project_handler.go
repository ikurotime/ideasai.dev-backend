package handler

import (
	"ikurotime/ideasai/internal/application/query"
	"ikurotime/ideasai/internal/domain/repositories"
	"reflect"

	"go.uber.org/zap"
)

type FindProjectHandler struct {
	repo repositories.ProjectRepository
	log  zap.Logger
}

func NewFindProjectHandler(log zap.Logger) *FindProjectHandler {
	return &FindProjectHandler{
		log: log,
	}
}

func (h *FindProjectHandler) Handle(query query.FindProjectQuery) {
	h.log.Info("FindProjectHandler.Handle", zap.String("id", query.ID))
	project, err := h.repo.FindById(query.ID)
	if err != nil {
		h.log.Error("FindProjectHandler.Handle", zap.Error(err))
		return
	}
	h.log.Info("FindProjectHandler.Handle", zap.Any("project", project))
}

func (h *FindProjectHandler) QueryType() reflect.Type {
	return reflect.TypeOf(query.FindProjectQuery{})
}
