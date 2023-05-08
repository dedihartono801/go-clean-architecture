package handler

import (
	"github.com/dedihartono801/go-clean-architecture/infrastructure/customstatus"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/helper"
	"github.com/dedihartono801/go-clean-architecture/usecase/sku"
	"github.com/gofiber/fiber/v2"
)

type SkuHandler interface {
	List(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

type skuHandler struct {
	service sku.Service
}

func NewSkuHandler(service sku.Service) SkuHandler {
	return &skuHandler{service}
}

// List godoc
// @Summary      List sku
// @Tags         skus
// @Accept       json
// @Produce      json
// @Success      200  {array} domain.Sku
// @Security ApiKeyAuth
// @Router       /sku [get]
func (h *skuHandler) List(ctx *fiber.Ctx) error {
	skus, err := h.service.List()
	if err != nil {
		return helper.CustomResponse(ctx, nil, err.Error(), customstatus.ErrInternalServerError.Code)
	}
	return helper.CustomResponse(ctx, skus, customstatus.StatusOk.Message, customstatus.StatusOk.Code)
}

// Create godoc
// @Summary      Create sku
// @Tags         skus
// @Accept       json
// @Produce      json
// @Param		 raw	body	object		true	"body raw"
// @Success      200  {object} domain.Sku
// @Security ApiKeyAuth
// @Router       /sku [post]
func (h *skuHandler) Create(ctx *fiber.Ctx) error {
	skuDto := new(sku.CreateDto)
	if err := ctx.BodyParser(skuDto); err != nil {
		return helper.CustomResponse(ctx, nil, err.Error(), customstatus.ErrBadRequest.Code)
	}

	sku, statusCode, err := h.service.Create(skuDto)
	if err != nil {
		return helper.CustomResponse(ctx, nil, err.Error(), statusCode)
	}
	return helper.CustomResponse(ctx, sku, customstatus.StatusCreated.Message, statusCode)
}
