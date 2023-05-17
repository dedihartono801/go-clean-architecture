package handler

import (
	"github.com/dedihartono801/go-clean-architecture/infrastructure/customstatus"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/helper"
	"github.com/dedihartono801/go-clean-architecture/usecase/transaction"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler interface {
	Checkout(ctx *fiber.Ctx) error
}

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) TransactionHandler {
	return &transactionHandler{service}
}

// Create godoc
// @Summary      Checkout Items
// @Tags         checkout
// @Accept       json
// @Produce      json
// @Param		 raw	body	object		true	"body raw"
// @Success      201  {object} domain.Transaction
// @Security ApiKeyAuth
// @Router       /checkout [post]
func (h *transactionHandler) Checkout(ctx *fiber.Ctx) error {
	checkoutDto := new(transaction.CheckoutDto)
	if err := ctx.BodyParser(checkoutDto); err != nil {
		return helper.CustomResponse(ctx, nil, err.Error(), customstatus.ErrBadRequest.Code)
	}

	transaction, statusCode, err := h.service.Checkout(ctx, checkoutDto)
	if err != nil {
		return helper.CustomResponse(ctx, nil, err.Error(), statusCode)
	}
	return helper.CustomResponse(ctx, transaction, customstatus.StatusCreated.Message, statusCode)
}
