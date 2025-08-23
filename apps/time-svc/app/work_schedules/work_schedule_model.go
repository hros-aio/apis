package work_schedules

import (
	"github.com/hros-aio/apis/libs/mongodoc/common/base"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Duration struct {
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
}

type WorkScheduleModel struct {
	base.Model
	TenantID   string             `json:"tenantId"`
	CompanyID  string             `json:"companyId"`
	LocationID primitive.ObjectID `json:"locationId"`
	Name       string             `json:"name"`
	IsDefault  bool               `json:"isDefault"`
	TotalHours float32            `json:"totalHours"`
	StartAt    string             `json:"startAt"`
	EndAt      string             `json:"endAt"`
	BreakTime  []Duration         `json:"breakTime"`
	WorkDays   []int              `json:"workDays"`
}

func (model WorkScheduleModel) DataMapper() *WorkScheduleSchema {
	data := &WorkScheduleSchema{
		TenantID:   model.TenantID,
		CompanyID:  model.CompanyID,
		LocationID: model.LocationID,
		Name:       model.Name,
		IsDefault:  model.IsDefault,
		TotalHours: model.TotalHours,
		StartAt:    model.StartAt,
		EndAt:      model.EndAt,
		WorkDays:   model.WorkDays,
	}

	data.BreakTime = make([]DurationSchema, len(model.BreakTime))
	for i, duration := range model.BreakTime {
		data.BreakTime[i] = DurationSchema(duration)
	}

	return data
}
