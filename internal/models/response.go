package models

import (
	"errors"
	"fmt"
)

func ResponseGood() error {
	return errors.New("good")
}

func ResponseErrorAtServer() error {
	return errors.New("internal error")
}

func ResponseBadRequest() error {
	return errors.New("bad request")
}

func ResponseBalanceGood(balance float64) error {

	newBalance := fmt.Sprint(balance)
	return errors.New("balance is: " + newBalance)
}

func ResponseTooSlow() error {
	return errors.New("The waiting time has been exceeded")
}
