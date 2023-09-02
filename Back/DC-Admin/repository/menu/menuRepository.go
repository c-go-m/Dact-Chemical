package menu

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/database"
	"github.com/c-go-m/DC-Admin/repository"
)

type MenuRepository struct {
	repository.BaseRepository
}

func New(dB *database.DataBase) *MenuRepository {
	repo := MenuRepository{
		repository.BaseRepository{
			Data: *dB,
		},
	}

	return &repo
}

func (base *MenuRepository) FindAll() (menus []entity.Menu, err error) {
	err = base.Data.Db.Find(&menus).Error
	return
}

func (base *MenuRepository) Create(menu *entity.Menu) error {
	result := base.Data.Db.Create(menu)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *MenuRepository) Update(menu *entity.Menu) error {
	result := base.Data.Db.Updates(menu)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *MenuRepository) Delete(menu *entity.Menu) error {
	result := base.Data.Db.Delete(&menu, menu.Id)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *MenuRepository) FindById(id int) (menu entity.Menu, err error) {
	err = base.Data.Db.Find(&menu, id).Error
	return
}
