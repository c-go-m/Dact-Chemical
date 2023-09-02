package storage

type IStorage interface {
	SaveFile(name string, data []byte) (err error)
	LoadFile(name string) (data []byte, err error)
	DeleteFile(url string) (err error)
}
