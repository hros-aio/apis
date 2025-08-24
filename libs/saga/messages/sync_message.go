package messages

import (
	"github.com/google/uuid"
)

type SyncDataPayload struct {
	PreviousData any `json:"previousData,omitempty"`
	Data         any `json:"data,omitempty"`
}

type SyncRegisteredPayload struct {
	SyncDataPayload
	SessionId uuid.UUID `json:"sessionId,omitempty"`
	Event     string    `json:"event,omitempty"`
	GroupIds  []string  `json:"groupIds,omitempty"`
}

type SyncRetryPayload struct {
	SessionId uuid.UUID `json:"sessionId,omitempty"`
}
