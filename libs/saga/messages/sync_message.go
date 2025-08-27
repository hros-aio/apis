package messages

import (
	"github.com/google/uuid"
)

type SyncRegisteredPayload struct {
	Data      any       `json:"data,omitempty"`
	SessionId uuid.UUID `json:"sessionId,omitempty"`
	Event     string    `json:"event,omitempty"`
	GroupIds  []string  `json:"groupIds,omitempty"`
}

type SyncRetryPayload struct {
	SessionId uuid.UUID `json:"sessionId,omitempty"`
}
