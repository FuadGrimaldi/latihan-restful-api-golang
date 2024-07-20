package handlers

import (
	"database/sql"
	"fmt"
	"golang-crud/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// handler to get all the data todos
func getAllTodoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		todos, err := service.GetAllTodos(db)
		fmt.Println(todos)
		if err != nil {
			return JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return JSONResponse(c, http.StatusOK, "successfully get all todos", todos)
	}
}
// handler to get todo by id
func getTodoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		t, err := service.GetTodoById(db, id)
		if err == sql.ErrNoRows {
			return JSONResponse(c, http.StatusNotFound, err.Error(), nil)
		} else if err != nil {
			c.Logger().Error(err) 
			return JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return JSONResponse(c, http.StatusOK, "successfully get todo by id", t)
	}
}

// handler to create todo
func createTodoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := new(service.Todo)
		if err := c.Bind(t); err != nil {
			c.Logger().Error(err)
			return JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		}
		if err :=t.Create(db); err != nil {
			c.Logger().Error(err)
			return JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return JSONResponse(c, http.StatusCreated, "Create successfully",t)
	}
}


