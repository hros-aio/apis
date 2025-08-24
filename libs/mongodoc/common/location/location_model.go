package location

import "github.com/hros-aio/apis/libs/mongodoc/common/base"

type AddressInfo struct {
	Line     string `json:"addressLine"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Zipcode  string `json:"zipCode"`
	Timezone string `json:"timezone"`
}

type LocationModel struct {
	base.Model
	AddressInfo
	Name          string `json:"name"`
	TenantID      string `json:"tenantId"`
	CompanyID     string `json:"companyId"`
	SyncID        string `json:"syncId"`
	MapUrl        string `json:"mapUrl"`
	IsHeadquarter bool   `json:"isHeadquarter"`
}

func (model LocationModel) DataMapper() *LocationSchema {
	data := &LocationSchema{
		AddressInfo:   AddressInfoSchema(model.AddressInfo),
		TenantID:      model.TenantID,
		Name:          model.Name,
		CompanyID:     model.CompanyID,
		MapUrl:        model.MapUrl,
		IsHeadquarter: model.IsHeadquarter,
		SyncID:        model.SyncID,
	}

	return data
}
