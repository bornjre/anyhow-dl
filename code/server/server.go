package server

import (
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ServerRun() {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {

	})

	r.GET("/:file", func(ctx *gin.Context) {

		file := ctx.Param("file")
		if strings.HasSuffix(file, ".torrent") {

		} else {

		}

	})

	r.GET("/:folder/:file", func(ctx *gin.Context) {

	})

	r.NoRoute(func(ctx *gin.Context) {

	})

	r.Run()
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
