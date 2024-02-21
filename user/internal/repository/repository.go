package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"user/internal/models"
)

type UserRepository interface {
	Create(email, hashepassword string) error
	Check(email, password string) error
	Profile(email string) (*models.UserDTO, error)
	List() (*[]models.UserDTO, error)
}

type UserRepo struct {
	Postgres *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) Create(email, hashedpassword string) error {
	query := `INSERT INTO users (email, hashedpassword) VALUES ($1, $2)`
	result, err := u.Postgres.Exec(query, email, hashedpassword)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	// Получение количества затронутых строк
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %v", err)
	}

	return nil

}

func (u *UserRepo) Check(email, password string) error {
	var user models.User
	query := `SELECT id, email, hashedpassword FROM users WHERE email = $1`
	err := u.Postgres.Get(&user, query, email)
	if err != nil {
		log.Println("err not found user")
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashePassword), []byte(password))
	if err != nil {
		log.Printf("invalid password: %v", err)
		return err
	}
	return nil
}

func (u *UserRepo) Profile(email string) (*models.UserDTO, error) {
	var user models.UserDTO
	query := `SELECT id, email FROM users WHERE email = $1`
	err := u.Postgres.Get(&user, query, email)
	if err != nil {
		log.Println("err user not exist")
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) List() ([]models.UserDTO, error) {
	var users []models.UserDTO
	query := `SELECT id, email FROM users `
	err := u.Postgres.Select(&users, query)
	if err != nil {
		log.Println("err dont get users")
		return nil, err
	}

	return users, nil
}
