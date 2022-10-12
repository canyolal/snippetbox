package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// Define a new UserModel type which wraps a database connection pool.
type UserModel struct {
	DB *sql.DB
}

// We'll use the Insert method to add a new record to the "users" table.
func (m *UserModel) Insert(name, email, password string) (int, error) {

	// stmt := `INSERT INTO users (name,email,hashed_password,created)
	// VALUES(?, ?, ?, UTC_TIMESTAMP)`

	// hashed_pass :=

	// result, err := m.DB.Exec(stmt, name, email)
	// if err != nil {
	// 	return 0, err
	// }
	return 0, nil
}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {

	return 0, nil
}

func (m *UserModel) Exists(id int) (bool, error) {

	return false, nil
}
