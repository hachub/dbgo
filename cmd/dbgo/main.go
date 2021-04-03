package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Calc() int {
	return 1
}
func main() {
	fmt.Println("Hello")
	db, err := sql.Open("mysql", "test:test@tcp(127.0.0.1:43306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM main")
	if rows == nil {
		fmt.Println(err)
		panic("IS NIL")
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var todo string
		rows.Scan(&id, &todo)
		fmt.Println(id, todo)
	}
}
