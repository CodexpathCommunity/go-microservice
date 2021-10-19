package database

// SQL contains generic SQL methods.
type SQL interface {
	// Exec queries the database and returns the affected rows
	Exec(sql string, args ...interface{}) (int64, error)
	// Query queries the database and return the rows
	Query(sql string, args ...interface{}) (Rows, error)
	// QueryRow queries the database and return a single row
	QueryRow(sql string, args ...interface{}) Row
	// BeginTx Starts a database Transaction
	BeginTx() (Transaction, error)
}

// Row is the result returned from a query
type Row interface {
	// Scan reads the values from the current row into dest values positionally
	Scan(dest ...interface{}) error
}

// Rows is the result set returned from a query
type Rows interface {
	// Scan reads the values from the current row into dest values positionally
	Scan(dest ...interface{}) error
	// Next prepares the next row for reading. It returns true if there is another row and false if no more rows are available. It automatically closes rows when all rows are read.
	Next() bool
	// Close closes the rows, making the connection ready for use again. It is safe to call Close after rows is already closed.
	Close()
	// Err returns any error that occurred while reading.
	Err() error
}

// Transaction represents an SQL database transaction
type Transaction interface {
	// Commit commits the database transaction
	Commit() error
	// Rollback rollbacks the database transaction
	Rollback() error
	// Exec queries the database and returns the affected rows
	Exec(sql string, args ...interface{}) (int64, error)
	// Query queries the database and return the rows
	Query(sql string, args ...interface{}) (Rows, error)
	// QueryRow queries the database and return a single row
	QueryRow(sql string, args ...interface{}) Row
}
