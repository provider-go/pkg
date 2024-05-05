package ipfs

import (
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/provider-go/pkg/file/typefile"
	"io"
)

type StorageFileIPFS struct {
	client *shell.Shell
}

func NewStorageFileIPFS(cfg typefile.ConfigStorageFile) *StorageFileIPFS {
	client := shell.NewShell(cfg.Endpoints)

	return &StorageFileIPFS{client: client}
}

func (a *StorageFileIPFS) Upload(file io.Reader) (string, error) {
	return a.client.Add(file)
}

func (a *StorageFileIPFS) Download(hash string) ([]byte, error) {
	read, err := a.client.Cat(hash)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(read)
}
