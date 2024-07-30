package handler

import (
	"fmt"
	"ikurotime/ideasai/internal/application/query"
	q "ikurotime/ideasai/internal/application/query"
	"ikurotime/ideasai/internal/domain/repositories"
	pkg "ikurotime/ideasai/pkg/query"

	"reflect"

	"go.uber.org/zap"
)

type FindProjectByIdHandler struct {
	repository repositories.ProjectRepository
	log        zap.Logger
}

func NewFindProjectHandler(repository repositories.ProjectRepository, log zap.Logger) pkg.QueryHandler {
	return &FindProjectByIdHandler{
		repository: repository,
		log:        log,
	}
}

func (h *FindProjectByIdHandler) Handle(query pkg.Query) (pkg.QueryResponse, error) {
	findProjectQuery, ok := query.(q.FindProjectQuery)
	if !ok {
		h.log.Error("FindProjectHandler.Handle", zap.String("error", "invalid query type"))
		return nil, fmt.Errorf("invalid query type")
	}
	h.log.Info("FindProjectHandler.Handle", zap.String("id", findProjectQuery.ID))
	project, err := h.repository.FindById(findProjectQuery.ID)
	if err != nil {
		h.log.Error("FindProjectHandler.Handle", zap.Error(err))
		return nil, err
	}
	h.log.Info("FindProjectHandler.Handle", zap.Any("project", project))
	return project, nil
}

func (h *FindProjectByIdHandler) QueryType() reflect.Type {
	return reflect.TypeOf(query.FindProjectQuery{})
}
