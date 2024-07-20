package handlers

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo, db *sql.DB) {
	e.GET("/", getAllTodoHandler(db))
}

