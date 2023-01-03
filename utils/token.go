package utils

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type TokenData struct {
	Data string `json:"data"`
	jwt.StandardClaims
}

func GenerateToken(data string, typeKey string) (string, error) {
	claim := TokenData{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "New-App",
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		Data: data,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenstring, err := token.SignedString([]byte(typeKey))
	if err != nil {
		return "", err
	}

	return tokenstring, err
}

func ClaimsToken(tokenString string, typeKey string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenData{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(typeKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	payload := token.Claims.(*TokenData)
	var result = payload.Data
	return result, nil
}

func Authentication(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerType := c.Request.Header.Get("Content-Type")
		cookie, err := c.Cookie("session_token")
		if err != nil {
			if headerType == "application/json" {
				c.JSON(http.StatusUnauthorized, "Error")
				return
			} else {
				location := url.URL{Path: path}
				c.Redirect(http.StatusSeeOther, location.RequestURI())
				return
			}
		}

		id, err := ClaimsToken(cookie, "session_token")
		if err != nil {
			location := url.URL{Path: path}
			c.Redirect(http.StatusSeeOther, location.RequestURI())
			return
		}

		c.Set("id", id)
		c.Next()
	}
}
