package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	"github.com/sudachi0114/gin-rest-backend-trial/model"
)

var DbEngine *xorm.Engine
var err error

func init() {
	driverName := "mysql"
	DsName := "root:root@tcp(127.0.0.1:3306)/sample_db?charset=utf8mb4&parseTime=True&loc=Local"
	err = errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal("<DB 接続失敗>", err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.Note))
	fmt.Println("init Database OK")
}
