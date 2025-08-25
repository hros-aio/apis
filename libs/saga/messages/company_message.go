package messages

import (
	"time"

	"github.com/google/uuid"
)

type CompanyActivatedPayload struct {
	ID             string     `json:"id,omitempty"`
	TenantID       string     `json:"tenantId,omitempty"`
	Name           string     `json:"name,omitempty"`
	LegalName      string     `json:"legalName,omitempty"`
	Status         string     `json:"status,omitempty"`
	RegistrationNo string     `json:"registrationNo,omitempty"`
	TaxID          string     `json:"taxId,omitempty"`
	Website        string     `json:"website,omitempty"`
	Industry       string     `json:"industry,omitempty"`
	Size           int        `json:"size,omitempty"`
	Logo           string     `json:"logo,omitempty"`
	FoundedDate    time.Time  `json:"foundedDate,omitempty"`
	HoldingID      *uuid.UUID `json:"holdingId,omitempty"`
}
