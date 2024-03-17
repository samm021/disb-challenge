package api

import (
	"disbursement-service/domain"
	"disbursement-service/dto"

	"github.com/gofiber/fiber/v2"
)

type transactionApi struct {
	transactionService domain.TransactionService
}

func NewTransaction(app *fiber.App, middleWare fiber.Handler, transactionService domain.TransactionService) {
	h := transactionApi{transactionService: transactionService}

	app.Post("api/transaction", middleWare, h.CreateTransaction)
	app.Post("api/transaction/callback", h.UpdateTransaction)
}

func (t transactionApi) CreateTransaction(ctx *fiber.Ctx) error {
	var req dto.TransactionReq

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(400).JSON(dto.Response{
			Message: "Invalid Body",
		})
	}

	res, err := t.transactionService.CreateTransaction(ctx.Context(), req)
	if err != nil {
		return ctx.Status(400).JSON(dto.Response{
			Message: "Invalid Body",
		})
	}

	return ctx.Status(200).JSON(res)
}

func (t transactionApi) UpdateTransaction(ctx *fiber.Ctx) error {
	var req dto.TransactionReq

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(400).JSON(dto.Response{
			Message: "Invalid Body",
		})
	}

	res, err := t.transactionService.CreateTransaction(ctx.Context(), req)
	if err != nil {
		return ctx.Status(400).JSON(dto.Response{
			Message: "Invalid Body",
		})
	}

	return ctx.Status(200).JSON(res)
}
