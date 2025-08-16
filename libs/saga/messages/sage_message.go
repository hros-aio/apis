package messages

import (
	"github.com/google/uuid"
)

type SagaRegisteredPayload struct {
	SessionId    uuid.UUID `json:"sessionId,omitempty"`
	Event        string    `json:"event,omitempty"`
	PreviousData any       `json:"previousData,omitempty"`
	Data         any       `json:"data,omitempty"`
	Step         int       `json:"step,omitempty"`
}

type SagaRollbackPayload struct {
	SessionId uuid.UUID `json:"sessionId,omitempty"`
	Event     string    `json:"event,omitempty"`
}
