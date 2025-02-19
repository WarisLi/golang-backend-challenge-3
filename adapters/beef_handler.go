package adapters

import (
	"github.com/WarisLi/golang-backend-challenge-3/core"
	"github.com/gofiber/fiber/v2"
)

type HttpBeefHandler struct {
	service core.BeefService
}

func NewHttpBeefHandler(service core.BeefService) *HttpBeefHandler {
	return &HttpBeefHandler{service: service}
}

func (h *HttpBeefHandler) GetBeefs(c *fiber.Ctx) error {
	beefs, err := h.service.Summary()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(beefs)
}
