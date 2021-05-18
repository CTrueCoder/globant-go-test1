package controllers

import (
	"encoding/json"
	"globantapp/dbapi"
	"net/http"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type BooksController struct {
	beego.Controller
}

func (c *BooksController) GetById() {
	sid := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	} else {
		bookdb, err := dbapi.Sess.GetBookDBById(id)
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			if bookdb == nil {
				c.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
				return
			}
			c.Data["json"] = bookdb
		}
	}

	c.ServeJSON()
}

func (c *BooksController) GetByFilterGet() {
	var err error

	var filter dbapi.BooksDBGetFilter
	if fltname := c.GetString("name"); len(fltname) > 0 {
		filter.Name = &fltname
	}
	if sminprice := c.GetString("minprice"); len(sminprice) > 0 {
		minprice, _ := strconv.ParseFloat(sminprice, 32)
		filter.MinPrice = new(float32)
		*filter.MinPrice = float32(minprice)
	}
	if smaxprice := c.GetString("maxprice"); len(smaxprice) > 0 {
		maxprice, _ := strconv.ParseFloat(smaxprice, 32)
		filter.MaxPrice = new(float32)
		*filter.MaxPrice = float32(maxprice)
	}
	if sgenre := c.GetString("genre"); len(sgenre) > 0 {
		genre, _ := strconv.ParseInt(sgenre, 10, 32)
		filter.Genre = new(int)
		*filter.Genre = int(genre)
	}

	if err == nil {
		books, err := dbapi.Sess.GetBooksDB(&filter)
		if err == nil {
			c.Data["json"] = books
		}
	}
	if err != nil {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *BooksController) GetByFilterPost() {
	var filter dbapi.BooksDBGetFilter
	var err error
	if len(c.Ctx.Input.RequestBody) > 0 {
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &filter)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	if err == nil {
		books, err := dbapi.Sess.GetBooksDB(&filter)
		if err == nil {
			c.Data["json"] = books
		}
	}
	if err != nil {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *BooksController) Create() {
	var newBook dbapi.BookDB
	if len(c.Ctx.Input.RequestBody) > 0 {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &newBook)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	_, err := dbapi.Sess.CreateBookDB(&newBook)
	if err != nil {
		c.Data["json"] = err.Error()
		return
	} else {
		c.Data["json"] = newBook
	}

	c.ServeJSON()
}

func (c *BooksController) UpdateById() {
	sid := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	var update dbapi.BookDBUpdate
	if len(c.Ctx.Input.RequestBody) > 0 {
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &update)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	update.Id = id

	updated, err := dbapi.Sess.UpdateBookDBUpdate(&update)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	if !updated {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
		return
	}
}

func (c *BooksController) DeleteById() {
	sid := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	deleted, err := dbapi.Sess.DeleteBookById(id)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	if deleted {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
	} else {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
	}
}
