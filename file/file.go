package file

import (
	"github.com/provider-go/pkg/file/ipfs"
	"github.com/provider-go/pkg/file/typefile"
	"io"
)

type StorageFile interface {
	Upload(io.Reader) (string, error)
	Download(string) ([]byte, error)
}

func NewStorageFile(provider string, cfg typefile.ConfigStorageFile) StorageFile {
	switch provider {
	case "ipfs":
		return ipfs.NewStorageFileIPFS(cfg)
	default:

		return ipfs.NewStorageFileIPFS(cfg)
	}
}
