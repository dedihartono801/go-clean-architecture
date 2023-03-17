package handler

import (
	"github.com/dedihartono801/go-clean-architecture/infrastructure/customstatus"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/helper"
	"github.com/dedihartono801/go-clean-architecture/usecase/admin"
	"github.com/gofiber/fiber/v2"
)

type AdminHandler interface {
	Login(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

type adminHandler struct {
	service admin.Service
}

func NewAdminHandler(service admin.Service) AdminHandler {
	return &adminHandler{service}
}

// Find godoc
// @Summary      Get Profile Admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Success      200  {object} domain.Admin
// @Security ApiKeyAuth
// @Router       /admin [get]
func (h *adminHandler) Find(ctx *fiber.Ctx) error {
	admin, err := h.service.Find(ctx.Locals("adminID").(string))
	if err != nil {
		return helper.CustomResponse(ctx, nil, customstatus.ErrNotFound.Message, customstatus.ErrNotFound.Code)
	}
	return helper.CustomResponse(ctx, admin, customstatus.StatusOk.Message, customstatus.StatusOk.Code)
}

// Create godoc
// @Summary      Create Admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param		 raw	body	object		true	"body raw"
// @Success      200  {object} domain.Admin
// @Router       /admin/create [post]
func (h *adminHandler) Create(ctx *fiber.Ctx) error {
	adminDto := new(admin.CreateDto)
	if err := ctx.BodyParser(adminDto); err != nil {
		return helper.CustomResponse(ctx, nil, err, 500)
	}

	admin, statusCode, err := h.service.Create(adminDto)
	if err != nil {
		return helper.CustomResponse(ctx, nil, err.Error(), statusCode)
	}
	return helper.CustomResponse(ctx, admin, customstatus.StatusCreated.Message, customstatus.StatusCreated.Code)
}

// Update godoc
// @Summary      Login admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param		 raw	body	object	true	"body raw"
// @Success      200  {object} domain.Admin
// @Router       /admin/login [post]
func (h *adminHandler) Login(ctx *fiber.Ctx) error {
	adminDto := new(admin.LoginDto)
	if err := ctx.BodyParser(adminDto); err != nil {
		return err
	}

	result, statusCode, err := h.service.Login(adminDto)
	if err != nil {
		return helper.CustomResponse(ctx, nil, err.Error(), statusCode)
	}

	return helper.CustomResponse(ctx, result, customstatus.StatusOk.Message, customstatus.StatusOk.Code)

}
