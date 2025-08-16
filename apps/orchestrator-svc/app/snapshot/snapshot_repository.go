package snapshot

import (
	"github.com/tinh-tinh/mongoose/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type SnapshotRepository struct {
	Model *mongoose.Model[SnapshotSchema]
}

func NewRepository(module core.Module) core.Provider {
	model := mongoose.InjectModel[SnapshotSchema](module)

	return module.NewProvider(&SnapshotRepository{
		Model: model,
	})
}
