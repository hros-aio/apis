package permissions

type UpsertPermissionInput struct {
	Name   string `json:"name" validate:"required"`
	Module string `json:"module" validate:"required"`
}

func (data UpsertPermissionInput) Dto() *PermissionModel {
	return &PermissionModel{
		Name:   data.Name,
		Module: data.Module,
	}
}
