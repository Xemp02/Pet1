package transport

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Usecase_1 interface {
	// Описание функционала слоя бизнес логики
}

func (h *HttpServer) authService(c echo.Context) error {

	// Проверка ошибки с ниже уровня

	response := fmt.Sprintf(`
		<html>
		
		<html>
		`)

	// Страница аунтификации + возможность регистрации , при отсутсвии токена/сеанса
	return c.HTML(http.StatusOK, response)

}

func (h *HttpServer) sendMessage(c echo.Context) error {

	// Проверка ошибки с ниже уровня

	response := fmt.Sprintf(` 
		<html>
		
		<html>
		`)
	// Основная страница с функционалом - отправка сообщений в обищй чат.

	return c.HTML(http.StatusOK, response)

}
