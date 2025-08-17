package departments

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/department"
)

type CreateDepartmentInput struct {
	Name       string     `json:"name" example:"Department 1" validate:"required"`
	IsDivision bool       `json:"isDivision" example:"false" validate:"required"`
	ParentID   *uuid.UUID `json:"parentId" example:"6cdad833-ba6d-49e3-889c-da23b764bb21" validate:"required"`
}

func (data CreateDepartmentInput) Dto() *department.DepartmentModel {
	return &department.DepartmentModel{
		Name:       data.Name,
		IsDivision: data.IsDivision,
		ParentID:   data.ParentID,
	}
}

type UpdateDepartmentInput struct {
	Name       string     `json:"name" example:"Department 1"`
	IsDivision bool       `json:"isDivision" example:"false"`
	ParentID   *uuid.UUID `json:"parentId" example:"6cdad833-ba6d-49e3-889c-da23b764bb21"`
}

func (data UpdateDepartmentInput) Dto() *department.DepartmentModel {
	return &department.DepartmentModel{
		Name:       data.Name,
		IsDivision: data.IsDivision,
		ParentID:   data.ParentID,
	}
}
