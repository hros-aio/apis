package locations

import (
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
)

type CreateLocationInput struct {
	Name     string                `json:"name" example:"Location 1" validate:"required"`
	Line     string                `json:"addressLine" example:"203 Street" validate:"required"`
	City     string                `json:"city" example:"MN City" validate:"required"`
	State    string                `json:"state" example:"Manu" validate:"required"`
	Country  string                `json:"country" example:"USA" validate:"required"`
	Zipcode  string                `json:"zipCode" example:"70000" validate:"required"`
	Timezone string                `json:"timezone" example:"PST" validate:"required"`
	MapUrl   string                `json:"mapUrl" example:"https://abc.com" validate:"required"`
	Contact  *tenant.ContactPerson `json:"contact" validate:"nested"`
}

func (data CreateLocationInput) Dto() *location.LocationModel {
	model := &location.LocationModel{
		Name: data.Name,
		AddressInfo: location.AddressInfo{
			Line:     data.Line,
			City:     data.City,
			State:    data.State,
			Country:  data.Country,
			Timezone: data.Timezone,
			Zipcode:  data.Zipcode,
		},
		MapUrl: data.MapUrl,
	}
	if data.Contact != nil {
		model.Contact = *data.Contact
	}

	return model
}
