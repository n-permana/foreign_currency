package controller

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	myMw "foreign_currency/middleware"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestValidAddExchange(t *testing.T) {
	// Setup
	var (
		sampleInput = `{"from":"IDR","to":"USD"}`
	)
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
	e := echo.New()
	transactionMiddleware := myMw.TransactionTestHandler(sess)
	req := httptest.NewRequest(echo.POST, "/exchange", strings.NewReader(sampleInput))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, transactionMiddleware(AddExchange())(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestAddEmptyExchange(t *testing.T) {
	// Setup
	var (
		sampleInput = `{"from":"","to":""}`
	)
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
	e := echo.New()
	transactionMiddleware := myMw.TransactionTestHandler(sess)
	req := httptest.NewRequest(echo.POST, "/exchange", strings.NewReader(sampleInput))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, transactionMiddleware(AddExchange())(c)) {
		assert.Equal(t, 422, rec.Code)
	}
}

func TestInValidAddExchange(t *testing.T) {
	// Setup
	var (
		sampleInput = `{"from":"AAA","to":"ZZZ"}`
	)
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
	e := echo.New()
	transactionMiddleware := myMw.TransactionTestHandler(sess)
	req := httptest.NewRequest(echo.POST, "/exchange", strings.NewReader(sampleInput))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, transactionMiddleware(AddExchange())(c)) {
		assert.Equal(t, 422, rec.Code)
	}
}

func TestGetExchanges(t *testing.T) {
	// Setup
	var (
		sampleInput = `{"from":"IDR","to":"USD"}`
	)
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
	e := echo.New()
	transactionMiddleware := myMw.TransactionTestHandler(sess)
	req := httptest.NewRequest(echo.GET, "/exchange", strings.NewReader(sampleInput))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, transactionMiddleware(GetExchanges())(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}
