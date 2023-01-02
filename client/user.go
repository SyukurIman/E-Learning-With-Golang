package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type UserClient interface {
	Login(email, password string) (userId int, respCode int, err error)
	Register(fullname, email, password string) (userId int, respCode int, err error)
}

type userClient struct{}

func NewUserClient() *userClient {
	return &userClient{}
}

func (u *userClient) Login(email, password string) (userId int, respCode int, err error) {
	dataJson := map[string]string{
		"email":    email,
		"password": password,
	}

	data, err := json.Marshal(dataJson)
	if err != nil {
		return 0, 0, err
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/user/login", bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return 0, -1, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	var result map[string]interface{}

	if err != nil {
		log.Println(err)
		return 0, -1, err
	} else {
		json.Unmarshal(b, &result)
		if result["user_id"] != nil {
			return int(result["user_id"].(float64)), http.StatusOK, nil
		} else {
			return 0, resp.StatusCode, nil
		}
	}
}

func (u *userClient) Register(fullname, email, password string) (userId int, respCode int, err error) {
	dataJson := map[string]string{
		"fullname": fullname,
		"email":    email,
		"password": password,
	}
	data, err := json.Marshal(dataJson)
	if err != nil {
		return 0, -1, err
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/user/register", bytes.NewBuffer(data))
	if err != nil {
		return 0, -1, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return 0, -1, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	var result map[string]interface{}

	if err != nil {
		return 0, -1, err
	} else {
		json.Unmarshal(b, &result)

		if result["user_id"] != nil {
			return int(result["user_id"].(float64)), resp.StatusCode, nil
		} else {
			return 0, resp.StatusCode, nil
		}
	}
}
