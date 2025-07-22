package query

import "github.com/tinh-tinh/mongoose/v2"

type TenantStore struct {
	model mongoose.Model[TenantSchema]
}

func (s *TenantStore) GetById(id string) (*TenantSchema, error) {
	return s.model.FindByID(id)
}

func (s *TenantStore) GetAll() ([]*TenantSchema, error) {
	return s.model.Find(nil)
}
