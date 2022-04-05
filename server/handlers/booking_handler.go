package handlers

import (
	"net/http"

	request "github.com/agungmohmd/booking-api/server/requests"
	"github.com/agungmohmd/booking-api/usecase"
	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	Handler
}

func (h *BookingHandler) SelectAll(ctx *fiber.Ctx) error {
	bookingUC := usecase.BookingUC{ContractUC: h.ContractUC}
	res, err := bookingUC.SelectAll()
	return h.SendResponse(ctx, res, err, 0)
}

func (h *BookingHandler) FindOne(ctx *fiber.Ctx) error {
	bookingUC := usecase.BookingUC{ContractUC: h.ContractUC}
	id := ctx.Params("id")
	res, err := bookingUC.FindOne(id)
	return h.SendResponse(ctx, res, err, 0)
}

func (h *BookingHandler) Add(ctx *fiber.Ctx) error {
	bookingUC := usecase.BookingUC{ContractUC: h.ContractUC}
	input := new(request.BookingRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, err, http.StatusBadRequest)
	}
	res, err := bookingUC.Add(input)
	return h.SendResponse(ctx, res, err, 0)
}
