package messages

import "time"

type ContactPerson struct {
	ContactName  string `json:"name"`
	ContactEmail string `json:"email"`
	ContactPhone string `json:"phone"`
}

type TenantCreatedPayload struct {
	Id        string        `json:"id,omitempty"`
	Name      string        `json:"name.omitempty"`
	CreatedAt *time.Time    `json:"createdAt,omitempty"`
	TenantId  string        `json:"tenantId,omitempty"`
	Domain    string        `json:"domain,omitempty"`
	Contact   ContactPerson `json:"contact,omitzero"`
}
