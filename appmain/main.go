package main

import (
	"fmt"
    _ "com.jtthink/services"
	// . "com.jtthink/core"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"os"
	_ "com.jtthink/models"
)

type all interface {

}

func test(str ...string) {
	for _, v := range str {
		fmt.Println(v)
	}
}

func main()  {
	// fmt.Println(GetService().Get(123))

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4")
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:"+ err.Error())
		os.Exit(1)
	}

	rows, err := db.Query("select user_name, signup_at from tbl_user")
	if err != nil {
		fmt.Println("Failed to select to mysql, err:"+ err.Error())
		return
	}

	columns, _ := rows.Columns()

	allRows := make([]interface{}, 0)
	for rows.Next() {
	    fileMap := make(map[string]interface{})
		row := make([]interface{}, len(columns))
		scanRow := make([]interface{}, len(columns))
		for i, _ := range row {
			scanRow[i] = &row[i]
		}
		// rows.Scan(&row[0], &row[1])
		rows.Scan(scanRow...)

		for k, val := range row {
			v, ok := val.([]byte)
			if ok {
				// row[k] = string(v)
				fileMap[columns[k]] = string(v)
			}
		}

		allRows = append(allRows, fileMap)
	}

	for _, v := range allRows {
		fmt.Println(v)
	}
}