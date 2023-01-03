package web

import (
	"e-learning/client"
	"e-learning/utils"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthWeb interface {
	// User
	Login(c *gin.Context)
	LoginProses(c *gin.Context)

	Register(c *gin.Context)
	RegisterProses(c *gin.Context)

	Logout(c *gin.Context)

	// Admin
	LoginAdmin(c *gin.Context)
	LoginAdminProses(c *gin.Context)

	LogoutAdmin(c *gin.Context)
}

type authWeb struct {
	userClient   client.UserClient
	admindClient client.AdminClient
}

func NewAuthWeb(userClient client.UserClient, adminClient client.AdminClient) *authWeb {
	return &authWeb{userClient, adminClient}
}

// User Function
func (a *authWeb) Login(c *gin.Context) {
	login := path.Join("user", "loginUser.tmpl")
	c.HTML(http.StatusOK, login, nil)
}

func (a *authWeb) LoginProses(c *gin.Context) {
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	userId, status, err := a.userClient.Login(email, password)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	if status == 200 {
		token, err := utils.GenerateToken(strconv.Itoa(userId), "session_token")
		if err != nil {
			log.Println(err)
			http.Redirect(c.Writer, c.Request, "/user/login", http.StatusSeeOther)
			return
		}
		c.SetCookie("session_token", token, 60*60, "/", "", true, true)

		location := url.URL{Path: "/user/dashboard"}
		c.Redirect(http.StatusSeeOther, location.RequestURI())
	} else {
		http.Redirect(c.Writer, c.Request, "/user/login", http.StatusSeeOther)
	}
}

func (a *authWeb) Register(c *gin.Context) {
	register := path.Join("user", "registerUser.tmpl")
	c.HTML(http.StatusOK, register, nil)
}

func (a *authWeb) RegisterProses(c *gin.Context) {
	fullname := c.Request.FormValue("fullname")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	userId, status, err := a.userClient.Register(fullname, email, password)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	if status == 201 {
		token, err := utils.GenerateToken(strconv.Itoa(userId), "session_token")
		if err != nil {
			log.Println(err)
			http.Redirect(c.Writer, c.Request, "/user/register", http.StatusSeeOther)
			return
		}
		c.SetCookie("session_token", token, 60*60, "/", "", true, true)
		location := url.URL{Path: "/user/dashboard"}
		c.Redirect(http.StatusSeeOther, location.RequestURI())
	} else {
		location := url.URL{Path: "/user/register"}
		c.Redirect(http.StatusSeeOther, location.RequestURI())
	}
}

func (a *authWeb) Logout(c *gin.Context) {
	c.SetCookie("session_token", "", -1, "/", "", true, true)

	location := url.URL{Path: "/"}
	c.Redirect(http.StatusSeeOther, location.RequestURI())
}

// Admin Function
func (a *authWeb) LoginAdmin(c *gin.Context) {
	login := path.Join("admin", "loginAdmin.tmpl")
	c.HTML(http.StatusOK, login, nil)
}

func (a *authWeb) LoginAdminProses(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	userId, status, err := a.admindClient.Login(username, password)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	if status == 200 {
		token, err := utils.GenerateToken(strconv.Itoa(userId), "session_token")
		if err != nil {
			log.Println(err)
			http.Redirect(c.Writer, c.Request, "/admin/login", http.StatusSeeOther)
			return
		}
		c.SetCookie("session_token", token, 60*60, "/", "", true, true)

		location := url.URL{Path: "/admin/dashboard"}
		c.Redirect(http.StatusSeeOther, location.RequestURI())
	} else {
		log.Println(err)
		http.Redirect(c.Writer, c.Request, "/admin/login", http.StatusSeeOther)
		return
	}

}

func (a *authWeb) LogoutAdmin(c *gin.Context) {

}
