package banner

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/database"
	"github.com/c-go-m/DC-Admin/repository"
)

type BannerRepository struct {
	repository.BaseRepository
}

func New(dB *database.DataBase) *BannerRepository {
	repo := BannerRepository{
		repository.BaseRepository{
			Data: *dB,
		},
	}

	return &repo
}

func (base *BannerRepository) FindAll() (banners []entity.Banner, err error) {
	err = base.Data.Db.Joins("Attached").Find(&banners).Error
	return
}

func (base *BannerRepository) Create(banner *entity.Banner) error {
	result := base.Data.Db.Create(banner)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *BannerRepository) Update(banner *entity.Banner) error {
	result := base.Data.Db.Updates(banner)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *BannerRepository) Delete(banner *entity.Banner) error {
	result := base.Data.Db.Delete(&banner, banner.Id)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *BannerRepository) FindById(id int) (banner entity.Banner, err error) {
	err = base.Data.Db.Joins("Attached").Find(&banner, id).Error
	return
}
