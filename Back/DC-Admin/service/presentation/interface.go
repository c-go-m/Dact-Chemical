package presentation

import "github.com/c-go-m/DC-Admin/common/entity"

type IPresentationService interface {
	FindAll() ([]entity.Presentation, error)
	Create(*entity.Presentation) error
	Update(*entity.Presentation) error
	Delete(*entity.Presentation) error
	FindById(int) (entity.Presentation, error)
}
