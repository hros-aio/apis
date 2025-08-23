package messages

type CompanyCreatedPayload struct {
	ID       string `json:"id,omitempty"`
	TenantID string `json:"tenantId,omitempty"`
	Name     string `json:"name,omitempty"`
}
