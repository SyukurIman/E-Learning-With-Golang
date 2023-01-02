package api

import (
	"e-learning/entity"
	"e-learning/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminAPI interface {
	Login(c *gin.Context)
}

type adminAPI struct {
	adminService service.AdminService
}

func NewAdminAPI(adminService service.AdminService) *adminAPI {
	return &adminAPI{adminService}
}

func (a *adminAPI) Login(c *gin.Context) {
	var admin entity.AdminLogin
	err := c.BindJSON(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error Bad Request")
		return
	}

	if admin.Username == "" || admin.Password == "" {
		c.JSON(http.StatusBadRequest, "Error Bad Request")
		return
	}

	newAdmin := entity.Admin{
		Username: admin.Username,
		Password: admin.Password,
	}
	newID, err := a.adminService.Login(c.Request.Context(), &newAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error Internal Server")
	}
	id := float64(newID)
	c.JSON(http.StatusOK, map[string]interface{}{
		"admin_id": id,
		"message":  "login succes",
	})
}
