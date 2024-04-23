package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"sef-demo/rpc"
	"strconv"
)

type UserServiceClient struct {
	rpc.UserServiceClient
	rpc.CalculatorServiceClient
}

//class UserServiceClient {
//	private rpc.UserServiceClient $client;
//
//	$response = $client->get($request);
//}

func GetUserServiceRPCClientInstance() *UserServiceClient {
	connection, _ := grpc.Dial(
		"localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return &UserServiceClient{
		rpc.NewUserServiceClient(connection),
		rpc.NewCalculatorServiceClient(connection),
	}
}

func (u UserServiceClient) GetValid() {
	request := rpc.UserRequest{
		Name:        "Charbel",
		Age:         "26",
		Username:    "AliMsayleb",
		PhoneNumber: "70123456",
		Email:       "ali@se.io",
		Roles:       []float64{1, 2, 3},
	}
	ctx := context.Background()
	response, err := u.Get(ctx, &request)
	// [$response, $error] = $this->get($ctx, $request);
	if err != nil {
		fmt.Printf("RPC Request error %s\n", err)
		return
	}

	fmt.Printf("RPC Request success, response Name %s request name %s\n", response.Name, request.Name)
}

func (u UserServiceClient) GetInvalid() {
	request := rpc.UserRequest{
		Name:        "Ali",
		Age:         "",
		Username:    "",
		PhoneNumber: "",
		Email:       "ali@se.io",
		Roles:       []float64{1, 2, 3},
	}
	ctx := context.Background()
	response, err := u.Get(ctx, &request)
	if err != nil {
		fmt.Printf("RPC Request error %s\n", err)
		return
	}

	fmt.Printf("RPC Request success, response username %s\n", response.Username)
}

func (u UserServiceClient) Stream() {
	ctx := context.Background()
	client, err := u.StreamGet(ctx)
	if err != nil {
		fmt.Printf("Error creating stream client %s\n", err.Error())
		return
	}
	for i := 0; i < 50000; i++ {
		err = client.Send(&rpc.UserRequest{
			Name:        "Ali",
			Age:         strconv.Itoa(i),
			Username:    "AliMsayleb",
			PhoneNumber: "70123456",
			Email:       "ali@se.io",
			Roles:       []float64{1, 2, 3},
		})
		if err != nil {
			fmt.Printf("Error sending stream request %s\n", err.Error())
			break
		}
		fmt.Printf("Stream request sent\n")
		response, err := client.Recv()
		if err != nil {
			fmt.Printf("Error receiving stream response %s\n", err.Error())
			break
		}
		fmt.Printf("Received response of user with age %s\n", response.Age)
	}
	ctx.Done()
	fmt.Printf("Done streaming\n")
}

func (u UserServiceClient) Calculate() {
	ctx := context.Background()
	client, err := u.Sum(ctx)
	if err != nil {
		fmt.Printf("Error creating stream client %s\n", err.Error())
		return
	}
	for i := 0; i < 10; i++ {
		randomNumber := rand.Intn(20)
		err = client.Send(&rpc.Number{
			Number: float64(randomNumber),
		})
		if err != nil {
			fmt.Printf("Error sending stream request %s\n", err.Error())
			break
		}
		fmt.Printf("Stream request sent with random number %d\n", randomNumber)
	}
	finalNumber, err := client.CloseAndRecv()
	ctx.Done()
	if err != nil {
		fmt.Printf("Error receiving final response %v\n", err)
		return
	}
	fmt.Printf("Done calculating, the final number is %d\n", int(finalNumber.Number))
}
