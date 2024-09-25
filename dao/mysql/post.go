package mysql

import "fmt"

func Query() {
	rows, err := db.Query("SELECT id, title FROM post")
	if err != nil {
		fmt.Println("Query failed:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		err := rows.Scan(&id, &title)
		if err != nil {
			fmt.Println("Failed to Scan row:", err)
			return
		}
		fmt.Printf("ID: %d, Name: %s\n", id, title)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Row error", err)
	}
}
