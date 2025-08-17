package messages

type AddressInfo struct {
	Line     string `json:"addressLine"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Zipcode  string `json:"zipCode"`
	Timezone string `json:"timezone"`
}
type LocationCreatedPayload struct {
	AddressInfo
	ID            string `json:"id"`
	Name          string `json:"name"`
	TenantID      string `json:"tenantId,omitempty"`
	CompanyID     string `json:"companyId,omitempty"`
	MapUrl        string `json:"mapUrl"`
	IsHeadquarter bool   `json:"isHeadquarter"`
}

type LocationUpdatedPayload struct {
	PreviouseData LocationCreatedPayload `json:"previousData,omitzero"`
	Data          LocationCreatedPayload `json:"data,omitzero"`
}

type LocationDeletedPayload struct {
	Id        string `json:"id,omitempty"`
	TenantID  string `json:"tenantId,omitempty"`
	CompanyID string `json:"companyId,omitempty"`
}
