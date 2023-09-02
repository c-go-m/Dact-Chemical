package banner

import "github.com/c-go-m/DC-Admin/common/entity"

type IBannerRepository interface {
	FindAll() ([]entity.Banner, error)
	Create(*entity.Banner) error
	Update(*entity.Banner) error
	Delete(*entity.Banner) error
	FindById(int) (entity.Banner, error)
}
