package works_chedules

import (
	"github.com/hros-aio/apis/libs/mongodoc/common/base"
	"github.com/tinh-tinh/mongoose/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DurationSchema struct {
	StartAt string `bson:"startAt" json:"startAt"`
	EndAt   string `bson:"endAt" json:"endAt"`
}

type WorkScheduleSchema struct {
	mongoose.BaseSchema `bson:"inline"`
	TenantID            string             `bson:"tenantId" json:"tenantId"`
	CompanyID           string             `bson:"companyId" json:"companyId"`
	LocationID          primitive.ObjectID `bson:"locationId" json:"locationId"`
	Name                string             `bson:"name" json:"name"`
	IsDefault           bool               `bson:"isDefault" json:"isDefault"`
	TotalHours          float32            `bson:"totalHours" json:"totalHours"`
	StartAt             string             `bson:"startAt" json:"startAt"`
	EndAt               string             `bson:"endAt" json:"endAt"`
	BreakTime           []DurationSchema   `bson:"breakTime" json:"breakTime"`
	WorkDays            []int              `bson:"workDays" json:"workDays"`
}

func (WorkScheduleSchema) CollectionName() string {
	return "work_schedules"
}

func (data WorkScheduleSchema) Dto() *WorkScheduleModel {
	model := &WorkScheduleModel{
		Model:      base.Model(data.BaseSchema),
		TenantID:   data.TenantID,
		CompanyID:  data.CompanyID,
		LocationID: data.LocationID,
		Name:       data.Name,
		IsDefault:  data.IsDefault,
		TotalHours: data.TotalHours,
		StartAt:    data.StartAt,
		EndAt:      data.EndAt,
		BreakTime:  make([]Duration, len(data.BreakTime)),
		WorkDays:   data.WorkDays,
	}

	for i, duration := range data.BreakTime {
		model.BreakTime[i] = Duration(duration)
	}

	return model
}
