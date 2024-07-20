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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal server error",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "successfully get all todos",
			"data":    todos,
		})
	}
}
// handler to get todo by id
func getTodoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		t, err := service.GetTodoById(db, id)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Todo not found",
			})
		} else if err != nil {
			c.Logger().Error(err) 
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "successfully get todo by id",
			"data":    t,
		})
	}
}
