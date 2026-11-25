package link

import (
	"URLS/pkg/httpresponse"
	"errors"
	"github.com/labstack/echo/v5"
	"log/slog"
	"net/http"
)

const liveTimeLink = 21600

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(e *echo.Echo) {
	e.POST("/links", h.CreateLink)
	e.GET("/:code", h.Redirect)
	e.GET("/links/:code", h.GetCreatedLink)
	e.PATCH("/links/:code", h.Update)  // update
	e.DELETE("/links/:code", h.Delete) // delete
}

func (h *Handler) fail(c *echo.Context, err error) error { // for processing errors
	switch {
	case errors.Is(err, ErrInvalidInput):
		return c.JSON(http.StatusBadRequest, httpresponse.Error("Bad request"))
	case errors.Is(err, ErrCodeTaken):
		return c.JSON(http.StatusConflict, httpresponse.Error("Code already taken"))
	case errors.Is(err, ErrNotFound):
		return c.JSON(http.StatusNotFound, httpresponse.Error("Not found"))
	case errors.Is(err, ErrExpired):
		return c.JSON(http.StatusGone, httpresponse.Error("Link expired"))
	default:
		slog.Error("link handler", "err", err)
		return c.JSON(http.StatusInternalServerError, httpresponse.Error("Server error"))
	}
}

func (h *Handler) CreateLink(c *echo.Context) error {
	var hook BindLink
	if err := c.Bind(&hook); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error("Incorrect POST request"))
	}

	if hook.Orig_url == nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error("Incorrect POST request"))
	}

	if hook.Duration > liveTimeLink {
		return c.JSON(http.StatusBadRequest, httpresponse.Error("Out of duration range"))
	}

	monet := false
	if hook.Monet_enabled != nil {
		monet = *hook.Monet_enabled
	}

	Link := Link{
		Orig_url:      *hook.Orig_url,
		Monet_enabled: monet,
		Duration:      hook.Duration,
	}

	link, err := h.svc.CreateLink(c.Request().Context(), &Link)
	if err != nil {
		return h.fail(c, err)
	}

	return c.JSON(http.StatusCreated, map[string]any{"code": link.Code, "expires_at": link.Expires_at})
}

func (h *Handler) Redirect(c *echo.Context) error {
	code := c.Param("code")

	val, err := h.svc.GetLinkInfo(c.Request().Context(), code)
	if err != nil {
		return h.fail(c, err)
	}

	return c.Redirect(http.StatusFound, val.Orig_url)
}

func (h *Handler) GetCreatedLink(c *echo.Context) error {

	code := c.Param("code")

	val, err := h.svc.GetLinkInfo(c.Request().Context(), code)
	if err != nil {
		return h.fail(c, err)
	}

	return c.JSON(http.StatusOK, val)
}

func (h *Handler) Delete(c *echo.Context) error {
	code := c.Param("code")

	result, err := h.svc.Delete(c.Request().Context(), code)
	if err != nil {
		return h.fail(c, err)
	}

	if result == 0 {
		return c.JSON(http.StatusBadRequest, httpresponse.Error("Nothing to DELETE"))
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) Update(c *echo.Context) error {
	var hook BindLink

	if err := c.Bind(&hook); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error("Incorrect PATCH request"))
	}

	if hook.Duration > liveTimeLink {
		return c.JSON(http.StatusBadRequest, httpresponse.Error("Out of duration range"))
	}

	code := c.Param("code")

	PLink := PatchLink{
		Code:          code,
		Orig_url:      hook.Orig_url,
		Monet_enabled: hook.Monet_enabled,
		Duration:      hook.Duration,
	}

	result, err := h.svc.Update(c.Request().Context(), &PLink)
	if err != nil {
		return h.fail(c, err)
	}
	if result == 0 {
		return c.JSON(http.StatusBadRequest, httpresponse.Error("Nothing to PATCH"))
	}
	return c.NoContent(http.StatusOK)
}
