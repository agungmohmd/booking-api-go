package bootstrap

import (
	"github.com/agungmohmd/booking-api/server/bootstrap/routes"
	"github.com/agungmohmd/booking-api/server/handlers"
)

func (boot Bootstrap) RegisterRouters() {
	handler := handlers.Handler{
		FiberApp:   boot.App,
		ContractUC: &boot.ContractUC,
	}

	apiV1 := boot.App.Group("/V1")
	bookingRoutes := routes.BookingRoute{RouterGroup: apiV1, Handler: handler}
	bookingRoutes.RegisterRoute()
}
