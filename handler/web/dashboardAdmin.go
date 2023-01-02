package web

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type DashboardAdmin interface {
	Dashboard(c *gin.Context)
}

type dashboardAdmin struct {
}

func NewDashboardAdmin() *dashboardAdmin {
	return &dashboardAdmin{}
}

func (d *dashboardAdmin) Dashboard(c *gin.Context) {
	dashboard := path.Join("admin", "dashboardAdmin.tmpl")
	c.HTML(http.StatusOK, dashboard, nil)
}
