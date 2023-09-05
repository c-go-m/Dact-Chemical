package banner

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/constant"
	"github.com/c-go-m/DC-Admin/repository/banner"
	"github.com/c-go-m/DC-Admin/service/attached"
)

type BannerService struct {
	repository      banner.IBannerRepository
	attachedService attached.IAttachedService
}

func New(repository banner.IBannerRepository, attachedService attached.IAttachedService) *BannerService {
	service := BannerService{
		repository:      repository,
		attachedService: attachedService,
	}

	return &service
}

func (base *BannerService) FindAll() (banners []entity.Banner, err error) {
	return base.repository.FindAll()
}

func (base *BannerService) Create(banner *entity.Banner) (err error) {

	banner.Attached.Url = constant.PathImgSystem + constant.PathBanner + banner.Attached.Name + "." + banner.Attached.Extension

	if err = base.attachedService.Create(&banner.Attached); err != nil {
		return err
	}

	if err = base.repository.Create(banner); err != nil {
		base.attachedService.Delete(&banner.Attached)
		return err
	}

	return nil
}

func (base *BannerService) Update(banner *entity.Banner) error {
	return base.repository.Update(banner)
}

func (base *BannerService) Delete(id int) (err error) {
	banner, err := base.repository.FindById(id)
	if err != nil {
		return err
	}
	if err = base.repository.Delete(&banner); err != nil {
		return err
	}
	return base.attachedService.Delete(&banner.Attached)
}

func (base *BannerService) FindById(id int) (banner entity.Banner, err error) {
	return base.repository.FindById(id)
}
