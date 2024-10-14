package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (handler *HttpHandler) GetProduct(c echo.Context) error {
	return c.String(http.StatusOK, "sukses get all data")
}
