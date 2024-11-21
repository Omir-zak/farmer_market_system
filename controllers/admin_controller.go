package controllers

import (
	"database/sql"
	"farmer_market/models"
)

func GetPendingFarmers(db *sql.DB) ([]models.User, error) {
	query := `SELECT id, name, email, city FROM users WHERE role = 'farmer' AND is_approved = false`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var farmers []models.User
	for rows.Next() {
		var farmer models.User
		err := rows.Scan(&farmer.ID, &farmer.Name, &farmer.Email, &farmer.City)
		if err != nil {
			return nil, err
		}
		farmers = append(farmers, farmer)
	}
	return farmers, nil
}

func ApproveFarmer(db *sql.DB, id int) error {
	query := `UPDATE users SET is_approved = true WHERE id = $1 AND role = 'farmer'`
	_, err := db.Exec(query, id)
	return err
}

func RejectFarmer(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1 AND role = 'farmer'`
	_, err := db.Exec(query, id)
	return err
}
