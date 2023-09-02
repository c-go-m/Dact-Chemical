package attached

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/constant"
	"github.com/c-go-m/DC-Admin/common/utility/storage"
	"github.com/c-go-m/DC-Admin/repository/attached"
)

type AttachedService struct {
	repository attached.IAttachedRepository
	storage    storage.IStorage
}

func New(repo attached.IAttachedRepository, storage storage.IStorage) *AttachedService {
	service := AttachedService{
		repository: repo,
		storage:    storage,
	}

	return &service
}

func (base *AttachedService) FindAll() (attacheds []entity.Attached, err error) {
	return base.repository.FindAll()
}

func (base *AttachedService) Create(attached *entity.Attached) (err error) {

	if err = base.storage.SaveFile(attached.Url, attached.Data); err != nil {
		return err
	}

	if err = base.repository.Create(attached); err != nil {

		if err.Error() != constant.ErrDuplicateRegistry {
			base.storage.DeleteFile(attached.Url)
			return err
		}
		return err
	}

	return err
}

func (base *AttachedService) Update(attached *entity.Attached) error {
	return base.repository.Update(attached)
}

func (base *AttachedService) Delete(attached *entity.Attached) (err error) {

	if err = base.storage.DeleteFile(attached.Url); err != nil {
		return err
	}

	if err = base.repository.Delete(attached); err != nil {
		return err
	}

	return
}

func (base *AttachedService) FindById(id int) (attached entity.Attached, err error) {
	return base.repository.FindById(id)
}
