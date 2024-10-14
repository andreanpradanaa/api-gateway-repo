package delivery

import (
	"github.com/labstack/echo/v4"
)

type HttpHandler struct {
	router *echo.Echo
}

func NewHttpHandler() *HttpHandler {
	server := &HttpHandler{
		router: echo.New(),
	}

	server.setupRouter()
	return server
}

func (handler *HttpHandler) setupRouter() {
	publicAccess := handler.router.Group("/api/v1")
	publicAccess.GET("/products", handler.GetProduct)

	// TODO: add new router here
}

func (handler *HttpHandler) Start() error {
	return handler.router.Start(":8080")
}
