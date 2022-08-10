package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type PageCtr interface {
	Index(ctx *gin.Context)
}

type pageCtr struct {
}

func NewPageCtr() PageCtr {
	return &pageCtr{}
}

func (c *pageCtr) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "notes/index", gin.H{
		"title": "Latihan Golang",
		"url":   os.Getenv("URL"),
	})
}
