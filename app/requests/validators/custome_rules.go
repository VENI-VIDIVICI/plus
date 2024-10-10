package validators

import (
	"errors"
	"fmt"
	"strings"

	"github.com/VENI-VIDIVICI/plus/pkg/database"
	"github.com/thedevsaddam/govalidator"
)

func init() {
	govalidator.AddCustomRule("not_exists", func(field, rule, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")
		tableName := rng[0]
		dbFiled := rng[1]
		var exceptId string
		if len(rng) > 2 {
			exceptId = rng[2]
		}
		// sql := "SELECT COUNT(*) FROM users WHERE phone = ?"
		requestValue := value.(string)
		query := "SELECT COUNT(*) FROM " + tableName + " WHERE " + dbFiled + " = ?"
		if exceptId != "" {
			query += " AND id != " + exceptId
		}
		var count int
		// database.DB.QueryRow(sql, emial).Scan(&count)
		database.DB.QueryRow(query, requestValue, exceptId).Scan(&count)
		if count != 0 {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 已被占用", requestValue)
		}
		return nil
	})
}
