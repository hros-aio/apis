package messages

import "time"

type TenantCreatedPayload struct {
	Id        string     `json:"id,omitempty"`
	Name      string     `json:"name.omitempty"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	TenantId  string     `json:"tenantId,omitempty"`
	Domain    string     `json:"domain,omitempty"`
}
