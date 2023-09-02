package attached

import "github.com/c-go-m/DC-Admin/common/entity"

type IAttachedService interface {
	FindAll() ([]entity.Attached, error)
	Create(*entity.Attached) error
	Update(*entity.Attached) error
	Delete(*entity.Attached) error
	FindById(int) (entity.Attached, error)
}
