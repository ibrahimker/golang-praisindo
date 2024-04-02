package repository

import (
	"database/sql"
	"log"

	"github.com/ibrahimker/golang-praisindo/user-service-grpc/entity"
)

type UserRepository struct {
	db *sql.DB
}

type IUserRepository interface {
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	GetAll() ([]entity.User, error)
	GetByEmail(email string) (entity.User, error)
	Delete(email string) error
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetAll() ([]entity.User, error) {
	queryString := "SELECT username, email, password, age FROM users LIMIT 10"
	rows, err := u.db.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(
			&user.Username, &user.Email, &user.Password, &user.Age,
		); err != nil {
			log.Println(err)
			continue
		}
		res = append(res, user)
	}

	return res, nil
}

func (u *UserRepository) GetByEmail(email string) (entity.User, error) {
	var user entity.User
	queryString := "SELECT username, email, password, age FROM users WHERE email = $1 LIMIT 1"
	if err := u.db.QueryRow(queryString, email).Scan(
		&user.Username, &user.Email, &user.Password, &user.Age); err != nil {
		log.Println(err)
		return entity.User{}, err
	}

	return user, nil
}

// Create implements IUserRepository.
func (u *UserRepository) Create(user entity.User) (entity.User, error) {
	queryString := "INSERT INTO users (username, email, password, age) " +
		"VALUES ($1,$2,$3,$4) "
	_, err := u.db.Exec(queryString, user.Username, user.Email, user.Password, user.Age)
	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

// Delete implements IUserRepository.
func (u *UserRepository) Delete(email string) error {
	queryString := "DELETE FROM users WHERE email = $1"
	_, err := u.db.Exec(queryString, email)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Update implements IUserRepository.
func (u *UserRepository) Update(user entity.User) (entity.User, error) {
	queryString := "UPDATE users " +
		"SET username=$2, password=$3, age=$4" +
		"WHERE email=$1"
	_, err := u.db.Exec(queryString, user.Email, user.Username, user.Password, user.Age)
	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}
