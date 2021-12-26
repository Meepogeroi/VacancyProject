package user

import (
	"net/http"

	"github.com/Meepogeroi/VacancyProject/db/models"
	"github.com/go-pg/pg/v9"
	"github.com/labstack/echo/v4"
)

func CreateUser(conn *pg.DB) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		user := models.User{}
		if err := ctx.Bind(&user); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				struct{ Error string }{"binding error"})
		}

		if err := user.CreateUser(conn); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				struct{ Error string }{err.Error()})
		}
		return ctx.JSON(http.StatusCreated, struct {
			Responce string
		}{Responce: "user created"})
	}
}
