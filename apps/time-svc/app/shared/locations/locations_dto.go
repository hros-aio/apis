package locations

import (
	"github.com/hros-aio/apis/libs/mongodoc/common/location"
	"github.com/hros-aio/apis/libs/saga/messages"
)

func ToModel(data messages.LocationCreatedPayload) *location.LocationModel {
	return &location.LocationModel{
		AddressInfo:   location.AddressInfo(data.AddressInfo),
		Name:          data.Name,
		TenantID:      data.TenantID,
		CompanyID:     data.CompanyID,
		MapUrl:        data.MapUrl,
		SyncID:        data.ID,
		IsHeadquarter: data.IsHeadquarter,
	}
}
