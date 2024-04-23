package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sef-demo/rpc"
)

type PizzaClient struct {
	rpc.PizzaServiceClient
}

func (p PizzaClient) GetPizzaName(sauce string, isVegetarian bool) {
	ctx := context.Background()
	request := rpc.PizzaRequest{
		Sauce:      sauce,
		Vegetarian: isVegetarian,
	}
	response, err := p.GetName(ctx, &request)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}
	vegetarianPhrase := "is not vegetarian"
	if isVegetarian {
		vegetarianPhrase = "is vegetarian"
	}
	fmt.Printf("Pizza with sauce %s and %s is called %s\n", sauce, vegetarianPhrase, response.Name)
}

func GetPizzaRPCClientInstance() *PizzaClient {
	connection, _ := grpc.Dial(
		"localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return &PizzaClient{
		rpc.NewPizzaServiceClient(connection),
	}
}
