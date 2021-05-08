package service

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"example.com/golang-gin-demo/model"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var DbEngine *xorm.Engine

func init() {
	// read config file
	err := godotenv.Load()
	// set db info
	driverName := os.Getenv("DB_DRIVER")
	DsName := os.Getenv("DB_URI")

	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.Book))
	fmt.Println("init data base ok")
}
