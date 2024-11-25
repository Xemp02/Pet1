package app

import "github.com/Xemp02/Pet1/internal/transport"

const httpPort = ":8888"

func Run() {
	// start transport
	httpserver := transport.New(httpPort)
	httpserver.StartHttpServer()
	httpserver.AddRoute()

}
