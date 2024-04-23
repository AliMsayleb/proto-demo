package main

import (
	"fmt"
	"sef-demo/client"
	"sef-demo/server"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello world")
	server.StartGRPC()
	server.StartHTTP()
	rpcClientInstance := client.GetUserServiceRPCClientInstance()
	httpClientInstance := client.GetUserServiceHTTPClientInstance()
	rpcPizzaClientInstance := client.GetPizzaRPCClientInstance()
	run(rpcClientInstance, httpClientInstance, rpcPizzaClientInstance)
}

func run(rpcClientInstace *client.UserServiceClient, httpClientInstance *client.UserServiceHTTPClient, instance *client.PizzaClient) {
	input := ""
	for {
		printMenu()
		fmt.Scanln(&input)
		if strings.TrimSpace(input) == "0" {
			break
		}
		if strings.TrimSpace(input) == "1" {
			httpClientInstance.GetValid()
		}
		if strings.TrimSpace(input) == "2" {
			httpClientInstance.GetInvalid()
		}
		if strings.TrimSpace(input) == "3" {
			rpcClientInstace.GetValid()
		}
		if strings.TrimSpace(input) == "4" {
			rpcClientInstace.GetInvalid()
		}
		if strings.TrimSpace(input) == "5" {
			start := time.Now()
			rpcClientInstace.Stream()
			finish := time.Now()
			fmt.Printf("Streaming took %d milliseconds\n", finish.Sub(start).Milliseconds())
		}
		if strings.TrimSpace(input) == "6" {
			start := time.Now()
			for i := 0; i < 50000; i++ {
				rpcClientInstace.GetValid()
			}
			finish := time.Now()
			fmt.Printf("HTTP Loop took %d milliseconds\n", finish.Sub(start).Milliseconds())
		}
		if strings.TrimSpace(input) == "7" {
			rpcClientInstace.Calculate()
		}
		if strings.TrimSpace(input) == "8" {
			fmt.Println("Enter 1 for red sauce, anything else for non red sauce")
			var sauce string
			fmt.Scanln(&sauce)
			sauceString := "not_red"
			if strings.TrimSpace(sauce) == "1" {
				sauceString = "red"
			}
			fmt.Println("Enter 1 for vegetarian, anything else for non vegetarian")
			var isVegString string
			fmt.Scanln(&isVegString)
			isVeg := false
			if strings.TrimSpace(isVegString) == "1" {
				isVeg = true
			}
			instance.GetPizzaName(sauceString, isVeg)
		}
	}
}

func printMenu() {
	fmt.Println("Enter 0 to terminate")
	fmt.Println("Enter 1 to call http valid request")
	fmt.Println("Enter 2 to call http invaliad request")
	fmt.Println("Enter 3 to call rpc valiad request")
	fmt.Println("Enter 4 to call rpc invaliad request")
	fmt.Println("Enter 5 to call rpc stream call of 5000 objects")
	fmt.Println("Enter 6 to call http 5000 times")
	fmt.Println("Enter 7 to generate 10 random numbers and sum them by the server")
	fmt.Println("Enter 8 to ask about pizza")
}
