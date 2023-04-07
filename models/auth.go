package models

import (
	generatetoken "courses/pkg/generateToken"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type UserAuth interface {
	Login(db *sqlx.DB) (bool, int, error)

	CreateUser(db *sqlx.DB) (int, error)
}

type UserAuthData struct {
	Id         int64  `json:"_id" db:"id"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"password" db:"password"`
	FirstName  string `json:"firstName" db:"first_name"`
	SecondName string `json:"secondName" db:"second_name"`
	ThirdName  string `json:"thirdName" db:"third_name"`
	Token      string `json:"token" db:"token"`
	UserID     string `json:"_userId" db:"user_id"`
}

func hashPassword(password, email string) string {
	pwd := sha256.New()
	pwd.Write([]byte(fmt.Sprintf(`%s:%s`, password, email)))

	hash := pwd.Sum(nil)

	return string(fmt.Sprintf("%x", hash))
}

func setHashUserID(email, f_name string) string {
	pwd := sha256.New()
	pwd.Write([]byte(fmt.Sprintf(`%s:%s`, email, f_name)))
	hash := pwd.Sum(nil)

	return string(fmt.Sprintf("%x", hash))
}

func (s *UserAuthData) Login(db *sqlx.DB) (bool, int, error) {
	var isAuth bool
	var err error

	password := hashPassword(s.Password, s.Email)

	s.Token, err = generatetoken.CreateToken(256)
	if err != nil {
		return false, http.StatusInternalServerError, err
	}

	if err := db.Get(&isAuth, `SELECT EXISTS (SELECT * FROM users WHERE password=$1 AND email=$2)`, password, s.Email); err != nil {
		return false, http.StatusInternalServerError, err
	}

	if err := db.Get(&s.UserID, `SELECT user_id FROM users WHERE password=$1 AND email=$2`, password, s.Email); err != nil {
		return false, http.StatusInternalServerError, err
	}

	if _, err := db.Exec(`
			UPDATE
				users
			SET
				token=$1
			WHERE
				user_id=$2`, s.Token, s.UserID); err != nil {
		return false, http.StatusInternalServerError, err
	}

	return isAuth, http.StatusOK, nil
}

func (s *UserAuthData) CreateUser(db *sqlx.DB) (int, error) {
	var err error
	var id int64

	hashPassword := hashPassword(s.Password, s.Email)

	s.Token, err = generatetoken.CreateToken(256)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	s.UserID = setHashUserID(s.Email, s.FirstName)

	if err := db.Get(&id, `
	INSERT INTO
		users
		(password, email, first_name, user_id, token)
	VALUES
		($1, $2, $3, $4, $5) RETURNING id`, hashPassword, s.Email, s.FirstName, s.UserID, s.Token); err != nil {
		return http.StatusInternalServerError, err
	}

	if id != 0 {
		if s.SecondName != "" {
			if _, err := db.Exec(`
			UPDATE
				users
			SET
				second_name=$1
			WHERE
				id=$2
				AND
				user_id=$3`, s.SecondName, id, s.UserID); err != nil {
				return http.StatusInternalServerError, err
			}
		}

		if s.ThirdName != "" {
			if _, err := db.Exec(`
			UPDATE
				users
			SET
				third_name=$1
			WHERE
				id=$2
				AND
				user_id=$3`, s.ThirdName, id, s.UserID); err != nil {
				return http.StatusInternalServerError, err
			}
		}

		return http.StatusOK, nil
	}

	return http.StatusInternalServerError, nil
}

func (s *UserAuthData) Logout(db *sqlx.DB) (int, error) {
	if _, err := db.Exec("UPDATE users SET token='' WHERE token=$1", s.Token); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
