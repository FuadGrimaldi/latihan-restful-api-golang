package handlers

import (
	"database/sql"
	"fmt"
	"golang-crud/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getAllTodoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		todos, err := service.GetAllTodos(db)
		fmt.Println(todos)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, todos)
	}
}

