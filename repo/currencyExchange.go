package repo

import (
	"fmt"
	"foreign_currency/model"

	"github.com/gocraft/dbr"
)

func SaveExchange(exchange *model.Exchange, tx *dbr.Tx) (map[string]string, error) {
	existingExchange, err := findOrCreateExchange(exchange, tx)
	if err != nil {
		return response("something wrong", 0, err)
	}
	if existingExchange["new"] == 0 {
		return response("the exchange is already exist", 0, err)
	}
	return response("the exhange has successfully saved", 1, err)
}

func GetExchanges(tx *dbr.Tx) ([]model.Exchange, error) {
	var exchanges []model.Exchange
	_, err := tx.Select("*").
		From("exchange").
		Load(&exchanges)
	if err != nil {
		return nil, err
	}
	return exchanges, nil
}

func GetDetailExchange(id string, tx *dbr.Tx) (interface{}, error) {
	type queryResult struct {
		Date string  `json:"date"`
		Rate float64 `json:"rate"`
	}
	type result struct {
		From     string  `json:"from"`
		To       string  `json:"to"`
		Average  float64 `json:"average"`
		Variance float64 `json:"variance"`
		Rates    []queryResult
	}
	exchange := new(model.Exchange)
	_, err := tx.Select("*").
		From("exchange").
		Where("id = ?", id).
		Load(&exchange)

	var queryResults []queryResult
	query := fmt.Sprintf(`
    SELECT 
    exchange_rate.date,
    exchange_rate.rate
    FROM exchange_rate
    WHERE exchange_rate.exchange_id = %s order by exchange_rate.date desc limit 7 
    `, id)
	_, err = tx.SelectBySql(query).Load(&queryResults)
	var finalResult result
	totalRate := float64(0)
	maxRate := float64(0)
	minRate := float64(0)
	for i, data := range queryResults {
		totalRate += float64(data.Rate)
		if float64(data.Rate) > maxRate {
			maxRate = float64(data.Rate)
		}
		if float64(data.Rate) < minRate || i == 0 {
			minRate = float64(data.Rate)
		}
	}
	finalResult.From = exchange.From
	finalResult.To = exchange.To
	finalResult.Average = float64(totalRate / 7)
	finalResult.Variance = float64(maxRate - minRate)
	finalResult.Rates = queryResults
	return finalResult, err
}

func SaveExchangeRate(exchangeRate *model.ExhangeRate, tx *dbr.Tx) (map[string]string, error) {
	var exchange model.Exchange
	exchange.From = exchangeRate.From
	exchange.To = exchangeRate.To
	existingExchange, err := findOrCreateExchange(&exchange, tx)
	if err != nil {
		return response("something wrong", 0, err)
	}
	_, err = tx.InsertInto("exchange_rate").
		Columns("date", "exchange_id", "rate").
		Values(exchangeRate.Date, existingExchange["exchangeId"], exchangeRate.Rate).
		Exec()
	if err != nil {
		return response("something wrong", 0, err)
	}
	return response("the exhange rate has successfully saved", 1, err)
}

func GetExchangeRate(date string, tx *dbr.Tx) (interface{}, error) {
	type queryResult struct {
		Id      string         `json:"id"`
		From    string         `json:"from"`
		To      string         `json:"to"`
		Rate    dbr.NullString `json:"rate"`
		Average dbr.NullString `json:"average"`
	}
	var queryResults []queryResult
	query := fmt.Sprintf(`
  SELECT
  exchange.id,
	exchange.from,
	exchange.to,
	exchange_rate.rate,
	(SELECT avg(rate) FROM exchange_rate as ier where ier.date <= '%s' and ier.date >= DATE_ADD('%s', INTERVAL -7 DAY) and ier.exchange_id = exchange.id) as average
  FROM exchange
  left join exchange_rate on exchange.id = exchange_rate.exchange_id and exchange_rate.date = '%s'
  group by exchange.id,exchange.from,exchange.to,exchange_rate.rate order by exchange_rate.rate desc
  `, date, date, date)
	_, err := tx.SelectBySql(query).Load(&queryResults)
	for i, data := range queryResults {
		if data.Rate.Valid != true {
			queryResults[i].Rate = dbr.NewNullString("insufficient data")
			queryResults[i].Average = dbr.NewNullString("insufficient data")
		}
	}
	return queryResults, err
}

func DeleteExchange(id string, tx *dbr.Tx) (map[string]string, error) {
	res, err := tx.DeleteFrom("exchange_rate").Where("exchange_id = ?", id).Exec()
	if err != nil {
		return response("something wrong", 0, err)
	}
	res, err = tx.DeleteFrom("exchange").Where("id = ?", id).Exec()
	if err != nil {
		return response("something wrong", 0, err)
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		return response("something wrong", 0, err)
	}
	if rowAffected < 1 {
		return response("exchange id is not exist", 0, err)
	}
	return response("the exhange has successfully deleted", 1, err)
}

// private function

func findOrCreateExchange(exchange *model.Exchange, tx *dbr.Tx) (map[string]interface{}, error) {
	// check if the exchange is already exist and return it's ID
	var existingExchange model.Exchange
	result := make(map[string]interface{})
	_, err := tx.Select("*").
		From("exchange").
		Where("`from` = ?", exchange.From).
		Where("`to` = ?", exchange.To).
		Load(&existingExchange)
	if err != nil {
		return result, err
	}
	if existingExchange.Id != 0 {
		// if the exchange exist return it's ID
		result["new"] = 0
		result["exchangeId"] = existingExchange.Id
	} else {
		// if not exist create new exchange and return the ID
		res, err := tx.InsertInto("exchange").
			Columns("from", "to").
			Record(exchange).
			Exec()
		if err != nil {
			return result, err
		}
		result["new"] = 1
		result["exchangeId"], err = res.LastInsertId()
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

func response(msg string, succes int, err error) (map[string]string, error) {
	responseFormat := make(map[string]string)
	responseFormat["succes"] = "true"
	if succes != 1 {
		responseFormat["succes"] = "false"
	}
	responseFormat["message"] = msg
	return responseFormat, err
}
