package grades

import "github.com/hros-aio/apis/libs/factory/middleware"

type CreateGradeInput struct {
	Name string `json:"name" validate:"required" example:"Grade 1"`
}

func (data CreateGradeInput) Dto(ctx middleware.ContextInfo) *GradeModel {
	return &GradeModel{
		TenantID:  ctx.TenantID,
		CompanyID: ctx.CompanyID,
		Name:      data.Name,
	}
}

type UpdateGradeInput struct {
	Name string `json:"name" example:"Grade 1"`
}

func (data UpdateGradeInput) Dto() *GradeModel {
	return &GradeModel{
		Name: data.Name,
	}
}
