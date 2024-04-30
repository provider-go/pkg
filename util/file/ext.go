package file

import (
	"mime"
	"path"
)

func GetContentType(filePath string) string {
	ext := path.Ext(filePath)
	if ext != "" {
		// 将扩展名前的点去除
		ext = ext[1:]
	}
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		return "application/octet-stream"
	}
	return contentType
}
