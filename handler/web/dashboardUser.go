package web

import (
	"e-learning/repository"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DashboardUser interface {
	Dashboard(c *gin.Context)
}

type dashboardUser struct {
	userRepository repository.UserRepository
}

func NewDashboardUser(userRepository repository.UserRepository) *dashboardUser {
	return &dashboardUser{userRepository}
}

func (d *dashboardUser) Dashboard(c *gin.Context) {
	userId := c.GetString("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	user, err := d.userRepository.GetUserById(c.Request.Context(), id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	data := map[string]interface{}{
		"name": user.Fullname,
	}

	dashboard := path.Join("user", "dashboardUser.tmpl")
	c.HTML(http.StatusOK, dashboard, data)
}
