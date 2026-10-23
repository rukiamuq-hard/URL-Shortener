package link

import (
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
	echo.POST("/links", h.CreateLink)
	//echo.GET("/:code" h.Redirect)        // redirect
	//echo.GET("/links/:code" h.Get)       // get
	//echo.PATCH("/links/:code" h.Update)  // update
	//echo.DELETE("/links/:code" h.Delete) // delete
}

func (h *Handler) CreateLink(c *echo.Context) error {
	var hook Link
	if err := c.Bind(&hook); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Status": "Incorrect POST request"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Created"})
}
