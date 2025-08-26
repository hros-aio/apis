package locations

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/hros-aio/apis/libs/saga/messages"
)

func ToModel(msg messages.LocationCreatedPayload) *location.LocationModel {
	return &location.LocationModel{
		Model: base.Model{
			ID: uuid.MustParse(msg.ID),
		},
		Name:          msg.Name,
		TenantId:      msg.TenantID,
		CompanyID:     uuid.MustParse(msg.CompanyID),
		MapUrl:        msg.MapUrl,
		IsHeadquarter: msg.IsHeadquarter,
		AddressInfo:   location.AddressInfo(msg.AddressInfo),
	}
}
