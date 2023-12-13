package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
	"vending-machine/model"
)

var (
	rows *sql.Rows
)

// GetItemList : get data item list from db
func GetItemList(db *sql.DB, params map[string]interface{}, orderBy string) (retData []*model.Item, err error) {
	var (
		values []interface{}
		order  string
	)

	query := "select id, name, price, status, created_at, updated_at, deleted_at from items" // set initial query
	// add query filter if parameter params sent from controller
	if len(params) > 0 {
		query += " where "
		for i, v := range params {
			if i == "total_payment" {
				query += "price <= ? and "
				values = append(values, v)
				continue
			}

			query += i + " = ? and "
			values = append(values, v)
		}

		query = strings.TrimSuffix(query, " and ")
	}
	// add ordering for the query
	if orderBy != "" {
		if string(orderBy[0]) == "-" {
			order = "desc"
		} else {
			order = "asc"
		}

		orderBy = strings.TrimPrefix(orderBy, "-")

		query += " order by " + orderBy + " " + order
	}
	// execute query
	rows, err = db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	// loop through query result for return data variable
	for rows.Next() {
		var item = new(model.Item)

		err = rows.Scan(&item.ID, &item.Name, &item.Price, &item.Status, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)
		if err != nil {
			return nil, err
		}

		retData = append(retData, item)
	}

	return retData, nil
}

// AddItem : function to insert new items into db
func AddItem(db *sql.DB, values []interface{}) (err error) {
	query := "insert into items (name, price, status) values (?, ?, ?)"
	_, err = db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

// ModifyItem : function to changes data of spesific item in db
func ModifyItem(db *sql.DB, id int64, mapValues map[string]interface{}) (rowsAffected int64, err error) {
	var (
		values []interface{}
		rows   sql.Result
	)

	query := "update items set "
	for i, v := range mapValues {
		query += i + " = ?, "
		values = append(values, v)
	}
	query += "updated_at = ?"
	values = append(values, time.Now())

	query = strings.TrimSuffix(query, ", ")

	query += " where id = ?"
	values = append(values, id)
	fmt.Println(query, values)
	rows, err = db.Exec(query, values...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = rows.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// RemoveItem : function to delete data from db. deleted data will not be fully erased but only it's status is set into value 2
func RemoveItem(db *sql.DB, id int64) (rowsAffected int64, err error) {
	var rows sql.Result

	query := "update items set status = ?, deleted_at = ? where id = ?"
	rows, err = db.Exec(query, 2, time.Now(), id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = rows.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
