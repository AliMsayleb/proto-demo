package services

import (
	"context"
	"sef-demo/rpc"
)

type PizzaRPCService struct {
	rpc.PizzaServiceServer
}

func (p PizzaRPCService) GetName(_ context.Context, request *rpc.PizzaRequest) (*rpc.PizzaResponse, error) {
	pizzaName := ""
	if request.Vegetarian && request.Sauce == "red" {
		pizzaName = "Vegetarian"
		// pizza is vegetarian
	} else if request.Vegetarian {
		pizzaName = "Alfredo"
		// pizza is alfredo since sauce is not red and it has no meet
	} else if request.Sauce == "red" {
		pizzaName = "Peperoni"
		// peperoni
	} else {
		pizzaName = "Chicken Alfredo"
		// chicken alfredo
	}
	return &rpc.PizzaResponse{
		Name: pizzaName,
	}, nil
}
