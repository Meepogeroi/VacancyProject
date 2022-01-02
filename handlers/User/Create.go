package user

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/Meepogeroi/VacancyProject/db/models"
	"github.com/go-pg/pg/v9"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func CreateUser(conn *pg.DB) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		user := models.User{}
		if err := ctx.Bind(&user); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				struct{ Error string }{"binding error"})
		}

		var errors error
		user.UUID = uuid.Must(uuid.NewV4(), errors).String()

		shaEncoder := sha256.New()
		shaEncoder.Write([]byte(user.Password))

		user.Password = base64.URLEncoding.EncodeToString(shaEncoder.Sum(nil))
		if err := user.CreateUser(conn); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				struct{ Error string }{err.Error()})
		}
		return ctx.JSON(http.StatusCreated, struct {
			Responce string
		}{Responce: "user created"})
	}
}
