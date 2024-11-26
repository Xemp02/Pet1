package transport

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *HttpServer) autharization(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *HttpServer) sendMessage(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
