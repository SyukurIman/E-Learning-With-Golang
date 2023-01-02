package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type AdminClient interface {
	Login(username, password string) (userId int, respCode int, err error)
}

type adminClient struct{}

func NewAdminClient() *adminClient {
	return &adminClient{}
}

func (a *adminClient) Login(username, password string) (userId int, respCode int, err error) {
	dataJson := map[string]string{
		"username": username,
		"password": password,
	}

	data, err := json.Marshal(dataJson)
	if err != nil {
		return 0, 0, err
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/admin/v1/login", bytes.NewBuffer(data))
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
		if result["admin_id"] != nil {
			return int(result["admin_id"].(float64)), http.StatusOK, nil
		} else {
			return 0, resp.StatusCode, nil
		}
	}
}
