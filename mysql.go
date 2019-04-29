package mantis

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"os"
)

// Define the query struct
type Query struct {
	Query      string
	Connection *sql.DB
	DSN        mysql.Config
}

// Connect to the database
func (q *Query) Connect() bool {
	q.DSN = mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  os.Getenv("DB_NET"),
		Addr:                 os.Getenv("DB_HOST"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
		AllowOldPasswords:    true,
		RejectReadOnly:       true,
	}

	var err error
	q.Connection, err = sql.Open("mysql", q.DSN.FormatDSN())

	HandleError("Error creating MySQL Connection", err)
	return true
}

// Select for when one result is expected
func (q *Query) SelectOne(args ...interface{}) *sql.Row {
	row := q.Connection.QueryRow(q.Query, args...)
	return row
}

// Select for when more than one result is expected
func (q *Query) Select(args ...interface{}) *sql.Rows {
	rows, err := q.Connection.Query(q.Query, args...)
	HandleError("Error connecting to database", err)
	return rows
}

// Insert a query
func (q *Query) Insert(args ...interface{}) int64 {
	stmt, err := q.Connection.Prepare(q.Query)
	HandleError("Error preparing insertion query", err)

	res, err := stmt.Exec(args...)
	HandleError("Error executing insertion query", err)

	id, err := res.LastInsertId()
	HandleError("Error fetching last ID in insertion query", err)
	return id
}

// Perform an update
func (q *Query) Update(args ...interface{}) int64 {
	stmt, err := q.Connection.Prepare(q.Query)
	HandleError("Error preparing update query", err)

	res, err := stmt.Exec(args...)
	HandleError("Error executing update query", err)

	affect, err := res.RowsAffected()
	HandleError("Error fetching affected rows in update query", err)
	return affect
}

// Perform a deletion
func (q *Query) Delete(args ...interface{}) int64 {
	stmt, err := q.Connection.Prepare(q.Query)
	HandleError("Error preparing deletion query", err)

	res, err := stmt.Exec(args...)
	HandleError("Error executing deletion query", err)

	affect, err := res.RowsAffected()
	HandleError("Error fetching affected rows in deletion query", err)
	return affect
}
