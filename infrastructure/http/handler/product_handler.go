package handler

import (
	"github.com/dedihartono801/go-clean-architecture/infrastructure/customstatus"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/helper"
	"github.com/dedihartono801/go-clean-architecture/usecase/product"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	Product(ctx *fiber.Ctx) error
}

type productHandler struct {
	service product.Service
}

func NewFilmHandler(service product.Service) ProductHandler {
	return &productHandler{service: service}
}

// Update godoc
// @Summary      List Product
// @Tags         product
// @Accept       json
// @Produce      json
// @Success      200  {object} product.ProductDto
// @Security ApiKeyAuth
// @Router       /product [get]
func (h *productHandler) Product(ctx *fiber.Ctx) error {

	dt, statusCode, err := h.service.Product()
	if err != nil {
		return helper.CustomResponse(ctx, nil, err.Error(), statusCode)
	}

	return helper.CustomResponse(ctx, dt, customstatus.StatusOk.Message, statusCode)

}
