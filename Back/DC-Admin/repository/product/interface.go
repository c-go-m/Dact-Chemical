package product

import "github.com/c-go-m/DC-Admin/common/entity"

type IProductRepository interface {
	FindAll() ([]entity.Product, error)
	Create(*entity.Product) error
	Update(*entity.Product) error
	Delete(*entity.Product) error
	FindById(int) (entity.Product, error)
}
