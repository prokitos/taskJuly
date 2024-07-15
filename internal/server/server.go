package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ServerStart(port string) *fiber.App {

	app := fiber.New()
	handlers(app)
	log.Fatal(app.Listen(port))

	return app
}

func handlers(instance *fiber.App) {

	instance.Post("/accounts", routeCreateAccount)
	instance.Post("/accounts/:id/deposit", routeDepositAccount)
	instance.Post("/accounts/:id/withdraw", routeWithdrawAccount)
	instance.Get("/accounts/:id/balance", routeBalanceCheckAccount)

}
