package models

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Order interface {
	AddToOrder(db *sqlx.DB) (int, error)

	GetUserOrders(db *sqlx.DB) ([]ProductItem, int, error)

	DeleteOrder(db *sqlx.DB) (int, error)
}

type DataForOrder struct {
	Id            int64         `json:"id" db:"id"`
	CourseItemsID pq.Int64Array `json:"courseItemsId" db:"course_items_id"`
	UserID        string        `json:"userId" db:"user_id"`
	Count         pq.Int64Array `json:"count" db:"count"`
	AddDateOrder  string        `json:"addDateOrder" db:"add_date_order"`
	Token         string        `json:"-" db:"token"`
}

type OrderItems struct {
	ID           int64         `json:"_id"`
	FirstName    string        `json:"firstName" db:"first_name"`
	SecondName   string        `json:"secondName" db:"second_name"`
	ThirdName    string        `json:"thirdName" db:"third_name"`
	Email        string        `json:"email" db:"email"`
	Count        pq.Int64Array `json:"count" db:"count"`
	AddDateOrder string        `json:"addDateOrder" db:"add_date_order"`
	Items        []OrderItem   `json:"items"`
}

type OrderItem struct {
	Id         int    `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Price      int    `json:"price" db:"price"`
	CourseID   int    `json:"courseId" db:"course_id"`
	FirstName  string `json:"firstName" db:"first_name"`
	SecondName string `json:"secondName" db:"second_name"`
	ThirdName  string `json:"thirdName" db:"third_name"`
}

func (d *DataForOrder) AddToOrder(db *sqlx.DB) (int, error) {
	var cart ProductData

	if err := db.Get(&d.UserID, `SELECT user_id FROM users WHERE token != '' AND token=$1`, d.Token); err != nil {
		return http.StatusInternalServerError, err
	}

	cart.UserID = d.UserID

	cartItems, status, err := cart.GetAllCartItems(db)
	if err != nil {
		return status, err
	}

	for _, item := range cartItems {
		d.CourseItemsID = append(d.CourseItemsID, item.CourseID)
		d.Count = append(d.Count, int64(item.Count))
	}

	stmt, err := db.Preparex(`
		INSERT INTO
			orders
			(course_items_id, user_id, add_date_order, count)
		VALUES
			($1, $2, $3, $4)`)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if _, err := stmt.Exec(pq.Array(d.CourseItemsID), d.UserID, d.AddDateOrder, pq.Array(d.Count)); err != nil {
		return http.StatusInternalServerError, err
	}

	if status, err := d.deleteAllCartAfterAddToOrder(db); err != nil {
		return status, err
	}

	return http.StatusOK, nil
}

func generateRequestForGetUserOrder(ids []int64, userID string) string {
	var query string = `
	SELECT 
		c.id AS course_id,
		c.title,
		c.price
	FROM
		courses c
	WHERE
		c.%s
		AND
		(%s)`

	var queryWhere []string

	for _, id := range ids {
		queryWhere = append(queryWhere, fmt.Sprintf(`c.id=%d`, id))
	}

	return fmt.Sprintf(query, fmt.Sprintf("user_id='%s'", userID), strings.Join(queryWhere, " OR "))
}

func (d *DataForOrder) GetUserOrders(db *sqlx.DB) ([]OrderItems, int, error) {
	var orderItems []DataForOrder
	var productInOrderItems []OrderItems
	var orderData OrderItems

	if err := db.Get(&d.UserID, `SELECT user_id FROM users WHERE token != '' AND token=$1`, d.Token); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := db.Get(&orderData, `
	SELECT 
		first_name,
		second_name,
		third_name,
		email
	FROM 
		users
	WHERE 
		user_id=$1`, d.UserID); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := db.Select(&orderItems, `
	SELECT
		*
	FROM
		orders
	WHERE
		user_id=$1`, d.UserID); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	for _, order := range orderItems {
		var f []OrderItem
		if err := db.Select(&f, generateRequestForGetUserOrder(order.CourseItemsID, d.UserID)); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		productInOrderItems = append(productInOrderItems, OrderItems{
			ID:           order.Id,
			FirstName:    orderData.FirstName,
			SecondName:   orderData.SecondName,
			ThirdName:    orderData.ThirdName,
			Email:        orderData.Email,
			AddDateOrder: order.AddDateOrder,
			Count:        order.Count,
			Items:        f})
	}

	return productInOrderItems, http.StatusOK, nil
}

func (d *DataForOrder) DeleteOrder(db *sqlx.DB) (int, error) {
	if _, err := db.Exec(`
		DELETE FROM
			orders
		WHERE
			id=$1`, d.Id); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (d *DataForOrder) deleteAllCartAfterAddToOrder(db *sqlx.DB) (int, error) {
	if _, err := db.Exec(`DELETE FROM cart WHERE user_id=$1`, d.UserID); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
