package repositories

import . "ikurotime/ideasai/internal/domain/model/entity"

type ProjectRepository interface {
	FindAll() ([]*Project, error)
	FindById(id string) (*Project, error)
	Save(project *Project) error
	Update(project *Project) error
}
