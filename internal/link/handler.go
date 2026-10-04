package link

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(echo *echo.Echo) {
	echo.POST("/links", h.HelloWorld)
	//echo.GET("/:code" h.Redirect)        // redirect
	//echo.GET("/links/:code" h.Get)       // get
	//echo.PATCH("/links/:code" h.Update)  // update
	//echo.DELETE("/links/:code" h.Delete) // delete
}

func (h *Handler) HelloWorld(c *echo.Context) error {
	fmt.Println("hello world")
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello world"})
}
