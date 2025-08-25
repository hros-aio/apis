package permissions

import (
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/tinh-tinh/sqlorm/v2"
)

type PermissionModel struct {
	base.Model
	Name   string `json:"name"`
	Module string `json:"module"`
}

func (model PermissionModel) DataMapper() *PermissionDB {
	return &PermissionDB{
		Model:  sqlorm.Model(model.Model),
		Name:   model.Name,
		Module: model.Module,
	}
}
