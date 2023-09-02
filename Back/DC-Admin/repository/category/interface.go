package category

import "github.com/c-go-m/DC-Admin/common/entity"

type ICategoryRepository interface {
	FindAll() ([]entity.Category, error)
	Create(*entity.Category) error
	Update(*entity.Category) error
	Delete(*entity.Category) error
	FindById(int) (entity.Category, error)
}
