package models

import "github.com/go-pg/pg/v9"

type Vacancy struct {
	ID           int    `pg:"id,pk"`        //id
	Position     string `pg:"position" json:"position"`     //Должность
	Charge       string `pg:"charge" json:"charge"`       //Обязанность
	Requirements string `pg:"requirements" json:"requirements"` //Требования
	Conditions   string `pg:"conditions" json:"conditions"`   //Условия работы
}

func (vac *Vacancy) Create(conn *pg.DB) error {
	err := conn.Insert(vac)
	if err != nil {
		return err
	}
	return nil
}

func (vac *Vacancy) GetVac(conn *pg.DB) (*Vacancy, error) {
	vacancy := &Vacancy{}
	err := conn.Model(vacancy).
		Where("id = ?0", vac.ID).
		Select()
	if err != nil {
		return nil, err
	}
	return vacancy, nil
}

func (vac *Vacancy) GetAll(conn *pg.DB) (*[]Vacancy, error) {
	vacs := &[]Vacancy{}
	err := conn.Model(vacs).Select()
	if err != nil {
		return nil, err
	}
	return vacs, nil
}
