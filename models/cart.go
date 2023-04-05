package models

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Cart interface {
	AddToCart(db *sqlx.DB) (int, error)

	DeleteItemFromCart(db *sqlx.DB) (int, error)

	GetAllProducts(db *sqlx.DB) ([]ProductItem, int, error)

	GetAllCartItems(db *sqlx.DB) ([]ProductData, int, error)
}

type ProductData struct {
	Id       int64  `db:"id"`
	CourseID int    `json:"courseId" db:"course_id"`
	Count    int    `db:"count"`
	UserID   string `json:"userId" db:"user_id"`
}

type ProductItem struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Price    int    `json:"price" db:"price"`
	Count    int    `json:"count" db:"count"`
	CourseID int    `json:"courseId" db:"course_id"`
}

func (p *ProductData) AddToCart(db *sqlx.DB) (int, error) {
	var isHasCourseInCart bool
	if err := db.Get(&isHasCourseInCart, `
	SELECT EXISTS
		(
			SELECT
				*
			FROM
				cart
			WHERE
				course_id=$1
				AND
				user_id=$2
		)`, p.CourseID, p.UserID); err != nil {
		return http.StatusInternalServerError, err
	}

	if isHasCourseInCart {
		if _, err := db.Exec(`
		UPDATE 
			cart 
		SET 
			count=count+1 
		WHERE 
			course_id=$1
			AND
			user_id=$2`, p.CourseID, p.UserID); err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		if _, err := db.Exec(`
		INSERT INTO
			cart
			(course_id, user_id)
		VALUES
			($1, $2)`, p.CourseID, p.UserID); err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}

func (p *ProductData) GetAllProducts(db *sqlx.DB) ([]ProductItem, int, error) {
	var res []ProductItem
	stmt, err := db.Preparex(`
	SELECT
		ca.id as id,
		c.title as title,
		c.price as price,
		ca.count as count,
		ca.course_id as course_id
	FROM
		courses c,
		cart ca
	WHERE
		c.id=ca.course_id
		AND
		c.user_id=ca.user_id
		AND
		c.user_id=$1`)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := stmt.Select(&res, p.UserID); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}

func (p *ProductData) GetAllCartItems(db *sqlx.DB) ([]ProductData, int, error) {
	var res []ProductData
	stmt, err := db.Preparex(`
	SELECT
		*
	FROM
		cart
	WHERE
		user_id=$1`)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := stmt.Select(&res, p.UserID); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}

func (p *ProductData) DeleteItemFromCart(db *sqlx.DB) (int, error) {
	var countCoursesInCart int
	if err := db.Get(&countCoursesInCart, `
		SELECT
			count
		FROM
			cart
		WHERE
			id=$1
		`, p.Id); err != nil {
		return http.StatusInternalServerError, err
	}

	if countCoursesInCart > 1 {
		if _, err := db.Exec(`
		UPDATE 
			cart 
		SET 
			count=count-1 
		WHERE 
			id=$1`, p.Id); err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		if _, err := db.Exec(`
		DELETE FROM
			cart
		WHERE
			id=$1`, p.Id); err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}
