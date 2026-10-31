package link

import (
	"URLS/pkg/httpresponse"
	"errors"
	"fmt"
	"github.com/labstack/echo/v5"
	"net/http"
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
		return c.JSON(http.StatusBadRequest, httpresponse.Error("Incorrect POST request"))
	}

	ctx := c.Request().Context()
	link, err := h.svc.CreateLink(ctx, &hook)
	if err != nil {
		switch {
		case errors.Is(err, ErrCodeTaken):
			return c.JSON(http.StatusConflict, httpresponse.Error("Code already taken"))
		case errors.Is(err, ErrNotFound):
			return c.JSON(http.StatusNotFound, httpresponse.Error("Not found"))
		default:
			fmt.Println("Error:", err.Error())
			return c.JSON(http.StatusInternalServerError, httpresponse.Error("Server error"))
		}
	}

	return c.JSON(http.StatusCreated, map[string]any{"code": link.Code, "expires_at": link.Expires_at})
}
