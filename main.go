package main

import (
	"rpc-http-json-golang/server"
)

// to test runing server, hit below curl code to terminal
// curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}' http://localhost:1234

func main() {
	server.StartServer()
}
