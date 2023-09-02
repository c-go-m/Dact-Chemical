package menu

import "github.com/c-go-m/DC-Admin/common/entity"

type IMenuRepository interface {
	FindAll() ([]entity.Menu, error)
	Create(*entity.Menu) error
	Update(*entity.Menu) error
	Delete(*entity.Menu) error
	FindById(int) (entity.Menu, error)
}
