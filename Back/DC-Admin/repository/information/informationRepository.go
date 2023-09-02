package information

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/database"
	"github.com/c-go-m/DC-Admin/repository"
)

type InformationRepository struct {
	repository.BaseRepository
}

func New(dB *database.DataBase) *InformationRepository {
	repo := InformationRepository{
		repository.BaseRepository{
			Data: *dB,
		},
	}

	return &repo
}

func (base *InformationRepository) FindAll() (informations []entity.Information, err error) {
	err = base.Data.Db.Joins("Attached").Find(&informations).Error
	return
}

func (base *InformationRepository) Create(information *entity.Information) error {
	result := base.Data.Db.Create(information)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *InformationRepository) Update(information *entity.Information) error {
	result := base.Data.Db.Updates(information)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *InformationRepository) Delete(information *entity.Information) error {
	result := base.Data.Db.Delete(&information, information.Id)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *InformationRepository) FindById(id int) (information entity.Information, err error) {
	err = base.Data.Db.Joins("Attached").Find(&information, id).Error
	return
}
