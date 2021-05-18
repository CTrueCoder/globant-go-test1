package main

import (
	"globantapp/dbapi"
	_ "globantapp/routers"
	"os"

	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
)

func main() {
	connstr := "admin:pass)@tcp(db:3306)/globant_books" //?charset=utf8mb4&parseTime=True&loc=Local
	if dbapi.CreateApiLocal(mysql.Open(connstr)) != nil {
		os.Exit(1)
	}

	beego.Run()
}
