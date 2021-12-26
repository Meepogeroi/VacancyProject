package vacancy

import (
	"net/http"

	"github.com/Meepogeroi/VacancyProject/db/models"
	"github.com/go-pg/pg/v9"
	"github.com/labstack/echo/v4"
)

func GetVac(conn *pg.DB) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		vac := &models.Vacancy{}
		if err := ctx.Bind(&vac); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				struct{ Error string }{err.Error()})
		}
		vac_get, err := vac.GetVac(conn)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				struct{ Error string }{err.Error()})
		}
		return ctx.JSON(http.StatusOK, vac_get)
	}
}