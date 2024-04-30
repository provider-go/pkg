package ipfs

import (
	shell "github.com/ipfs/go-ipfs-api"
	"io"
)

// CatIPFS 从ipfs下载数据
func CatIPFS(sh *shell.Shell, hash string) ([]byte, error) {
	read, err := sh.Cat(hash)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(read)
}
