package ipfs

import (
	shell "github.com/ipfs/go-ipfs-api"
	"io"
)

// UploadIPFS 数据上传到ipfs
func UploadIPFS(sh *shell.Shell, file io.Reader) (string, error) {
	return sh.Add(file)
}
