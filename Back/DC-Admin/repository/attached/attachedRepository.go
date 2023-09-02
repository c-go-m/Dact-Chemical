package attached

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/database"
	"github.com/c-go-m/DC-Admin/repository"
)

type AttachedRepository struct {
	repository.BaseRepository
}

func New(dB *database.DataBase) *AttachedRepository {
	repo := AttachedRepository{
		repository.BaseRepository{
			Data: *dB,
		},
	}
	return &repo
}

func (base *AttachedRepository) FindAll() (attacheds []entity.Attached, err error) {
	err = base.Data.Db.Find(&attacheds).Error
	return
}

func (base *AttachedRepository) Create(attached *entity.Attached) error {
	result := base.Data.Db.Create(attached)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *AttachedRepository) Update(attached *entity.Attached) error {
	result := base.Data.Db.Updates(attached)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *AttachedRepository) Delete(attached *entity.Attached) error {
	result := base.Data.Db.Delete(&attached, attached.Id)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *AttachedRepository) FindById(id int) (attached entity.Attached, err error) {
	err = base.Data.Db.Find(&attached, id).Error
	return
}
