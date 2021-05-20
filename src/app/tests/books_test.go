package test

import (
	"bytes"
	"encoding/json"
	"globantapp/dbapi"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	_ "globantapp/routers"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/mysql"
)

func init() {
	connstr := "admin:pass)@tcp(db:3306)/globant_books?charset=utf8mb4&parseTime=True&loc=Local"
	dbapi.CreateApiLocal(mysql.Open(connstr))

	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBookGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/books/get/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	/*
		Convey("Subject: Test Station Endpoint\n", t, func() {
			Convey("Status Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})
	*/
}

func TestBooksGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/books/get", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestBooksGetByFilterPost(t *testing.T) {
	var filter dbapi.BooksDBGetFilter
	filter.MinPrice = new(float32)
	*filter.MinPrice = 7.0
	filter.MaxPrice = new(float32)
	*filter.MaxPrice = 12.0

	newBookJson, _ := json.Marshal(filter)
	r, _ := http.NewRequest("POST", "/books/get", bytes.NewReader(newBookJson))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())
}

func TestBooksGetByFilterGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/books/get?genre=2", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())
}
func TestBookCreate(t *testing.T) {
	var newBook dbapi.BookDB
	newBook.Name = "beeTest1"
	newBook.Price = 16.0
	newBook.Genre = 3
	newBook.Amount = 3

	newBookJson, _ := json.Marshal(newBook)
	r, _ := http.NewRequest("POST", "/books/create", bytes.NewReader(newBookJson))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())
}

func TestBookDelete(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/books/delete/21", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())
}
