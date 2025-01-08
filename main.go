package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type user struct {
	id int
	name string
}

func main() {
	fmt.Println("Начало положено :)")

	connectionString := "user=bogdan password=123 dbname=blog-db sslmode=disable port=5430"

	db, err := sql.Open("postgres", connectionString)
    
	if err != nil {
        panic(err)
    }

    defer db.Close()

	rows, err := db.Query("select * from users")
	
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	usersArray := []user{}

	for rows.Next() {
		user := user{}

		err := rows.Scan(&user.id, &user.name)

		if err != nil {
			fmt.Println(err)
			continue
		}

		usersArray = append(usersArray, user)
	}

	for _, user := range usersArray {
		fmt.Println(user.id, user.name)
	}

}