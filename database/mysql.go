package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"regexp"
	"strings"
)

type MySQLError struct {
	Name  string
	Error error
}

func (m *MySQL) MatchError(err error) MySQLError {
	matched, _ := regexp.MatchString("duplicate", strings.ToLower(err.Error()))
	if matched {
		return MySQLError{
			Name:  "duplicate",
			Error: err,
		}
	}
	return MySQLError{}
}

// MySQL is our primary struct
type MySQL struct {
	LastQuery          string
	Connection         *sqlx.DB
	Config             mysql.Config
	MaxOpenConnections int
	Connected          bool
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
	m.MaxOpenConnections = 50
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
	m.Connection, err = sqlx.Open("mysql", m.Config.FormatDSN())
	if err != nil {
		return err
	}
	m.Connection.SetMaxOpenConns(m.MaxOpenConnections)
	m.Connected = true
	return nil
}

// SelectOne single result, stored within arg:into
//      country := Country{}
//      country, err := db.Select(&country, "SELECT * FROM countries WHERE name='Germany' ORDER BY name ASC")
func (m *MySQL) SelectOne(into any, query string, args ...any) (any, error) {
	if !m.Connected {
		return into, errors.New("not connected")
	}
	err := m.Connection.Get(&into, query, args...)
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	return into, err
}

// Select for more than one result is expected
//      countries := []Countries{}
//      countries, err := db.Select(&countries, "SELECT * FROM countries ORDER BY name ASC")
func (m *MySQL) Select(into []any, query string, args ...any) ([]any, error) {
	if !m.Connected {
		return into, errors.New("not connected")
	}
	err := m.Connection.Select(into, query, args...)
	return into, err
}

// InsertOne one struct into a named query using sqlx standards
// 		person := Person{ FirstName: "Ardie" }
// 		lastInsertId, err = db.NamedExec(`INSERT INTO persons (first_name) VALUES (:first_name)`, person)
func (m *MySQL) InsertOne(namedQuery string, insertStruct any) (int64, error) {
	if !m.Connected {
		return -1, errors.New("not connected")
	}
	result, err := m.Connection.NamedExec(namedQuery, insertStruct)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

// Insert many structs into a named query using sqlx standards
// 		persons := []Person{
// 			{FirstName: "Ardie"},
//			{FirstName: "Sonny"},
// 		}
// 		err = db.NamedExec(`INSERT INTO persons (first_name) VALUES (:first_name)`, persons)
func (m *MySQL) Insert(namedQuery string, insertStruct []any) error {
	if !m.Connected {
		return errors.New("not connected")
	}
	_, err := m.Connection.NamedExec(namedQuery, insertStruct)
	return err
}

// UpdateOne performs an update of one record
// 		persons := Person{ FirstName: "Ardie" }
// 		err = db.NamedExec(`UPDATE persons SET first_name=:first_name`, persons)
func (m *MySQL) UpdateOne(namedQuery string, updateStruct any) (int64, error) {
	if !m.Connected {
		return -1, errors.New("not connected")
	}
	result, err := m.Connection.NamedExec(namedQuery, updateStruct)
	if err != nil {
		return -1, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return affected, nil
}

// Update performs an update of many records
// 		persons := []Person{
// 			{FirstName: "Ardie"},
//			{FirstName: "Sonny"},
// 		}
// 		err = db.NamedExec(`UPDATE persons SET first_name=:first_name`, persons)
func (m *MySQL) Update(namedQuery string, updateStructs []any) error {
	if !m.Connected {
		return errors.New("not connected")
	}
	_, err := m.Connection.NamedExec(namedQuery, updateStructs)
	return err
}

// DeleteOne performs a deletion
// 		persons := Person{Id: 0}
// 		err = db.NamedExec(`DELETE FROM persons WHERE id=:id`, persons)
func (m *MySQL) DeleteOne(namedQuery string, deleteStruct any) error {
	_, err := m.Connection.NamedExec(namedQuery, deleteStruct)
	return err
}

// Delete performs a deletion
// 		persons := []Person{
// 			{Id: 0},
//			{Id: 1},
// 		}
// 		err = db.NamedExec(`DELETE FROM persons WHERE id=:id`, persons)
func (m *MySQL) Delete(namedQuery string, deleteStructs []any) (int64, error) {
	results, err := m.Connection.NamedExec(namedQuery, deleteStructs)
	if err != nil {
		return -1, err
	}

	affected, err := results.RowsAffected()
	if err != nil {
		return -1, err
	}
	return affected, nil
}
