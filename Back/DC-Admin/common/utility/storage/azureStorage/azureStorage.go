package azurestorage

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/c-go-m/DC-Admin/common/utility/config"
	"github.com/c-go-m/DC-Admin/common/utility/useError"
)

type AzureStorage struct {
	storage    *azblob.Client
	urlStorage string
	container  string
	context    context.Context
}

func New() (*AzureStorage, error) {
	url := config.AzureStorage

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, useError.ErrorConnectionStorage
	}

	client, err := azblob.NewClient(url, credential, nil)
	if err != nil {
		return nil, useError.ErrorConnectionStorage
	}

	storage := AzureStorage{
		storage:    client,
		urlStorage: config.AzureStorage,
		container:  config.Container,
		context:    context.Background(),
	}
	return &storage, nil
}

func (base *AzureStorage) SaveFile(name string, data []byte) (err error) {
	_, err = base.storage.UploadBuffer(base.context, base.container, name, data, &azblob.UploadBufferOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return useError.ErrSaveFile
	}
	fmt.Println("Creo el archivo")
	return err
}

func (base *AzureStorage) LoadFile(name string) (data []byte, err error) {
	get, err := base.storage.DownloadStream(base.context, base.container, name, nil)
	if err != nil {
		return nil, useError.ErrSaveFile
	}

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(base.context, &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	if err != nil {
		return nil, useError.ErrSaveFile
	}

	err = retryReader.Close()
	if err != nil {
		return nil, useError.ErrSaveFile
	}

	return nil, nil
}

func (base *AzureStorage) DeleteFile(url string) (err error) {
	_, err = base.storage.DeleteBlob(base.context, base.container, url, nil)
	if err != nil {
		return useError.ErrSaveFile
	}

	return nil
}
