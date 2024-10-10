package user

import (
	"github.com/VENI-VIDIVICI/plus/pkg/database"
)

func IsEmailExit(emial string) bool {
	sql := "SELECT COUNT(*) FROM users WHERE email = ?"
	var count int
	database.DB.QueryRow(sql, emial).Scan(&count)
	return count > 0
}

func IsPhoneExit(phone string) bool {
	sql := "SELECT COUNT(*) FROM users WHERE phone = ?"
	var count int
	database.DB.QueryRow(sql, phone).Scan(&count)
	return count > 0
}

func (u *User) Create() error {
	query := "INSERT INTO users (name, phone, password) VALUES (?, ?, ?)"
	_, err := database.DB.Exec(query, u.Name, u.Phone, u.Password)
	if err != nil {
		return err
	}
	return nil
}
