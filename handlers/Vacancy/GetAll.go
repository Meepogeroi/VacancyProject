package vacancy

import (
	"net/http"

	"github.com/Meepogeroi/VacancyProject/db/models"
	"github.com/go-pg/pg/v9"
	"github.com/labstack/echo/v4"
)

func GetAll(conn *pg.DB) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		vac := &models.Vacancy{}
		vacs, err := vac.GetAll(conn)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				struct{ Error string }{err.Error()})
		}
		return ctx.JSON(http.StatusOK, vacs)
	}
}