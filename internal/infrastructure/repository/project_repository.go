package repository

import (
	"ikurotime/ideasai/internal/domain/repositories"

	"ikurotime/ideasai/internal/domain/model/entity"

	"go.uber.org/zap"
)

type ProjectRepository struct {
	log *zap.Logger
}

func NewProjectRepository(log *zap.Logger) repositories.ProjectRepository {
	return &ProjectRepository{
		log: log,
	}
}

func (r *ProjectRepository) FindAll() ([]*entity.Project, error) {
	r.log.Info("ProjectRepository.FindAll")
	return nil, nil
}

func (r *ProjectRepository) FindById(id string) (*entity.Project, error) {
	r.log.Info("ProjectRepository.FindById", zap.String("id", id))
	return nil, nil
}

func (r *ProjectRepository) Save(project *entity.Project) error {
	r.log.Info("ProjectRepository.Save", zap.Any("project", project))
	return nil
}
func (r *ProjectRepository) Update(project *entity.Project) error {
	r.log.Info("ProjectRepository.Update", zap.Any("project", project))
	return nil
}
