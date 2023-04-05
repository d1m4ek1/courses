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
	Id          int64         `json:"id" db:"id"`
	CartItemsID pq.Int64Array `json:"cartItemsId" db:"cart_items_id"`
	UserID      string        `json:"userId" db:"user_id"`
}

type OrderItems struct {
	ID    int64         `json:"_id"`
	Items []ProductItem `json:"items"`
}

func (d *DataForOrder) AddToOrder(db *sqlx.DB) (int, error) {
	cart := ProductData{UserID: d.UserID}

	cartItems, status, err := cart.GetAllCartItems(db)
	if err != nil {
		return status, err
	}

	for _, item := range cartItems {
		d.CartItemsID = append(d.CartItemsID, item.Id)
	}

	stmt, err := db.Preparex(`
		INSERT INTO
			orders
			(cart_items_id, user_id)
		VALUES
			($1, $2)`)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if _, err := stmt.Exec(pq.Array(d.CartItemsID), d.UserID); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func generateRequestForGetUserOrder(ids []int64, userID string) string {
	var query string = `
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
		ca.course_id=c.id
		AND
		ca.%s
		AND
		(%s)`

	var queryWhere []string

	for _, id := range ids {
		queryWhere = append(queryWhere, fmt.Sprintf(`ca.id=%d`, id))
	}

	return fmt.Sprintf(query, fmt.Sprintf("user_id='%s'", userID), strings.Join(queryWhere, "OR"))
}

func (d *DataForOrder) GetUserOrders(db *sqlx.DB) ([]OrderItems, int, error) {
	var orderItems []DataForOrder
	var productInOrderItems []OrderItems

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
		var f []ProductItem
		if err := db.Select(&f, generateRequestForGetUserOrder(order.CartItemsID, d.UserID)); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		productInOrderItems = append(productInOrderItems, OrderItems{ID: order.Id, Items: f})
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
