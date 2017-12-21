package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"test-golang/contracts"
)

const port = 1234

func StartServer() {
	helloworld := new(HelloWorldHandler)
	rpc.Register(helloworld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable start server on port: %v", err))
	}

	http.Serve(l, http.HandlerFunc(httpHandler))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
	err := rpc.ServeRequest(serverCodec)
	if err != nil {
		log.Printf("Error while serving JSON request: %v", err)
		http.Error(w, "Error while serving JSON request, detail has been logged.", 500)
	}
	return
}

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contracts.Requests, reply *contracts.Responses) error {
	reply.Message = "hallo from rpc, " + args.Name
	return nil
}

func (c *HttpConn) Read(p []byte) (n int, err error) {
	return c.in.Read(p)
}

func (c *HttpConn) Write(d []byte) (n int, err error) {
	return c.out.Write(d)
}

func (c *HttpConn) Close() error {
	return nil
}
