package user

import "github.com/VENI-VIDIVICI/plus/pkg/database"

func (u User) IsEmailExit(emial string) bool {
	sql := "SELECT COUNT(*) FROM user WHERE emial = ?"
	var count int
	database.DB.QueryRow(sql, emial).Scan(&count)
	return count > 0
}

func (u User) IsPhoneExit(phone string) bool {
	sql := "SELECT COUNT(*) FROM user WHERE phone = ?"
	var count int
	database.DB.QueryRow(sql, phone).Scan(&count)
	return count > 0
}
