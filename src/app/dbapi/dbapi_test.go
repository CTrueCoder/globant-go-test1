package dbapi

import (
	"testing"

	"gorm.io/driver/mysql"
)

func getApi() (*DBApi, error) {
	connstr := "admin:pass)@tcp(127.0.0.1:3306)/globant_books?charset=utf8mb4&parseTime=True&loc=Local"
	return CreateApi(mysql.Open(connstr))
}

func TestCreateDelete(t *testing.T) {
	api, err := getApi()
	if err != nil {
		t.Fatal("error get api: ", err)
	}

	var book BookDB
	book.Name = "test5"
	book.Price = 3
	book.Genre = 2
	book.Amount = 1
	_, err = api.CreateBookDB(&book)
	if err != nil {
		t.Fatal("error create book : ", err)
	}

	deleted, err := api.DeleteBookDB(&book)
	if err != nil {
		t.Fatal("error delete book: ", err)
	}

	if !deleted {
		t.Fatal("already deleted")
	}

}

func TestGetBook(t *testing.T) {
	api, err := getApi()
	if err != nil {
		t.Fatal("error get api: ", err)
	}

	bookdb, err := api.GetBookDBById(1)
	if err != nil {
		t.Fatal("error get book")
	}
	if bookdb != nil {
		t.Fatal("ok")
	} else {
		t.Fatal("not found")
	}
}

func TestGetBooks(t *testing.T) {
	api, err := getApi()
	if err != nil {
		t.Fatal("error get api: ", err)
	}

	var filter BooksDBGetFilter
	name := "1"
	var minprice float32 = 5.0
	filter.Name = &name
	filter.MinPrice = &minprice

	books, err := api.GetBooksDB(&filter)
	if err != nil {
		t.Fatal("error get books: ", err)
	}

	if len(books) == 0 {
		t.Fatal("error books count")
	}
}

func TestUpdateBook(t *testing.T) {
	api, err := getApi()
	if err != nil {
		t.Fatal("error get api: ", err)
	}

	var update BookDBUpdate
	name := string("update1")
	update.Id = 1
	update.Name = &name
	_, err = api.UpdateBookDBUpdate(&update)
	if err != nil {
		t.Fatal("error update book: ", err)
	}
}

func BenchmarkGetBooks(b *testing.B) {
	api, err := getApi()
	if err != nil {
		b.Fatal("error get api: ", err)
	}

	var filter BooksDBGetFilter
	minprice := float32(5)
	filter.Name = new(string)
	*filter.Name = ""
	filter.MinPrice = &minprice
	filter.MaxPrice = new(float32)
	*filter.MaxPrice = 7

	b.ResetTimer()
	b.StartTimer()
	books, err := api.GetBooksDB(&filter)
	if err != nil {
		b.Fatal("error get books: ", err)
	}
	b.StopTimer()

	if len(books) == 4 {
		b.Fatal("error books count")
	}
}
