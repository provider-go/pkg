package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/provider-go/pkg/ipfs"
	"net/http"
)

func main() {
	router := gin.Default()
	sh := shell.NewShell("192.168.0.103:5001")
	// 设置文件上传的路由
	router.POST("/upload", func(c *gin.Context) {
		// 单文件上传
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("上传文件错误: %s", err.Error()))
			return
		}
		fmt.Println(file.Size)
		fmt.Println(file.Filename)
		src, err := file.Open()
		hash, _ := ipfs.UploadIPFS(sh, src)

		c.String(http.StatusOK, fmt.Sprintf("文件上传成功: %s", hash))
		// QmSZ43EHSdzdYMYsMNhBMs6Ex6Gxgyn7BoxijPQe3bBRUG
	})

	router.GET("/view", func(c *gin.Context) {
		hash := c.Query("hash")
		s, _ := ipfs.CatIPFS(sh, hash)
		//c.File("./ipfs/tests/3.jpg")
		c.Writer.Write(s)
	})

	// 启动服务器
	router.Run(":8081")
}
