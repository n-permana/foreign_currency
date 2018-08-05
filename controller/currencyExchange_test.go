package controller

import (
	"net/http/httptest"
	"strings"
	"testing"

	"foreign_currency/db"
	myMw "foreign_currency/middleware"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]string{
		"success": "true",
		"message": "the exhange has successfully saved",
	}
	validExchangeJSON = `{"from":"JPY","to":"USD"}`
)

func TestValidAddExchange(t *testing.T) {
	// Setup
	e := echo.New()
	transactionMiddleware := myMw.TransactionTestHandler(db.Init())
	req := httptest.NewRequest(echo.POST, "/exchange", strings.NewReader(validExchangeJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, transactionMiddleware(AddExchange())(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}
