package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type MySQL struct {
	LastQuery  string
	Connection *sql.DB
	Config     mysql.Config
}

// Connect to the database
func (q *MySQL) Connect() error {
	var err error
	q.Connection, err = sql.Open("mysql", q.Config.FormatDSN())
	if err != nil {
		return err
	}
	q.Connection.SetMaxOpenConns(10)
	return nil
}

// SelectOne selects for a single result
func (q *MySQL) SelectOne(query string, args ...interface{}) *sql.Row {
	return q.Connection.QueryRow(query, args...)
}

// Select for more than one result is expected
func (q *MySQL) Select(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := q.Connection.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Insert a query
func (q *MySQL) Insert(query string, args ...interface{}) (int64, error) {
	stmt, err := q.Connection.Prepare(query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

// Update performs an update
func (q *MySQL) Update(query string, args ...interface{}) (int64, error) {
	stmt, err := q.Connection.Prepare(query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	return affected, nil
}

// Delete performs a deletion
func (q *MySQL) Delete(query string, args ...interface{}) (int64, error) {
	stmt, err := q.Connection.Prepare(query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	return affected, nil
}
