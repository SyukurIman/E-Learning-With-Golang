package api

import (
	"e-learning/entity"
	"e-learning/service"
	"e-learning/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminAPI interface {
	Login(c *gin.Context)
	AddNewAdmin(c *gin.Context)
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
func (a *adminAPI) AddNewAdmin(c *gin.Context) {
	var admin entity.Admin

	err := c.BindJSON(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if admin.Username == "" || admin.Password == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Username and Password Empty",
		})
		return
	}

	password, err := utils.HashPassword(admin.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	admin.Password = password

	admin, err = a.adminService.Register(c.Request.Context(), &admin)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error internal server",
		})
		return
	}
	id := float64(admin.ID)
	c.JSON(http.StatusCreated, map[string]interface{}{
		"admin_id": id,
		"message":  "register admin succes",
	})
}
