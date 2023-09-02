package product

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/database"
	"github.com/c-go-m/DC-Admin/repository"
)

type ProductRepository struct {
	repository.BaseRepository
}

func New(dB *database.DataBase) *ProductRepository {
	repo := ProductRepository{
		repository.BaseRepository{
			Data: *dB,
		},
	}

	return &repo
}

func (base *ProductRepository) FindAll() (products []entity.Product, err error) {
	err = base.Data.Db.Joins("Attached").Joins("Category").Preload("Presentation.Attached").Find(&products).Error
	return
}

func (base *ProductRepository) Create(product *entity.Product) error {
	result := base.Data.Db.Create(product)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *ProductRepository) Update(product *entity.Product) error {
	result := base.Data.Db.Updates(product)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *ProductRepository) Delete(product *entity.Product) error {
	result := base.Data.Db.Delete(&product, product.Id)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *ProductRepository) FindById(id int) (product entity.Product, err error) {
	err = base.Data.Db.Joins("Attached").Joins("Category").Find(&product, id).Error
	return
}
