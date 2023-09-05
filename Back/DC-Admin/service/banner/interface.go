package banner

import "github.com/c-go-m/DC-Admin/common/entity"

type IBannerService interface {
	FindAll() ([]entity.Banner, error)
	Create(*entity.Banner) error
	Update(*entity.Banner) error
	Delete(int) error
	FindById(int) (entity.Banner, error)
}
