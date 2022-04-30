package database

import (
	"database/sql"
	"encoding/json"
	"github.com/go-sql-driver/mysql"
)

// MySQL is our primary struct
type MySQL struct {
	LastQuery  string
	Connection *sql.DB
	Config     mysql.Config
}

func (m *MySQL) Default(user, password, address, dbName string) {
	m.Config = mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		AllowNativePasswords: true,
		Addr:                 address,
		DBName:               dbName,
	}
}

// ConfigString turns our configuration into a JSON string
func (m *MySQL) ConfigString() string {
	marshaledStruct, err := json.Marshal(m.Config)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

// String returns our MySQL struct as a JSON string
func (m *MySQL) String() string {
	marshaledStruct, err := json.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

// Connect to the database
func (m *MySQL) Connect() error {
	var err error
	m.Connection, err = sql.Open("mysql", m.Config.FormatDSN())
	if err != nil {
		return err
	}
	m.Connection.SetMaxOpenConns(10)
	return nil
}

// SelectOne selects for a single result
func (m *MySQL) SelectOne(query string, args ...any) (interface{}, error) {
	var into interface{}
	row := m.Connection.QueryRow(query, args...)
	err := row.Scan(&into)
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	return into, nil
}

// Select for more than one result is expected
func (m *MySQL) Select(query string, args ...any) (*sql.Rows, error) {
	rows, err := m.Connection.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Insert a query
func (m *MySQL) Insert(query string, args ...any) (int64, error) {
	stmt, err := m.Connection.Prepare(query)
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
func (m *MySQL) Update(query string, args ...any) (int64, error) {
	stmt, err := m.Connection.Prepare(query)
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
func (m *MySQL) Delete(query string, args ...any) (int64, error) {
	stmt, err := m.Connection.Prepare(query)
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
