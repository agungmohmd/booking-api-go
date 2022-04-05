package routes

import (
	"github.com/agungmohmd/booking-api/server/handlers"
	"github.com/gofiber/fiber/v2"
)

type BookingRoute struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

func (route BookingRoute) RegisterRoute() {
	handler := handlers.BookingHandler{Handler: route.Handler}
	r := route.RouterGroup.Group("/api/booking")

	r.Get("/", handler.SelectAll)
	r.Get("/:id", handler.FindOne)
	r.Post("/", handler.Add)
}
