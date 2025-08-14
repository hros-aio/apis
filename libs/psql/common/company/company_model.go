package company

import (
	"reflect"

	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
	"github.com/tinh-tinh/sqlorm/v2"
)

type CompanyModel struct {
	base.Model
	TenantId         string               `json:"tenantId"`
	Name             string               `json:"name"`
	Industry         string               `json:"industry"`
	Size             int                  `json:"size"`
	Logo             string               `json:"logo"`
	Contact          tenant.ContactPerson `json:"contact"`
	SecondaryContact tenant.ContactPerson `json:"secondaryContact"`
}

func (model CompanyModel) DataMapper() *CompanyDB {
	data := &CompanyDB{
		TenantId:         model.TenantId,
		Name:             model.Name,
		Industry:         model.Industry,
		Size:             model.Size,
		Logo:             model.Logo,
		Contact:          tenant.ContactPersonDb(model.Contact),
		SecondaryContact: tenant.ContactPersonDb(model.SecondaryContact),
	}

	if !reflect.ValueOf(model.Model).IsZero() {
		data.Model = sqlorm.Model(model.Model)
	}

	return data
}
