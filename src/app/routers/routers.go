package routers

import (
	"globantapp/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/books/get/:id:int", &controllers.BooksController{}, "get:GetById")
	beego.Router("/books/get", &controllers.BooksController{}, "get:GetByFilterGet;post:GetByFilterPost")
	beego.Router("/books/create", &controllers.BooksController{}, "post:Create")
	beego.Router("/books/update/:id:int", &controllers.BooksController{}, "put:UpdateById")
	beego.Router("/books/delete/:id:int", &controllers.BooksController{}, "delete:DeleteById")
}
