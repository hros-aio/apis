package messages

type DepartmentCreatedPayload struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	TenantID   string `json:"tenantId,omitempty"`
	CompanyID  string `json:"companyId,omitempty"`
	Code       string `json:"code,omitempty"`
	IsDivision bool   `json:"isDivision,omitempty"`
	ParentID   string `json:"parentId,omitempty"`
}

type DepartmentUpdatedPayload struct {
	PreviouseData DepartmentCreatedPayload `json:"previousData,omitzero"`
	Data          DepartmentCreatedPayload `json:"data,omitzero"`
}

type DepartmentDeletedPayload struct {
	ID        string `json:"id,omitempty"`
	TenantID  string `json:"tenantId,omitempty"`
	CompanyID string `json:"companyId,omitempty"`
}
