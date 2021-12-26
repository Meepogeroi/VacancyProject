package models

import "github.com/go-pg/pg/v9"

type User struct {
	ID       int    `pg:"id,pk"`
	FIO      string `pg:"fio" json:"fio"`
	Mail     string `pg:"mail, unique" json:"mail"`
	PhoneNum string `pg:"phone_num" json:"phone_num"`
	Password string `pg:"password" json:"password"`
	CV       []byte `pg:"cv" json:"cv"` //Резюме, я не знаю как файл загрузить в бд, поэтому оставил массив байт
	WorkExp  string `pg:"work_exp" json:"work_exp"`
}

func (usr *User) CreateUser(conn *pg.DB) error {
	err := conn.Insert(usr)
	if err != nil {
		return err
	}
	return nil
}

func (usr *User) GetUser(conn *pg.DB) (*User, error) {
	user := &User{}
	err := conn.Model(user).
		Where("mail = ?0", usr.Mail).
		Where("password = ?0", usr.Password).
		Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (usr *User) GetAll(conn *pg.DB) (*[]User, error) {
	users := &[]User{}
	err := conn.Model(users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}
