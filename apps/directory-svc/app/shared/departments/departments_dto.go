package departments

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/department"
	"github.com/hros-aio/apis/libs/saga/messages"
)

func ToModel(msg messages.DepartmentCreatedPayload) *department.DepartmentModel {
	model := &department.DepartmentModel{
		Model: base.Model{
			ID: uuid.MustParse(msg.ID),
		}, Name: msg.Name,
		TenantID:   msg.TenantID,
		CompanyID:  uuid.MustParse(msg.CompanyID),
		Code:       msg.Code,
		IsDivision: msg.IsDivision,
	}

	if msg.ParentID != "" {
		parentId := uuid.MustParse(msg.ParentID)
		model.ParentID = &parentId
	}

	return model
}
