package works_chedules

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/tinh-tinh/mongoose/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkScheduleRepository struct {
	Model *mongoose.Model[WorkScheduleSchema]
}

func NewRepository(module core.Module) core.Provider {
	model := mongoose.InjectModel[WorkScheduleSchema](module)

	return module.NewProvider(&WorkScheduleRepository{
		Model: model,
	})
}

func (r *WorkScheduleRepository) Create(model *WorkScheduleModel) (*mongo.InsertOneResult, error) {
	data := model.DataMapper()
	result, err := r.Model.Create(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *WorkScheduleRepository) FindAll(filter any, queryParams middleware.Paginate) ([]*WorkScheduleModel, error) {
	data, err := r.Model.Find(filter, mongoose.QueriesOptions{
		Skip:  int64(queryParams.Skip),
		Limit: int64(queryParams.Limit),
	})
	if err != nil {
		return nil, err
	}

	models := []*WorkScheduleModel{}
	for _, item := range data {
		models = append(models, item.Dto())
	}
	return models, nil
}
