package useError

import (
	"reflect"

	"github.com/c-go-m/DC-Admin/common/utility/constant"
)

func ControlError(err error) string {
	if err != nil {
		switch reflect.TypeOf(err).String() {
		case constant.TypeGeneralError:
			return err.Error()
		case constant.TypeBussinesError:
			return err.Error()
		default:
			return err.Error()
		}
	}
	return ""
}

func GetTypeError(err error) string {
	return reflect.TypeOf(err).String()
}
