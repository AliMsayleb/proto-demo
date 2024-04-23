package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"sef-demo/rpc"
	"strconv"
)

type UserRPCService struct {
	rpc.UserServiceServer
	rpc.CalculatorServiceServer
}

func (u UserRPCService) Get(_ context.Context, request *rpc.UserRequest) (*rpc.UserResponse, error) {
	validationError := validateRequest(request)
	if validationError != nil {
		return &rpc.UserResponse{}, validationError
	}
	newName := fmt.Sprintf("%s %s", request.Name, "SEF DEMO hello")
	return &rpc.UserResponse{
		Name:        newName,
		Age:         request.Age,
		Username:    request.Username,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Roles:       request.Roles,
	}, validationError
}

func (u UserRPCService) StreamGet(server rpc.UserService_StreamGetServer) error {
	ctx := server.Context()
	for {
		// exit if context is done
		select {
		case <-ctx.Done():
			{
				return ctx.Err()
			}
		default:
		}

		// receive data from stream
		request, err := server.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}
		validationError := validateRequest(request)
		if validationError != nil {
			return validationError
		}
		age, _ := strconv.Atoi(request.Age)

		err = server.Send(&rpc.UserResponse{
			Name:        request.Name,
			Age:         strconv.Itoa(age + 20),
			Username:    request.Username,
			PhoneNumber: request.PhoneNumber,
			Email:       request.Email,
			Roles:       request.Roles,
		})
		if err != nil {
			fmt.Printf("error sending message to client %s\n", err.Error())
		}
	}
}

func validateRequest(request *rpc.UserRequest) error {
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

func (u UserRPCService) Sum(server rpc.CalculatorService_SumServer) error {
	ctx := server.Context()
	var result float64 = 0
	for {
		// exit if context is done
		select {
		case <-ctx.Done():
			{
				break
			}
		default:
		}

		// receive data from stream
		request, err := server.Recv()
		if err == io.EOF {
			log.Println("Received EOF")
			break
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}
		result += request.Number
	}
	response := rpc.Number{Number: result}
	err := server.SendAndClose(&response)
	if err != nil {
		fmt.Printf("Failde to return response and close connection %v \n", err)
	}
	return nil
}
