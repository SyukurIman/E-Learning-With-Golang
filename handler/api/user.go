package api

import (
	"e-learning/entity"
	"e-learning/service"
	"e-learning/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Login(c *gin.Context) {
	var user entity.UserLogin

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid decode json")
		return
	}

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, "Email or Password is Empty")
		return
	}

	newUser := entity.User{
		Email:    user.Email,
		Password: user.Password,
	}
	newID, err := u.userService.Login(c.Request.Context(), &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error internal server")
		return
	}
	id := float64(newID)
	c.JSON(http.StatusOK, map[string]any{
		"user_id": id,
		"message": "login success",
	})
}

func (u *userAPI) Register(c *gin.Context) {
	var user entity.UserRegister

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if user.Fullname == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"Error": "Email or Password is Empty"})
		return
	}

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	NewUser := entity.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: password,
	}

	NewUser, err = u.userService.Register(c.Request.Context(), &NewUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "error internal server"})
		return
	}

	id := float64(NewUser.ID)
	c.JSON(http.StatusCreated, map[string]interface{}{"user_id": id, "message": "register succes"})
}
