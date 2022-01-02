package user

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"os"
	"time"

	"github.com/Meepogeroi/VacancyProject/config"
	"github.com/Meepogeroi/VacancyProject/db/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v9"
	"github.com/labstack/echo/v4"
)

func AuthUser(conn *pg.DB) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		user := models.User{}
		if err := ctx.Bind(&user); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				struct{ Error string }{"binding error"})
		}

		shaEncoder := sha256.New()
		shaEncoder.Write([]byte(user.Password))
		user.Password = base64.URLEncoding.EncodeToString(shaEncoder.Sum(nil))

		usr, err := user.GetUser(conn)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				struct{ Error string }{err.Error()})
		}

		uuid := user.UUID

		claims := &config.JWTClaims{
			Mail:           user.Mail,
			UUID:           user.UUID,
			StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 72).Unix()},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(os.Getenv("JWT")))

		if err != nil {
			return echo.ErrUnauthorized
		}

		return ctx.JSON(http.StatusOK, echo.Map{
			"token":     t,
			"uuid":      uuid,
			"fio":       usr.FIO,
			"mail":      user.Mail,
			"phone_num": user.PhoneNum,
		})
	}
}
