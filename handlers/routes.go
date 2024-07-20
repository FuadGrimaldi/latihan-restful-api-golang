package handlers

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

// routes API
func Route(e *echo.Echo, db *sql.DB) {
	e.GET("/", getAllTodoHandler(db))
	e.POST("/", createTodoHandler(db))
	e.GET("/:id", getTodoHandler(db))
}

