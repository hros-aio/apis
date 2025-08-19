package works_chedules

import "github.com/tinh-tinh/mongoose/v2"

type CreateWorkScheduleInput struct {
	Name       string     `json:"name" validate:"required" example:"Work Schedule"`
	TotalHours float32    `json:"totalHours" validate:"required" example:"8"`
	StartAt    string     `json:"startAt" validate:"required" example:"08:30"`
	EndAt      string     `json:"endAt" validate:"required" example:"17:30"`
	BreakTime  []Duration `json:"breakTime"`
	WorkDays   []int      `json:"workDays" validate:"required" example:"[1, 2, 3, 4, 5]"`
	Location   string     `json:"location" validate:"required" example:"Remote"`
}

func (data CreateWorkScheduleInput) Dto() *WorkScheduleModel {
	return &WorkScheduleModel{
		Name:       data.Name,
		TotalHours: data.TotalHours,
		StartAt:    data.StartAt,
		EndAt:      data.EndAt,
		BreakTime:  data.BreakTime,
		WorkDays:   data.WorkDays,
		LocationID: mongoose.ToObjectID(data.Location),
	}
}

type UpdateWorkScheduleInput struct {
	Name       string     `json:"name"`
	TotalHours float32    `json:"totalHours"`
	StartAt    string     `json:"startAt"`
	EndAt      string     `json:"endAt"`
	BreakTime  []Duration `json:"breakTime"`
	WorkDays   []int      `json:"workDays"`
}

func (data UpdateWorkScheduleInput) Dto() *WorkScheduleModel {
	return &WorkScheduleModel{
		Name:       data.Name,
		TotalHours: data.TotalHours,
		StartAt:    data.StartAt,
		EndAt:      data.EndAt,
		BreakTime:  data.BreakTime,
		WorkDays:   data.WorkDays,
	}
}
