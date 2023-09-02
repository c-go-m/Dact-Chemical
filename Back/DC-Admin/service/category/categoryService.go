package category

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/constant"
	"github.com/c-go-m/DC-Admin/repository/category"
	"github.com/c-go-m/DC-Admin/service/attached"
)

type CategoryService struct {
	repository      category.ICategoryRepository
	attachedService attached.IAttachedService
}

func New(repository category.ICategoryRepository, attachedService attached.IAttachedService) *CategoryService {
	service := CategoryService{
		repository:      repository,
		attachedService: attachedService,
	}

	return &service
}

func (base *CategoryService) FindAll() (categories []entity.Category, err error) {
	return base.repository.FindAll()
}

func (base *CategoryService) Create(category *entity.Category) (err error) {

	category.Attached.Url = constant.PathImgSystem + constant.PathCategory + category.Attached.Name + "." + category.Attached.Extension

	if err = base.attachedService.Create(&category.Attached); err != nil {
		return err
	}

	if err = base.repository.Create(category); err != nil {
		base.attachedService.Delete(&category.Attached)
		return err
	}

	return nil
}

func (base *CategoryService) Update(category *entity.Category) error {
	return base.repository.Update(category)
}

func (base *CategoryService) Delete(category *entity.Category) error {
	return base.repository.Delete(category)
}

func (base *CategoryService) FindById(id int) (category entity.Category, err error) {
	return base.repository.FindById(id)
}
