package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

var (
	mu sync.Mutex
)

type LikeRequest struct {
	Name  string
	Likes int
}

type Response struct {
	Message string
}

type Request struct {
	Name string
	Cost int
}

type MenuResponse struct {
	Prostitute map[string]int
}

type Brothel struct {
	Prostitute map[string]int
}

func (b *Brothel) RentProstitute(req *Request, res *Response) error {
	mu.Lock()
	defer mu.Unlock()

	name := req.Name
	giveCost := req.Cost
	cost, exist := b.Prostitute[name]

	if !exist {
		return fmt.Errorf("Prostitute not found")
	}
	if giveCost < cost {
		return fmt.Errorf("Prostitute %s is too expensive for you", name)
	} else {
		b.Prostitute[name]--
		res.Message = "Prostitute " + name + " rented successfully. Enjoy!"
		return nil
	}
}

func (b *Brothel) ShowCatalogue(args struct{}, res *MenuResponse) error { // Corrected method name
	res.Prostitute = b.Prostitute
	return nil
}

func (b *Brothel) LikeProstitute(req *LikeRequest, res *Response) error {
	name := req.Name
	likes := req.Likes

	cost, exist := b.Prostitute[name]

	if !exist {
		return fmt.Errorf("Prostitute not available")
	} else {
		b.Prostitute[name] = cost + (likes * 10)
		res.Message = "Updated. Thank you for your likes. This prostitute now costs " + fmt.Sprintf("%d", (cost+(likes*10)))
		return nil
	}
}

func main() {
	fmt.Println("Server running")
	brothel := &Brothel{
		Prostitute: map[string]int{
			"Ms":   25,
			"Pras": 1000,
			"Ram":  1,
			"Shek": 0,
			"bitch":8
		},
	}

	err := rpc.Register(brothel)
	if err != nil {
		fmt.Println("Error registering RPC service:", err)
		return
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error in listening:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}