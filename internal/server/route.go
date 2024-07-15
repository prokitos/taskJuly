package server

import (
	"module/internal/models"
	"module/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func routeCreateAccount(c *fiber.Ctx) error {

	go services.CreateAccount()

	return models.ResponseGood()
}

func routeDepositAccount(c *fiber.Ctx) error {

	id := c.Params("id")

	var amount bodyInput
	if err := c.BodyParser(&amount); err != nil {
		return models.ResponseBadRequest()
	}

	idNew, err := strconv.Atoi(id)
	if err != nil {
		log.Debug("id couldn't convert to a number: " + id)
		return models.ResponseBadRequest()
	}

	// контексты решил не делать
	superChan := make(chan error)
	go services.DepositeToBasicAccount(superChan, int64(idNew), amount.Amount)
	return <-superChan
}
func routeWithdrawAccount(c *fiber.Ctx) error {

	id := c.Params("id")

	var amount bodyInput
	if err := c.BodyParser(&amount); err != nil {
		return models.ResponseBadRequest()
	}

	idNew, err := strconv.Atoi(id)
	if err != nil {
		log.Debug("id couldn't convert to a number: " + id)
		return models.ResponseBadRequest()
	}

	// контексты решил не делать
	superChan := make(chan error)
	go services.WithdrawByBasicAccount(superChan, int64(idNew), amount.Amount)
	return <-superChan
}
func routeBalanceCheckAccount(c *fiber.Ctx) error {

	id := c.Params("id")

	idNew, err := strconv.Atoi(id)
	if err != nil {
		log.Debug("id couldn't convert to a number: " + id)
		return models.ResponseBadRequest()
	}

	// контексты решил не делать
	superChan := make(chan error)
	go services.BalanceFromBasicAccount(superChan, int64(idNew))
	return <-superChan
}

type bodyInput struct {
	Amount float64 `json:"amount"`
}
