package transport

import (
	"github.com/labstack/echo/v4"
	"log"
)

type HttpServer struct {
	*echo.Echo
	address string
}

func New(address string) *HttpServer {
	echoServer := echo.New()

	return &HttpServer{
		address: address,
		Echo:    echoServer,
	}
}

func (h *HttpServer) StartHttpServer() {
	err := h.Start(h.address)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *HttpServer) AddRoute() {
	version := h.Group("/v1")
	{
		version.GET("/authService", h.autharization)

		version.GET("/sendMessage", h.sendMessage)
	}
}
