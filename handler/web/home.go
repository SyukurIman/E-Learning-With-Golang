package web

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type HomeWeb interface {
	Index(c *gin.Context)
}

type homeWeb struct{}

func NewHomeWeb() *homeWeb {
	return &homeWeb{}
}

func (h *homeWeb) Index(c *gin.Context) {
	index := path.Join("general", "index.tmpl")
	c.HTML(http.StatusOK, index, nil)
}
