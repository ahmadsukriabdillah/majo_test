package repository

import (
	"database/sql"
	"errors"
	"log"
	"majo_test/models"
)

type UserRepository interface {
	FindByUsername(username string) (models.User, error)
	GetAccessOutlet(user_id uint) ([]string, error)
	GetAccessMerchant(user_id uint) ([]string, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	row := u.db.QueryRow("select id, user_name, name, password from users where user_name = ?", username)
	err := row.Scan(&user.Id, &user.Username, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("error: user not found")
		} else {
			return models.User{}, errors.New("error: user not found")
		}
	}
	return user, nil
}

func (u *userRepository) GetAccessMerchant(user_id uint) ([]string, error) {
	rows, err := u.db.Query("select id from merchants where user_id = ?", user_id)
	if err != nil {
		return nil, errors.New("error: system error")
	}
	defer rows.Close()
	var datas []string
	for rows.Next() {
		var item string
		err := rows.Scan(&item)
		if err != nil {
			log.Fatal(err)
		}
		datas = append(datas, item)
	}
	return datas, nil
}

func (u *userRepository) GetAccessOutlet(user_id uint) ([]string, error) {
	rows, err := u.db.Query("select o.id from outlets o left join merchants m on o.merchant_id = m.id where m.user_id = ?", user_id)
	if err != nil {
		return nil, errors.New("error: system error")
	}
	defer rows.Close()
	var datas []string
	for rows.Next() {
		var item string
		err := rows.Scan(&item)
		if err != nil {
			log.Fatal(err)
		}
		datas = append(datas, item)
	}
	return datas, nil
}
