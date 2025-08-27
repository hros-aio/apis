package departments

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql/common/department"
	"github.com/hros-aio/apis/libs/saga/messages"
)

type CreateDepartmentInput struct {
	Name       string     `json:"name" example:"Department 1" validate:"required"`
	IsDivision bool       `json:"isDivision" example:"false" validate:"required"`
	ParentID   *uuid.UUID `json:"parentId" example:"6cdad833-ba6d-49e3-889c-da23b764bb21" validate:"required"`
}

func (data CreateDepartmentInput) Dto(ctx middleware.ContextInfo) *department.DepartmentModel {
	return &department.DepartmentModel{
		Name:       data.Name,
		IsDivision: data.IsDivision,
		ParentID:   data.ParentID,
		TenantID:   ctx.TenantID,
		CompanyID:  ctx.CompanyID,
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

func ToCreatedMessage(model *department.DepartmentModel) messages.DepartmentCreatedPayload {
	msg := messages.DepartmentCreatedPayload{
		ID:         model.ID.String(),
		Name:       model.Name,
		IsDivision: model.IsDivision,
		TenantID:   model.TenantID,
		CompanyID:  model.CompanyID.String(),
		Code:       model.Code,
	}

	if model.ParentID != nil {
		msg.ParentID = model.ParentID.String()
	}

	return msg
}

func ToUpdatedMessage(oldData *department.DepartmentModel, newData *department.DepartmentModel) messages.DepartmentUpdatedPayload {
	msg := messages.DepartmentUpdatedPayload{
		PreviouseData: ToCreatedMessage(oldData),
		Data:          ToCreatedMessage(newData),
	}

	return msg
}
