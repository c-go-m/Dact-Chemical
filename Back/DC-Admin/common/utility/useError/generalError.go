package useError

import (
	"github.com/c-go-m/DC-Admin/common/utility/constant"
)

type GeneralError struct {
	msg string
}

func (e *GeneralError) Error() string {
	return e.msg
}

var ErrorConnectionDataBase = &GeneralError{msg: constant.ErrConnectionDataBase}
var ErrorConnectionStorage = &GeneralError{msg: constant.ErrConnectionDataBase}
var ErrorEnvironmentVariable = &GeneralError{msg: constant.ErrEnvironmentVariable}
var ErrorInitServer = &GeneralError{msg: constant.ErrInitServer}
var ErrDataInvalid = &GeneralError{msg: constant.ErrDataInvalid}
var ErrSaveFile = &GeneralError{msg: constant.ErrDataInvalid}
