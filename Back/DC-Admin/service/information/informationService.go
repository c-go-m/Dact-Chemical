package information

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/constant"
	"github.com/c-go-m/DC-Admin/repository/information"
	"github.com/c-go-m/DC-Admin/service/attached"
)

type InformationService struct {
	repository      information.IInformationRepository
	attachedService attached.IAttachedService
}

func New(repository information.IInformationRepository, attachedService attached.IAttachedService) *InformationService {
	service := InformationService{
		repository:      repository,
		attachedService: attachedService,
	}

	return &service
}

func (base *InformationService) FindAll() (informations []entity.Information, err error) {
	return base.repository.FindAll()
}

func (base *InformationService) Create(information *entity.Information) (err error) {

	information.Attached.Url = constant.PathImgSystem + constant.PathInformation + information.Attached.Name + "." + information.Attached.Extension

	if err = base.attachedService.Create(&information.Attached); err != nil {
		return err
	}

	if err = base.repository.Create(information); err != nil {
		base.attachedService.Delete(&information.Attached)
		return err
	}

	return nil
}

func (base *InformationService) Update(information *entity.Information) error {
	return base.repository.Update(information)
}

func (base *InformationService) Delete(information *entity.Information) error {
	return base.repository.Delete(information)
}

func (base *InformationService) FindById(id int) (information entity.Information, err error) {
	return base.repository.FindById(id)
}
