package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sef-demo/http_models"
)

type UserServiceHTTPClient struct {
}

func GetUserServiceHTTPClientInstance() *UserServiceHTTPClient {
	return &UserServiceHTTPClient{}
}

func (u UserServiceHTTPClient) GetValid() {
	request := http_models.UserRequest{
		Name:        "Ali",
		Age:         "26",
		Username:    "AliMsayleb",
		PhoneNumber: "70123456",
		Email:       "ali@se.io",
		Roles:       []float64{1, 2, 3},
	}
	payload, err := json.Marshal(request)
	body := bytes.NewBuffer(payload)
	httpResponse, err := http.Post("http://127.0.0.1:7070/user", "application/json", body)
	if err != nil {
		fmt.Printf("Error HTTP calling the API: %s\n", err)
		return
	}
	if httpResponse.StatusCode != http.StatusOK {
		fmt.Printf("HTTP Request code %d\n", httpResponse.StatusCode)
		return
	}
	response := http_models.UserResponse{}
	bodyBytes, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Printf("Error unmarshaling into json %s\n", err)
		return
	}

	fmt.Printf("HTTP Request success, response username %s, resposne name %s\n", response.Username, response.Name)
}

func (u UserServiceHTTPClient) GetInvalid() {
	request := http_models.UserRequest{}
	payload, err := json.Marshal(request)
	body := bytes.NewBuffer(payload)
	httpResponse, err := http.Post("http://127.0.0.1:7070/user", "application/json", body)
	if err != nil {
		fmt.Printf("Error HTTP calling the API: %s\n", err)
		return
	}
	if httpResponse.StatusCode != http.StatusOK {
		fmt.Printf("HTTP Request code %d\n", httpResponse.StatusCode)
		return
	}
	response := http_models.UserResponse{}
	bodyBytes, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Printf("Error unmarshaling into json %s\n", err)
		return
	}

	fmt.Printf("HTTP Request success, response username %s\n", response.Username)
}
