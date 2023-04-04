package models

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type User interface {
	AddCourse(db *sqlx.DB) (int, error)
}

type CourseData struct {
	Id     int64  `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Price  int    `json:"price" db:"price"`
	Banner string `json:"banner" db:"banner"`
}

func (c *CourseData) AddCourse(db *sqlx.DB) (int, error) {
	stmt, err := db.Prepare(`
	INSERT INTO
		courses
		(title, price, banner)
	VALUES
		($1, $2, $3)`)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if _, err := stmt.Exec(c.Title, c.Price, c.Banner); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (c *CourseData) GetByIDCourse(db *sqlx.DB) (int, error) {
	if err := db.Get(c, `
	SELECT 
		* 
	FROM 
		courses 
	WHERE 
		id=$1`, c.Id); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (c *CourseData) EditCourse(db *sqlx.DB) (int, error) {
	stmt, err := db.Prepare(`
	UPDATE
		courses
	SET
		title=$1,
		price=$2,
		banner=$3
	WHERE
		id=$4`)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if _, err := stmt.Exec(c.Title, c.Price, c.Banner, c.Id); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func GetAllCourses(db *sqlx.DB) ([]CourseData, int, error) {
	var res []CourseData

	if err := db.Select(&res, `
	SELECT 
		* 
	FROM 
		courses`); err != nil {
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}

func DeleteCourse(db *sqlx.DB, id int64) (int, error) {
	if _, err := db.Exec(`
	DELETE FROM
		courses
	WHERE
		id=$1`, id); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
