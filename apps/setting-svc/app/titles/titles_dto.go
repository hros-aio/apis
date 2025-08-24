package titles

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/factory/middleware"
)

type CreateTitleInput struct {
	DepartmentID string `json:"departmentId" validate:"required,isUUID" example:"6cdad833-ba6d-49e3-889c-da23b764bb21"`
	GradeID      string `json:"gradeId" validate:"required,isUUID" example:"6cdad833-ba6d-49e3-889c-da23b764bb21"`
	Name         string `json:"name" validate:"required" example:"Senior Developer"`
}

func (data CreateTitleInput) Dto(ctx middleware.ContextInfo) *TitleModel {
	return &TitleModel{
		TenantID:     ctx.TenantID,
		CompanyID:    ctx.CompanyID,
		DepartmentID: uuid.MustParse(data.DepartmentID),
		GradeID:      uuid.MustParse(data.GradeID),
		Name:         data.Name,
	}
}

type UpdateTitleInput struct {
	DepartmentID string `json:"departmentId" validate:"isUUID" example:"6cdad833-ba6d-49e3-889c-da23b764bb21"`
	GradeID      string `json:"gradeId" validate:"isUUID" example:"6cdad833-ba6d-49e3-889c-da23b764bb21"`
	Name         string `json:"name" example:"Senior Developer"`
}

func (data UpdateTitleInput) Dto() *TitleModel {
	return &TitleModel{
		DepartmentID: uuid.MustParse(data.DepartmentID),
		GradeID:      uuid.MustParse(data.GradeID),
		Name:         data.Name,
	}
}
