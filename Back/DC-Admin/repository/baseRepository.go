package repository

import (
	"encoding/json"

	"github.com/c-go-m/DC-Admin/common/utility/constant"
	"github.com/c-go-m/DC-Admin/common/utility/database"
	"github.com/c-go-m/DC-Admin/common/utility/useError"
)

type BaseRepository struct {
	Data database.DataBase
}

type GormErr struct {
	Code    string
	Message string
}

func (base *BaseRepository) ControlError(err error) error {

	if err != nil {

		byteErr, _ := json.Marshal(err)
		var newError GormErr

		json.Unmarshal((byteErr), &newError)

		switch newError.Code {
		case constant.CodKeyDuplicate:
			return useError.ErrorDuplicateRegistry
		default:
			return useError.ErrDataInvalid
		}
	}
	return nil
}
