package controller

import (
	"fmt"
	"foreign_currency/helper"
	"foreign_currency/model"
	"foreign_currency/repo"
	"regexp"

	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

func AddExchange() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// Add new exchange based on from and to currency
		exchange := new(model.Exchange)
		if err = c.Bind(exchange); err != nil {
			logrus.Debug(err)
			return
		}
		// validation
		// make sure all data field (from, to) is not empty
		// make sure all data field (from, to) is valid currency based on currency list in helper
		var validationMessage []string
		if exchange.From == "" {
			validationMessage = append(validationMessage, "field 'from' is required")
		}
		if exchange.To == "" {
			validationMessage = append(validationMessage, "field 'to' is required")
		}
		if exchange.From != "" && helper.ValidateCurrency(exchange.From) != true {
			validationMessage = append(validationMessage, "'from' currency is invalid")
		}
		if exchange.To != "" && helper.ValidateCurrency(exchange.To) != true {
			validationMessage = append(validationMessage, "'to' currency is invalid")
		}
		// if some of validation failed return the message
		if len(validationMessage) > 0 {
			return c.JSON(422, validationMessage)
		}

		tx := c.Get("Tx").(*dbr.Tx)
		msg, err := repo.SaveExchange(exchange, tx)
		if err != nil {
			logrus.Debug(err)
			return
		}
		return c.JSON(200, msg)
	}
}

func GetExchanges() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// Get all exchange list
		tx := c.Get("Tx").(*dbr.Tx)
		fmt.Println("ok")
		exchanges, err := repo.GetExchanges(tx)
		return c.JSON(200, exchanges)
	}
}

func AddExchangeRate() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// Add new exchange rate based on from, to currency and date
		exchangeRate := new(model.ExhangeRate)
		if err = c.Bind(exchangeRate); err != nil {
			logrus.Debug(err)
			return
		}
		// validation
		// make sure all data field is not empty
		// make sure currency data field (from, to) is valid currency based on currency list in helper
		// make sure date field use yyyy-mm-dd format
		var validationMessage []string
		if exchangeRate.Date == "" {
			validationMessage = append(validationMessage, "field 'from' is required")
		}
		validDateFromat := dateRegex()
		if validDateFromat.MatchString(exchangeRate.Date) != true {
			validationMessage = append(validationMessage, "date format is not valid, please use yyyy-mm-dd")
		}
		if exchangeRate.From == "" {
			validationMessage = append(validationMessage, "field 'to' is required")
		}
		if exchangeRate.To == "" {
			validationMessage = append(validationMessage, "field 'to' is required")
		}
		if exchangeRate.From != "" && helper.ValidateCurrency(exchangeRate.From) != true {
			validationMessage = append(validationMessage, "'from' currency is invalid")
		}
		if exchangeRate.To != "" && helper.ValidateCurrency(exchangeRate.To) != true {
			validationMessage = append(validationMessage, "'to' currency is invalid")
		}
		if exchangeRate.Rate == 0 {
			validationMessage = append(validationMessage, "field 'to' is required")
		}
		// if some of validation failed return the message
		if len(validationMessage) > 0 {
			return c.JSON(422, validationMessage)
		}

		tx := c.Get("Tx").(*dbr.Tx)
		msg, err := repo.SaveExchangeRate(exchangeRate, tx)
		if err != nil {
			logrus.Debug(err)
			return
		}
		return c.JSON(200, msg)
	}
}

func GetExchangeRate() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// Get all exchange rate on specific date given

		// validation
		// make sure date field use yyyy-mm-dd format
		validDateFromat := dateRegex()
		date := c.Param("date")
		var validationMessage []string
		if validDateFromat.MatchString(date) != true {
			validationMessage = append(validationMessage, "date format is not valid, please use yyyy-mm-dd")
		}
		// if some of validation failed return the message
		if len(validationMessage) > 0 {
			return c.JSON(422, validationMessage)
		}
		tx := c.Get("Tx").(*dbr.Tx)
		exchangeRates, err := repo.GetExchangeRate(date, tx)
		return c.JSON(200, exchangeRates)
	}
}

func GetDetailExchange() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// Get Exchange Detail for the last 7 data point
		id := c.Param("id")
		tx := c.Get("Tx").(*dbr.Tx)
		detailExchange, err := repo.GetDetailExchange(id, tx)
		if err != nil {
			logrus.Debug(err)
			return
		}
		return c.JSON(200, detailExchange)
	}
}

func DeleteExchange() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// Delete all record of exchange_rate that associated with exchange id
		// Delete an exchange based on the exchange id given
		id := c.Param("id")
		tx := c.Get("Tx").(*dbr.Tx)
		msg, err := repo.DeleteExchange(id, tx)
		if err != nil {
			logrus.Debug(err)
			return
		}
		return c.JSON(200, msg)
	}
}

func dateRegex() *regexp.Regexp {
	// make sure user input for date is yyyy-mm-dd
	validDateFromat := regexp.MustCompile("(\\d\\d\\d\\d)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])")
	return validDateFromat
}
