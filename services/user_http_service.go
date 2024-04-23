package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sef-demo/http_models"
)

type UserHTTPService struct{}

func (u UserHTTPService) GetUser(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error unmarshalling request %s", err)))
		return
	}
	request := http_models.UserRequest{}
	err = json.Unmarshal(bodyBytes, &request)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error unmarshalling request %s", err)))
		return
	}
	validationError := validateHTTPRequest(request)
	if validationError != nil {
		io.WriteString(w, fmt.Sprintf("%s", validationError))
		return
	}
	modifiedName := fmt.Sprintf("%s %s", request.Name, "Ali modified me")
	responseObject := http_models.UserResponse{
		Name:        modifiedName,
		Age:         request.Age,
		Username:    request.Username,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Roles:       request.Roles,
	}
	response, err := json.Marshal(responseObject)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Error marshalling response %s", err))
		return
	}
	io.WriteString(w, string(response))
}

func validateHTTPRequest(request http_models.UserRequest) error {
	errorString := ""
	if request.Name == "" {
		errorString = fmt.Sprintf("%sField Name cannot be empty\n", errorString)
	}
	if request.Age == "" {
		errorString = fmt.Sprintf("%sField Age cannot be empty\n", errorString)
	}
	if request.Username == "" {
		errorString = fmt.Sprintf("%sField Username cannot be empty\n", errorString)
	}
	if request.PhoneNumber == "" {
		errorString = fmt.Sprintf("%sField PhoneNumber cannot be empty\n", errorString)
	}
	if request.Email == "" {
		errorString = fmt.Sprintf("%sField Email cannot be empty\n", errorString)
	}
	if len(request.Roles) == 0 {
		errorString = fmt.Sprintf("%sField Roles cannot be empty\n", errorString)
	}
	if errorString == "" {
		return nil
	}
	return errors.New(errorString)
}
