package controllers

import (
	"database/sql"
	"errors"
	"farmer_market/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var JwtKey = []byte("your_secret_key")

// Структура токена
type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func LoginAdmin(db *sql.DB, email, password string) (string, error) {
	// Ищем пользователя по email
	query := `SELECT id, password, role FROM users WHERE email = $1 AND role = 'admin'`
	var user models.User
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Password, &user.Role)
	if err != nil {
		return "", err
	}

	// Проверяем пароль
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// Генерируем JWT токен
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

func RegisterBuyer(db *sql.DB, buyer models.User) error {
	buyer.Role = "buyer"
	buyer.IsApproved = true // Покупателю подтверждение не требуется

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(buyer.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	buyer.Password = string(hashedPassword)

	query := `INSERT INTO users (name, email, password, role, is_approved) VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(query, buyer.Name, buyer.Email, buyer.Password, buyer.Role, buyer.IsApproved)
	return err
}

func RegisterFarmer(db *sql.DB, farmer models.User) error {
	// Проверяем, существует ли email
	query := `SELECT COUNT(*) FROM users WHERE email = $1`
	var count int
	err := db.QueryRow(query, farmer.Email).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("email already exists")
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(farmer.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	farmer.Password = string(hashedPassword)

	// Добавляем фермера в базу
	query = `INSERT INTO users (name, email, password, role, city, is_approved) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = db.Exec(query, farmer.Name, farmer.Email, farmer.Password, "farmer", farmer.City, false)
	if err != nil {
		return err
	}

	return nil
}

// controllers/user_controller.go
/*package controllers

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
*/
