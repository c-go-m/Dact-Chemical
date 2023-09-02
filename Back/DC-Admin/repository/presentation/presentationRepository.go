package presentation

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/database"
	"github.com/c-go-m/DC-Admin/repository"
)

type PresentationRepository struct {
	repository.BaseRepository
}

func New(dB *database.DataBase) *PresentationRepository {
	repo := PresentationRepository{
		repository.BaseRepository{
			Data: *dB,
		},
	}

	return &repo
}

func (base *PresentationRepository) FindAll() (presentations []entity.Presentation, err error) {
	err = base.Data.Db.Joins("Attached").Find(&presentations).Error
	return
}

func (base *PresentationRepository) Create(presentation *entity.Presentation) error {
	result := base.Data.Db.Create(presentation)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *PresentationRepository) Update(presentation *entity.Presentation) error {
	result := base.Data.Db.Updates(presentation)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *PresentationRepository) Delete(presentation *entity.Presentation) error {
	result := base.Data.Db.Delete(&presentation, presentation.Id)
	return base.BaseRepository.ControlError(result.Error)
}

func (base *PresentationRepository) FindById(id int) (presentation entity.Presentation, err error) {
	err = base.Data.Db.Joins("Attached").Find(&presentation, id).Error
	return
}
