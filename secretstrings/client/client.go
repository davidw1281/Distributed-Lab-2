package main

import (
	"flag"
	"net/rpc"

	"bufio"
	"fmt"
	"os"

	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	//TODO: connect to the RPC server and send the request(s)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	request := stubs.Request{Message: ""}
	response := new(stubs.Response)

	file, _ := os.Open("/home/davidw128/Distributed-Lab-2/secretstrings/wordlist")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		request.Message = scanner.Text()
		client.Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println("Responded " + response.Message)
	}
}
