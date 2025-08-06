package base

import (
	"time"

	"github.com/google/uuid"
	"github.com/tinh-tinh/sqlorm/v2"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID      `json:"id"`
	CreatedAt *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt *time.Time     `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}

func (Model) FromData(data sqlorm.Model) Model {
	return Model{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}
