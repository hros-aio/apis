package location

import (
	"github.com/hros-aio/apis/libs/mongodoc/common/base"
	"github.com/tinh-tinh/mongoose/v2"
)

type AddressInfoSchema struct {
	Line     string `bson:"addressLine" json:"addressLine"`
	City     string `bson:"city" json:"city"`
	State    string `bson:"state" json:"state"`
	Country  string `bson:"country" json:"country"`
	Zipcode  string `bson:"zipCode" json:"zipCode"`
	Timezone string `bson:"timezone" json:"timezone"`
}

type LocationSchema struct {
	mongoose.BaseSchema `bson:"inline"`
	SyncID              string            `bson:"syncId"`
	AddressInfo         AddressInfoSchema `bson:"addressInfo"`
	TenantID            string            `bson:"tenantId" json:"tenantId"`
	CompanyID           string            `bson:"companyId" json:"companyId"`
	Name                string            `bson:"name" json:"name"`
	MapUrl              string            `bson:"mapUrl" json:"mapUrl"`
}

func (LocationSchema) CollectionName() string {
	return "locations"
}

func (data LocationSchema) Dto() *LocationModel {
	model := &LocationModel{
		Model:       base.Model(data.BaseSchema),
		AddressInfo: AddressInfo(data.AddressInfo),
		SyncID:      data.SyncID,
		TenantID:    data.TenantID,
		CompanyID:   data.CompanyID,
		Name:        data.Name,
		MapUrl:      data.MapUrl,
	}

	return model
}
