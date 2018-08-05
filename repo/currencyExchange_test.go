package repo

import (
	"foreign_currency/db"
	"foreign_currency/model"
	"testing"
)

func TestValidSaveExchange(t *testing.T) {
	expectedResult := map[string]string{
		"succes":  "true",
		"message": "the exhange has successfully saved",
	}
	session := db.Init()
	tx, _ := session.Begin()
	defer tx.Rollback()
	mockData := new(model.Exchange)
	mockData.From = "JPY"
	mockData.To = "USD"
	res, _ := SaveExchange(mockData, tx)
	if res["succes"] != expectedResult["succes"] && res["message"] != expectedResult["message"] {
		t.Errorf(`SaveExchange was incorrect, got: status %v with message %v, want: status %v with message %v.`,
			res["succes"], res["message"], expectedResult["succes"], expectedResult["message"])
	}

}
