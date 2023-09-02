package presentation

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/constant"
	"github.com/c-go-m/DC-Admin/repository/presentation"
	"github.com/c-go-m/DC-Admin/service/attached"
)

type PresentationService struct {
	repository      presentation.IPresentationRepository
	attachedService attached.IAttachedService
}

func New(repository presentation.IPresentationRepository, attachedService attached.IAttachedService) *PresentationService {
	service := PresentationService{
		repository:      repository,
		attachedService: attachedService,
	}

	return &service
}

func (base *PresentationService) FindAll() (presentations []entity.Presentation, err error) {
	return base.repository.FindAll()
}

func (base *PresentationService) Create(presentation *entity.Presentation) (err error) {

	presentation.Attached.Url = constant.PathImgSystem + constant.PathPresentation + presentation.Attached.Name + "." + presentation.Attached.Extension

	if err = base.attachedService.Create(&presentation.Attached); err != nil {
		return err
	}

	if err = base.repository.Create(presentation); err != nil {
		base.attachedService.Delete(&presentation.Attached)
		return err
	}

	return nil
}

func (base *PresentationService) Update(presentation *entity.Presentation) error {
	return base.repository.Update(presentation)
}

func (base *PresentationService) Delete(presentation *entity.Presentation) error {
	return base.repository.Delete(presentation)
}

func (base *PresentationService) FindById(id int) (presentation entity.Presentation, err error) {
	return base.repository.FindById(id)
}
