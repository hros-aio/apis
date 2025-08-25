package roles

import "github.com/hros-aio/apis/libs/factory/middleware"

type CreateRoleInput struct {
	Name        string `json:"name" validate:"required" example:"Admin"`
	Description string `json:"description" example:"Administrator role"`
}

func (data CreateRoleInput) Dto(ctx middleware.ContextInfo) *RoleModel {
	return &RoleModel{
		TenantID:    ctx.TenantID,
		CompanyID:   ctx.CompanyID,
		Name:        data.Name,
		Description: data.Description,
	}
}

type UpdateRoleInput struct {
	Name        string `json:"name" example:"Admin"`
	Description string `json:"description" example:"Administrator role"`
}

func (data UpdateRoleInput) Dto() *RoleModel {
	return &RoleModel{
		Name:        data.Name,
		Description: data.Description,
	}
}
