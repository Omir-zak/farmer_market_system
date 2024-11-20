// controllers/user_controller.go
package controllers

import (
	"database/sql"
	"errors"
	"farmer_market/models"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateUser(db *sql.DB, user models.User) error {
	if db == nil {
		return errors.New("database connection is not initialized")
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v\n", err)
		return err
	}
	user.Password = string(hashedPassword)

	// SQL query to insert the user
	query := `INSERT INTO users (name, email, password, role, is_approved) VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(query, user.Name, user.Email, user.Password, user.Role, user.IsApproved)

	// Handle duplicate email error
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { // Unique violation error code
			return errors.New("email already exists")
		}
		log.Printf("Error creating user: %v\n", err)
		return err
	}

	return nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	query := `SELECT id, name, email, password, role, is_approved FROM users WHERE email = $1`
	var user models.User
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.IsApproved)
	if err != nil {
		log.Printf("Error fetching user by email: %v\n", err)
		return nil, err
	}
	return &user, nil
}
