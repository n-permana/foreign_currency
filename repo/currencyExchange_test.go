package repo

import (
	"fmt"
	"foreign_currency/model"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"github.com/stretchr/testify/require"
)

func TestSaveValidExchange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"from", "to"}).CloseError(fmt.Errorf("close error"))
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(rows)
	mock.ExpectExec("INSERT INTO `exchange`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectClose()

	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}

	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		panic(err)
	}

	mockData := new(model.Exchange)
	mockData.From = "IDR"
	mockData.To = "USD"

	res, err := SaveExchange(mockData, tx)
	if err != nil {
		panic(err)
	} else {
		tx.Commit()
		db.Close()
		conn.Close()
	}
	expectedResult := map[string]string{
		"succes":  "true",
		"message": "the exhange has successfully saved",
	}
	if res["succes"] != expectedResult["succes"] || res["message"] != expectedResult["message"] {
		t.Errorf(`SaveExchange was incorrect, got: status %v with message %v, want: status %v with message %v.`,
			res["succes"], res["message"], expectedResult["succes"], expectedResult["message"])
	}
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveExistingExchange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mock.ExpectBegin()
	emptyrows := sqlmock.NewRows([]string{"from", "to"}).CloseError(fmt.Errorf("close error"))
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(emptyrows)
	mock.ExpectExec("INSERT INTO `exchange`").WillReturnResult(sqlmock.NewResult(1, 1))
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(rows)
	mock.ExpectCommit()
	mock.ExpectClose()

	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}

	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		panic(err)
	}

	mockData := new(model.Exchange)
	mockData.From = "IDR"
	mockData.To = "USD"

	// save new exchange
	_, err = SaveExchange(mockData, tx)
	if err != nil {
		panic(err)
	}

	// save existing exchange
	res, err := SaveExchange(mockData, tx)
	if err != nil {
		panic(err)
	} else {
		tx.Commit()
		db.Close()
		conn.Close()
	}
	expectedResult := map[string]string{
		"succes":  "false",
		"message": "the exchange is already exist",
	}
	if res["succes"] != expectedResult["succes"] || res["message"] != expectedResult["message"] {
		t.Errorf(`SaveExchange was incorrect, got: status %v with message %v, want: status %v with message %v.`,
			res["succes"], res["message"], expectedResult["succes"], expectedResult["message"])
	}
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetExchangesReturnValues(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mock.ExpectBegin()
	emptyrows := sqlmock.NewRows([]string{"from", "to"}).CloseError(fmt.Errorf("close error"))
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(emptyrows)
	mock.ExpectExec("INSERT INTO `exchange`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(emptyrows)
	mock.ExpectExec("INSERT INTO `exchange`").WillReturnResult(sqlmock.NewResult(1, 1))
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2)
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(rows)
	mock.ExpectCommit()
	mock.ExpectClose()

	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}

	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		panic(err)
	}

	mockData1 := new(model.Exchange)
	mockData1.From = "IDR"
	mockData1.To = "USD"
	// save new exchange
	_, err = SaveExchange(mockData1, tx)
	if err != nil {
		panic(err)
	}
	mockData2 := new(model.Exchange)
	mockData2.From = "IDR"
	mockData2.To = "JPY"
	// save new exchange
	_, err = SaveExchange(mockData2, tx)
	if err != nil {
		panic(err)
	}

	// get all exchange
	res, err := GetExchanges(tx)
	if err != nil {
		panic(err)
	} else {
		tx.Commit()
		db.Close()
		conn.Close()
	}

	expectedResult := 2
	if len(res) != expectedResult {
		t.Errorf(`GetExchanges was incorrect, got: %v data, want: %v data.`, len(res), expectedResult)
	}

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetExchangesReturnEmpty(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mock.ExpectBegin()
	emptyrows := sqlmock.NewRows([]string{"from", "to"}).CloseError(fmt.Errorf("close error"))
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(emptyrows)
	mock.ExpectCommit()
	mock.ExpectClose()

	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}

	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		panic(err)
	}

	// get all exchange
	res, err := GetExchanges(tx)
	if err != nil {
		panic(err)
	} else {
		tx.Commit()
		db.Close()
		conn.Close()
	}

	expectedResult := 0
	if len(res) != expectedResult {
		t.Errorf(`GetExchanges was incorrect, got: %v data, want: %v data.`, len(res), expectedResult)
	}

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetValidExchangesRateWithNotExistExchange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mock.ExpectBegin()
	emptyrows := sqlmock.NewRows([]string{"from", "to"}).CloseError(fmt.Errorf("close error"))
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(emptyrows)
	mock.ExpectExec("INSERT INTO `exchange`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO `exchange_rate`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectClose()

	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}

	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		panic(err)
	}
	mockData := new(model.ExhangeRate)
	mockData.Date = "2018-07-02"
	mockData.From = "IDR"
	mockData.To = "USD"
	mockData.Rate = 1.299

	res, err := SaveExchangeRate(mockData, tx)
	if err != nil {
		panic(err)
	} else {
		tx.Commit()
		db.Close()
		conn.Close()
	}

	expectedResult := map[string]string{
		"succes":  "true",
		"message": "the exhange rate has successfully saved",
	}
	if res["succes"] != expectedResult["succes"] || res["message"] != expectedResult["message"] {
		t.Errorf(`SaveExchange was incorrect, got: status %v with message %v, want: status %v with message %v.`,
			res["succes"], res["message"], expectedResult["succes"], expectedResult["message"])
	}

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetValidExchangesRateWithExistingExchange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mock.ExpectBegin()
	emptyrows := sqlmock.NewRows([]string{"from", "to"}).CloseError(fmt.Errorf("close error"))
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(emptyrows)
	mock.ExpectExec("INSERT INTO `exchange`").WillReturnResult(sqlmock.NewResult(1, 1))
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(rows)
	mock.ExpectExec("INSERT INTO `exchange_rate`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectClose()

	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}

	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		panic(err)
	}

	mockData := new(model.Exchange)
	mockData.From = "IDR"
	mockData.To = "USD"

	_, err = SaveExchange(mockData, tx)
	if err != nil {
		panic(err)
	}

	mockData2 := new(model.ExhangeRate)
	mockData2.Date = "2018-07-02"
	mockData2.From = "IDR"
	mockData2.To = "USD"
	mockData2.Rate = 1.299

	res, err := SaveExchangeRate(mockData2, tx)
	if err != nil {
		panic(err)
	} else {
		tx.Commit()
		db.Close()
		conn.Close()
	}

	expectedResult := map[string]string{
		"succes":  "true",
		"message": "the exhange rate has successfully saved",
	}
	if res["succes"] != expectedResult["succes"] || res["message"] != expectedResult["message"] {
		t.Errorf(`SaveExchange was incorrect, got: status %v with message %v, want: status %v with message %v.`,
			res["succes"], res["message"], expectedResult["succes"], expectedResult["message"])
	}

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetDetailExchanges(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mock.ExpectBegin()
	emptyrows := sqlmock.NewRows([]string{"from", "to"}).CloseError(fmt.Errorf("close error"))
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(emptyrows)
	mock.ExpectExec("INSERT INTO `exchange`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO `exchange_rate`").WillReturnResult(sqlmock.NewResult(1, 1))
	rows := sqlmock.NewRows([]string{"from", "to"}).AddRow("IDR", "USD")
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(rows)
	rows2 := sqlmock.NewRows([]string{"from", "to", "avergae", "variance"}).AddRow("IDR", "USD", 1.299, 1.299)
	mock.ExpectQuery("SELECT (.+) FROM exchange_rate").WillReturnRows(rows2)
	mock.ExpectCommit()
	mock.ExpectClose()

	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}

	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		panic(err)
	}

	mockData := new(model.ExhangeRate)
	mockData.Date = "2018-07-02"
	mockData.From = "IDR"
	mockData.To = "USD"
	mockData.Rate = 1.299

	_, err = SaveExchangeRate(mockData, tx)
	if err != nil {
		panic(err)
	}

	_, err = GetDetailExchange("1", tx)
	if err != nil {
		panic(err)
	} else {
		tx.Commit()
		db.Close()
		conn.Close()
	}

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetExchangeRate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mock.ExpectBegin()
	emptyrows := sqlmock.NewRows([]string{"from", "to"}).CloseError(fmt.Errorf("close error"))
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(emptyrows)
	mock.ExpectExec("INSERT INTO `exchange`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO `exchange_rate`").WillReturnResult(sqlmock.NewResult(1, 1))
	rows2 := sqlmock.NewRows([]string{"id", "from", "to", "rate", "variance"}).AddRow("1", "IDR", "USD", 1.299, 1.299)
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(rows2)
	mock.ExpectCommit()
	mock.ExpectClose()

	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}

	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		panic(err)
	}

	mockData := new(model.ExhangeRate)
	mockData.Date = "2018-07-03"
	mockData.From = "IDR"
	mockData.To = "USD"
	mockData.Rate = 1.299

	_, err = SaveExchangeRate(mockData, tx)
	if err != nil {
		panic(err)
	}

	_, err = GetExchangeRate("2018-07-02", tx)
	if err != nil {
		panic(err)
	} else {
		tx.Commit()
		db.Close()
		conn.Close()
	}

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteExchange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mock.ExpectBegin()
	emptyrows := sqlmock.NewRows([]string{"from", "to"}).CloseError(fmt.Errorf("close error"))
	mock.ExpectQuery("SELECT (.+) FROM exchange").WillReturnRows(emptyrows)
	mock.ExpectExec("INSERT INTO `exchange`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO `exchange_rate`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("DELETE FROM `exchange_rate`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("DELETE FROM `exchange`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectClose()

	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}

	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		panic(err)
	}

	mockData := new(model.ExhangeRate)
	mockData.Date = "2018-07-02"
	mockData.From = "IDR"
	mockData.To = "USD"
	mockData.Rate = 1.299

	_, err = SaveExchangeRate(mockData, tx)
	if err != nil {
		panic(err)
	}

	res, err := DeleteExchange("1", tx)
	if err != nil {
		panic(err)
	} else {
		tx.Commit()
		db.Close()
		conn.Close()
	}

	expectedResult := map[string]string{
		"succes":  "true",
		"message": "the exhange has successfully deleted",
	}
	if res["succes"] != expectedResult["succes"] || res["message"] != expectedResult["message"] {
		t.Errorf(`DeleteExchange was incorrect, got: status %v with message %v, want: status %v with message %v.`,
			res["succes"], res["message"], expectedResult["succes"], expectedResult["message"])
	}

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteInvalidExchange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `exchange_rate`").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("DELETE FROM `exchange`").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()
	mock.ExpectClose()

	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}

	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		panic(err)
	}

	res, err := DeleteExchange("1", tx)
	if err != nil {
		panic(err)
	} else {
		tx.Commit()
		db.Close()
		conn.Close()
	}

	expectedResult := map[string]string{
		"succes":  "false",
		"message": "exchange id is not exist",
	}
	if res["succes"] != expectedResult["succes"] || res["message"] != expectedResult["message"] {
		t.Errorf(`DeleteExchange was incorrect, got: status %v with message %v, want: status %v with message %v.`,
			res["succes"], res["message"], expectedResult["succes"], expectedResult["message"])
	}

	require.NoError(t, mock.ExpectationsWereMet())
}
