package repository

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"testing"
	"time"
	"vending-machine/model"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	db        *sql.DB
	mock      sqlmock.Sqlmock
	err       error
	xpctQuery string
	xpctVal   []driver.Value
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

// TestGetItemList : unit test for GetItemList function returned success
func TestGetItemList(t *testing.T) {
	var retData []*model.Item

	// new mock db connection
	db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db connection : %v", err)
	}
	defer db.Close()

	// set up expected query and values
	xpctQuery = "select id, name, price, status, created_at, updated_at, deleted_at from items"
	xpctVal = []driver.Value{1}

	// exec query
	mock.ExpectQuery(xpctQuery).WithArgs(xpctVal...).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "status", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "Item1", 11000, 1, time.Now(), nil, nil).
		AddRow(2, "Item2", 12000, 1, time.Now(), nil, nil))

	// call function with mock db
	retData, err = GetItemList(db, map[string]interface{}{"status": 1}, "id")
	if err != nil {
		t.Fatalf("Failed to get item list : %v", err)
	}

	// check returned data
	if len(retData) != 2 {
		t.Fatalf("Expect 2 items, got %d instead", len(retData))
	}

	// check mock expectations
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed mock. Expectations were not met : %v", err)
	}
}

// TestGetItemListError : unit test for GetItemList function returned error
func TestGetItemListError(t *testing.T) {
	var xpctError string

	// new mock db connection
	db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db connection : %v", err)
	}
	defer db.Close()

	// set up expected query and values
	xpctQuery = "select id, name, price, status, created_at, updated_at, deleted_at from items"
	xpctVal = []driver.Value{1}

	xpctError = "mocked error"

	// exec query
	mock.ExpectQuery(xpctQuery).WithArgs(xpctVal...).WillReturnError(errors.New(xpctError))

	// call function with mock db
	_, err = GetItemList(db, map[string]interface{}{"status": 1}, "id")
	if err == nil || err.Error() != xpctError {
		t.Fatalf("Failed to get item list or error doesn't returned %v, got %v instead", xpctError, err)
	}

	// check mock expectations
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed mock. Expectations were not met : %v", err)
	}
}

// TestAddItem : unit test for AddItem function returned success
func TestAddItem(t *testing.T) {
	var val []interface{}

	// new mock db connection
	db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db connection : %v", err)
	}
	defer db.Close()

	// set up expected query and values
	xpctQuery = "insert into items \\(name, price, status\\) values \\(\\?, \\?, \\?\\)"
	xpctVal = []driver.Value{"TestItem", 11000, 1}

	val = []interface{}{"TestItem", 11000, 1}

	// exec query
	mock.ExpectExec(xpctQuery).WithArgs(xpctVal...).WillReturnResult(sqlmock.NewResult(1, 1))

	// call function with mock db
	err = AddItem(db, val)
	if err != nil {
		t.Fatalf("Failed to insert item : %v", err)
	}

	// check mock expectations
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed mock. Expectations were not met : %v", err)
	}
}

// TestAddItemError : unit test for AddItem function returned error
func TestAddItemError(t *testing.T) {
	var (
		val       []interface{}
		xpctError string
	)

	// new mock db connection
	db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db connection : %v", err)
	}
	defer db.Close()

	// set up expected query and values
	xpctQuery = "insert into items \\(name, price, status\\) values \\(\\?, \\?, \\?\\)"
	xpctVal = []driver.Value{"TestItem", 11000, 1}

	val = []interface{}{"TestItem", 11000, 1}

	xpctError = "mocked error"

	// exec query
	mock.ExpectExec(xpctQuery).WithArgs(xpctVal...).WillReturnError(errors.New(xpctError))

	// call function with mock db
	err = AddItem(db, val)
	if err == nil || err.Error() != xpctError {
		t.Fatalf("Failed to insert item or error doesn't returned %v, got %v instead", xpctError, err)
	}

	// check mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Failed mock. Expectations were not met : %v", err)
	}
}

// TestModifyItem : unit test for ModifyItem function returned success
func TestModifyItem(t *testing.T) {
	// new mock db connection
	db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db connection : %v", err)
	}
	defer db.Close()

	// set up expected query and values
	xpctQuery := "update items set name = \\?, price = \\?, updated_at = \\? where id = \\?"
	xpctVal := []driver.Value{"NewName", 15000, AnyTime{}, 1}

	// setup parameters
	id := int64(1)
	mapValues := map[string]interface{}{
		"name":  "NewName",
		"price": 15000,
	}

	// exec query
	mock.ExpectExec(xpctQuery).WithArgs(xpctVal...).WillReturnResult(sqlmock.NewResult(0, 1))

	// call function with mock db
	rowsAffected, err := ModifyItem(db, id, mapValues)
	if err != nil {
		t.Fatalf("Failed to update item : %v", err)
	}

	// check rows affected
	if rowsAffected != 1 {
		t.Fatalf("Expect 1 row affected, got %d instead", rowsAffected)
	}

	// check mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Failed mock. Expectations were not met : %v", err)
	}
}

// TestModifyItemError : unit test for ModifyItem function returned error
func TestModifyItemError(t *testing.T) {
	var xpctError string

	// new mock db connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db connection : %v", err)
	}
	defer db.Close()

	// set up expected query and values
	xpctQuery := "update items set name = \\?, price = \\?, updated_at = \\? where id = \\?"
	xpctVal := []driver.Value{"NewName", 15000, AnyTime{}, 1}

	// setup parameters
	id := int64(1)
	mapValues := map[string]interface{}{
		"name":  "NewName",
		"price": 15000,
	}

	xpctError = "mocked error"

	// exec query
	mock.ExpectExec(xpctQuery).WithArgs(xpctVal...).WillReturnError(errors.New(xpctError))

	// call function with mock db
	rowsAffected, err := ModifyItem(db, id, mapValues)
	if err == nil || err.Error() != xpctError {
		t.Fatalf("Failed to update item or error doesn't returned %v, got %v instead", xpctError, err)
	}

	// check rows affected
	if rowsAffected != 0 {
		t.Fatalf("Expect 0 row affected, got %d instead", rowsAffected)
	}

	// check mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Failed mock. Expectations were not met : %v", err)
	}
}

// TestRemoveItem : unit test for RemoveItem function returned success
func TestRemoveItem(t *testing.T) {
	// new mock db connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db connection : %v", err)
	}
	defer db.Close()

	// set up expected query and values
	xpctQuery := "update items set status = \\?, deleted_at = \\? where id = \\?"
	xpctVal := []driver.Value{2, AnyTime{}, 1}

	// setup parameters
	id := int64(1)

	// exec query
	mock.ExpectExec(xpctQuery).WithArgs(xpctVal...).WillReturnResult(sqlmock.NewResult(0, 1))

	// call function with mock db
	rowsAffected, err := RemoveItem(db, id)
	if err != nil {
		t.Fatalf("Failed to delete item : %v", err)
	}

	// check rows affected
	if rowsAffected != 1 {
		t.Fatalf("Expect 1 row affected, got %d instead", rowsAffected)
	}

	// check mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Failed mock. Expectations were not met : %v", err)
	}
}

// TestRemoveItemError : unit test for RemoveItem function returned error
func TestRemoveItemError(t *testing.T) {
	var xpctError string

	// new mock db connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db connection : %v", err)
	}
	defer db.Close()

	// set up expected query and values
	xpctQuery := "update items set status = \\?, deleted_at = \\? where id = \\?"
	xpctVal := []driver.Value{2, AnyTime{}, 1}

	// setup parameters
	id := int64(1)
	xpctError = "mocked error"

	// exec query
	mock.ExpectExec(xpctQuery).WithArgs(xpctVal...).WillReturnError(errors.New(xpctError))

	// call function with mock db
	rowsAffected, err := RemoveItem(db, id)
	if err == nil || err.Error() != xpctError {
		t.Fatalf("Failed to delete item or error doesn't returned %v, got %v instead", xpctError, err)
	}

	// check rows affected
	if rowsAffected != 0 {
		t.Fatalf("Expect 0 row affected, got %d instead", rowsAffected)
	}

	// check mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Failed mock. Expectations were not met : %v", err)
	}
}
