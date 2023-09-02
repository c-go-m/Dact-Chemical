package product

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/constant"
	"github.com/c-go-m/DC-Admin/repository/product"
	"github.com/c-go-m/DC-Admin/service/attached"
)

type ProductService struct {
	repository      product.IProductRepository
	attachedService attached.IAttachedService
}

func New(repository product.IProductRepository, attachedService attached.IAttachedService) *ProductService {
	service := ProductService{
		repository:      repository,
		attachedService: attachedService,
	}

	return &service
}

func (base *ProductService) FindAll() (products []entity.Product, err error) {
	return base.repository.FindAll()
}

func (base *ProductService) Create(product *entity.Product) (err error) {

	product.Attached.Url = constant.PathImgSystem + constant.PathProduct + product.Attached.Name + "." + product.Attached.Extension

	if err = base.attachedService.Create(&product.Attached); err != nil {
		return err
	}

	if err = base.repository.Create(product); err != nil {
		base.attachedService.Delete(&product.Attached)
		return err
	}

	return nil
}

func (base *ProductService) Update(product *entity.Product) error {
	return base.repository.Update(product)
}

func (base *ProductService) Delete(product *entity.Product) error {
	return base.repository.Delete(product)
}

func (base *ProductService) FindById(id int) (product entity.Product, err error) {
	return base.repository.FindById(id)
}
