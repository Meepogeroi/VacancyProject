package db

import (
	"github.com/Meepogeroi/VacancyProject/db/models"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{
		(*models.User)(nil),
		(*models.Vacancy)(nil),
	} {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			return err
		}
	}
	return nil
}
