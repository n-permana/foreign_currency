package db

import (
	"fmt"
	"foreign_currency/config"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

func Init() *dbr.Session {

	session := getSession()
	tables := initialTable()
	for _, table := range tables {
		fmt.Println(table)
		cek, err := session.Exec(table)
		fmt.Println(cek)
		fmt.Println(err)
	}
	return session
}

func getSession() *dbr.Session {

	db, err := dbr.Open("mysql",
		config.USER+":"+config.PASSWORD+"@tcp("+config.HOST+":"+config.PORT+")/"+config.DB,
		nil)
	fmt.Println(db)
	fmt.Println(err)
	if err != nil {
		logrus.Error(err)
	} else {
		session := db.NewSession(nil)
		return session
	}
	return nil
}
