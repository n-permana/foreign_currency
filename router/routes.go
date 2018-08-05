package router

import (
	"foreign_currency/controller"
	"foreign_currency/db"

	myMw "foreign_currency/middleware"

	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {

	e := echo.New()
	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	e.Use(myMw.TransactionHandler(db.Init()))
	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.POST("/exchange", controller.AddExchange())
		v1.GET("/exchange", controller.GetExchanges())
		v1.GET("/exchange/:id", controller.GetDetailExchange())
		v1.POST("/exchange_rate", controller.AddExchangeRate())
		v1.GET("/exchange_rate/:date", controller.GetExchangeRate())
		v1.DELETE("/exchange/:id", controller.DeleteExchange())
	}

	return e
}
