package dbapi

import (
	"errors"

	"gorm.io/gorm"
)

type DBApi struct {
	db *gorm.DB
}

var Sess *DBApi

type GenreDB struct {
	//gorm.Model

	Id   int `gorm:"primarykey"`
	Name string
}

type BookDB struct {
	//gorm.Model

	Id     int `gorm:"primarykey"`
	Name   string
	Price  float32
	Genre  int
	Amount int
}

type BookDBUpdate struct {
	Id     int `gorm:"primarykey"`
	Name   *string
	Price  *float32
	Genre  *int
	Amount *int
}

type BooksDBGetFilter struct {
	Name     *string
	MinPrice *float32
	MaxPrice *float32
	Genre    *int
}

func CreateDB(dialector gorm.Dialector) (db *gorm.DB, err error) {
	return gorm.Open(dialector, &gorm.Config{})
}

func CreateApi(dialector gorm.Dialector) (api *DBApi, err error) {
	db, lerr := CreateDB(dialector)
	if lerr != nil {
		return nil, lerr
	}

	return &DBApi{db}, nil
}

func CreateApiLocal(dialector gorm.Dialector) (err error) {
	Sess, err = CreateApi(dialector)
	return
}

func (api *DBApi) getBooksTable() *gorm.DB {
	return api.db.Table("books")
}

// create row using struct. field "id" not using
func (api *DBApi) CreateBookDB(book *BookDB) (int, error) {
	db := api.getBooksTable().Create(book)
	err := db.Error
	if err != nil {
		return 0, err
	}

	return book.Id, nil
}

// Create row using args
func (api *DBApi) CreateBook(name string, price float32, genre int, amount int) (int, error) {
	bookdb := BookDB{Name: name, Price: price, Genre: genre, Amount: amount}
	return api.CreateBookDB(&bookdb)
}

func (api *DBApi) GetBookDBById(id int) (bookdb *BookDB, err error) {
	var lbookdb BookDB
	lerr := api.getBooksTable().First(&lbookdb, id).Error
	if lerr != nil {
		if errors.Is(lerr, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, lerr
	}
	return &lbookdb, nil
}

func (api *DBApi) DeleteBookDB(bookdb *BookDB) (deleted bool, err error) {
	db := api.getBooksTable().Delete(bookdb)
	return db.RowsAffected == 1, db.Error
}

func (api *DBApi) DeleteBookById(id int) (deleted bool, err error) {
	db := api.getBooksTable().Where("id = ?", id).Delete(id)
	return db.RowsAffected == 1, db.Error
}

func (api *DBApi) UpdateBookDB(bookdb *BookDB) (updated bool, err error) {
	//return api.getBooksTable().Update(bookdb).Error
	db := api.getBooksTable().Updates(bookdb)
	return db.RowsAffected == 1, db.Error
}

// Update table book row one by using id in bookdb.id
func (api *DBApi) UpdateBookDBUpdate(bookdb *BookDBUpdate) (updated bool, err error) {
	//return api.getBooksTable().Update(bookdb).Error
	db := api.getBooksTable()
	db = db.Model(bookdb)
	if bookdb.Name != nil {
		db = db.Select("name")
	}
	if bookdb.Price != nil {
		db = db.Select("price")
	}
	if bookdb.Genre != nil {
		db = db.Select("genre")
	}
	if bookdb.Amount != nil {
		db = db.Where("amount")
	}

	db = db.Updates(bookdb)
	return db.RowsAffected == 1, db.Error
	//return db.Where("id = ?", bookdb.Id).Updates(bookdb).Error
}

func (api *DBApi) GetBooksDB(filter *BooksDBGetFilter) (booksdb []BookDB, err error) {
	db := api.getBooksTable()

	if filter.Name != nil {
		db = db.Where("name LIKE ?", "%"+*filter.Name+"%")
	}
	if filter.MinPrice != nil {
		db = db.Where("price >= ?", filter.MinPrice)
	}
	if filter.MaxPrice != nil {
		db = db.Where("price <= ?", filter.MaxPrice)
	}
	if filter.Genre != nil {
		db = db.Where("genre = ?", filter.Genre)
	}
	err = db.Where("amount > 0").Find(&booksdb).Error

	return
}
