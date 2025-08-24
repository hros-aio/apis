package grades

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/tinh-tinh/sqlorm/v2"
)

type GradeModel struct {
	base.Model
	TenantID  string                `json:"tenantId"`
	CompanyID uuid.UUID             `json:"companyId"`
	Company   *company.CompanyModel `json:"company,omitempty"`
	Code      string                `json:"code"`
	Name      string                `json:"name"`
}

func (model GradeModel) DataMapper() *GradeDB {
	data := &GradeDB{
		TenantID:  model.TenantID,
		Name:      model.Name,
		Code:      model.Code,
		CompanyID: model.CompanyID,
	}

	if !reflect.ValueOf(model.Model).IsZero() {
		data.Model = sqlorm.Model(model.Model)
	}

	return data
}
