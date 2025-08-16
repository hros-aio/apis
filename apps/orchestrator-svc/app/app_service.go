
package app

import "github.com/tinh-tinh/tinhtinh/v2/core"

type appService struct {}

func (s *appService) Create(input interface{}) interface{} {
	return nil
}

func (s *appService) Find() interface{} {
	return nil
}

func (s *appService) FindById(id string) interface{} {
	return nil
}

func (s *appService) Update(id string,input interface{}) interface{} {
	return nil
}

func (s *appService) Delete(id string) interface{} {
	return nil
}

func NewService(module core.Module) core.Provider {
	svc := module.NewProvider(&appService{})

	return svc
}
	