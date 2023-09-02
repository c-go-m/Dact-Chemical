package information

import "github.com/c-go-m/DC-Admin/common/entity"

type IInformationRepository interface {
	FindAll() ([]entity.Information, error)
	Create(*entity.Information) error
	Update(*entity.Information) error
	Delete(*entity.Information) error
	FindById(int) (entity.Information, error)
}
