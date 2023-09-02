package filesystem

import (
	"os"

	"github.com/c-go-m/DC-Admin/common/utility/useError"
)

type FileSystem struct {
}

func New() *FileSystem {
	file := FileSystem{}
	return &file
}

func (base *FileSystem) SaveFile(name string, data []byte) (err error) {

	if err = os.WriteFile(name, data, 0655); err != nil {
		return useError.ErrSaveFile
	}

	return
}

func (base *FileSystem) LoadFile(name string) (data []byte, err error) {
	content, err := os.ReadFile(name)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func (base *FileSystem) DeleteFile(url string) (err error) {

	if err = os.Remove(url); err != nil {
		return useError.ErrSaveFile
	}

	return
}
