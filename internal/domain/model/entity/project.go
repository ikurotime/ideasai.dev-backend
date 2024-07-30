package entity

import vo "ikurotime/ideasai/internal/domain/model/value_object"

type Project struct {
	ID          vo.ProjectID
	Description vo.Description
	Category    vo.Category
}
