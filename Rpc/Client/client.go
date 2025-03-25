package main

import (
	"fmt"
	"net/rpc"
)

type Response struct {
	Message string // Ensure this matches the server
}

type LikeRequest struct {
	Name  string
	Likes int
}

type Request struct {
	Name string // Ensure this matches the server
	Cost int    // Ensure this matches the server
}

type MenuResponse struct {
	Prostitute map[string]int // Ensure this matches the server
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error in dialing:", err)
		return
	}
	defer client.Close()

	var res Response
	var req Request
	var likeReq LikeRequest
	var menures MenuResponse
	var choice int

	for {
		var name string
		var cost int

		fmt.Println("Enter 1 to Rent Prostitute, 2 to Like Prostitute, 3 to see Prostitute catalogue or -1 to exit:")
		fmt.Scan(&choice)

		if choice == -1 {
			break // Exit the loop if user inputs -1
		}

		switch choice {
		case 1:
			fmt.Println("Enter prostitute name:")
			fmt.Scan(&name)

			fmt.Println("Enter available fund:")
			fmt.Scan(&cost)
			req = Request{Name: name, Cost: cost}
			err = client.Call("Brothel.RentProstitute", req, &res)
			if err != nil {
				fmt.Println("Error in RPC call:", err)
				return
			}
			fmt.Println("Response:", res.Message)
		case 2:
			fmt.Println("Enter prostitute name:")
			fmt.Scan(&name)

			likeReq = LikeRequest{Name: name, Likes: 1}
			err = client.Call("Brothel.LikeProstitute", likeReq, &res) // Use likeReq for liking
			if err != nil {
				fmt.Println("Error in RPC call:", err)
				return
			}
			fmt.Println("Response:", res.Message)
		case 3:
			err = client.Call("Brothel.ShowCatalogue", struct{}{}, &menures) // Ensure method name matches server
			if err != nil {
				fmt.Println("Error in RPC call:", err)
				return
			}
			fmt.Println("Prostitute Catalogue:", menures.Prostitute)
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or -1.")
		}
	}
}