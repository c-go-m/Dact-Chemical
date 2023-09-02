package category

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/database"
	"github.com/c-go-m/DC-Admin/repository"
)

type CategoryRepository struct {
	repository.BaseRepository
}

func New(dB *database.DataBase) *CategoryRepository {
	repo := CategoryRepository{
		repository.BaseRepository{
			Data: *dB,
		},
	}

	return &repo
}

func (base *CategoryRepository) FindAll() (categories []entity.Category, err error) {
	err = base.Data.Db.Joins("Attached").Find(&categories).Error
	return
}

func (base *CategoryRepository) Create(category *entity.Category) error {
	result := base.Data.Db.Create(category)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *CategoryRepository) Update(category *entity.Category) error {
	result := base.Data.Db.Updates(category)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *CategoryRepository) Delete(category *entity.Category) error {
	result := base.Data.Db.Delete(&category, category.Id)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *CategoryRepository) FindById(id int) (category entity.Category, err error) {
	err = base.Data.Db.Joins("Attached").Find(&category, id).Error
	return
}
