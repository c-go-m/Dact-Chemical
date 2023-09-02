package useError

import (
	"github.com/c-go-m/DC-Admin/common/utility/constant"
)

type BussinesError struct {
	msg string
}

func (e *BussinesError) Error() string {
	return e.msg
}

var ErrorDuplicateRegistry = &BussinesError{msg: constant.ErrDuplicateRegistry}
var ErrCreateRegistry = &BussinesError{msg: constant.ErrCreateRegistry}
var ErrFindRegistry = &BussinesError{msg: constant.ErrFindRegistry}
var ErrUpdateRegistry = &BussinesError{msg: constant.ErrUpdateRegistry}
var ErrDeleteRegistry = &BussinesError{msg: constant.ErrDeleteRegistry}
