package locations

import (
	"reflect"

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

type UpdateLocationInput struct {
	Name     *string               `json:"name,omitempty" example:"Location 1"`
	Line     *string               `json:"addressLine,omitempty" example:"203 Street"`
	City     *string               `json:"city,omitempty" example:"MN City"`
	State    *string               `json:"state,omitempty" example:"Manu"`
	Country  *string               `json:"country,omitempty" example:"USA"`
	Zipcode  *string               `json:"zipCode,omitempty" example:"70000"`
	Timezone *string               `json:"timezone,omitempty" example:"PST"`
	MapUrl   *string               `json:"mapUrl,omitempty" example:"https://abc.com"`
	Contact  *tenant.ContactPerson `json:"contact,omitempty" validate:"nested"`
}

func (data UpdateLocationInput) Dto() *location.LocationModel {
	model := &location.LocationModel{}
	adddress := location.AddressInfo{}

	if data.Name != nil {
		model.Name = *data.Name
	}

	if data.Line != nil {
		adddress.Line = *data.Line
	}

	if data.City != nil {
		adddress.City = *data.City
	}

	if data.State != nil {
		adddress.State = *data.State
	}

	if data.Country != nil {
		adddress.Country = *data.Country
	}

	if data.Zipcode != nil {
		adddress.Zipcode = *data.Zipcode
	}

	if data.Timezone != nil {
		adddress.Timezone = *data.Timezone
	}

	if !reflect.ValueOf(adddress).IsZero() {
		model.AddressInfo = adddress
	}

	if data.MapUrl != nil {
		model.MapUrl = *data.MapUrl
	}

	if data.Contact != nil {
		model.Contact = *data.Contact
	}

	return model
}
