package menu

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/repository/menu"
)

type MenuService struct {
	repository menu.IMenuRepository
}

func New(repository menu.IMenuRepository) *MenuService {
	service := MenuService{
		repository: repository,
	}

	return &service
}

func (base *MenuService) FindAll() (menus []entity.Menu, err error) {
	return base.repository.FindAll()
}

func (base *MenuService) Create(menu *entity.Menu) (err error) {
	return base.repository.Create(menu)
}

func (base *MenuService) Update(menu *entity.Menu) error {
	return base.repository.Update(menu)
}

func (base *MenuService) Delete(menu *entity.Menu) error {
	return base.repository.Delete(menu)
}

func (base *MenuService) FindById(id int) (menu entity.Menu, err error) {
	return base.repository.FindById(id)
}
