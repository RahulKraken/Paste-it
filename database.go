package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// User struct
type User struct {
	ID int `json:"id"`
	UserName string `json:"user_name"`
}

// db connection var
var db *sql.DB

// CRUD application methods for user table
// create user
func CreateUser(db *sql.DB, user User) {
	insert, err := db.Query("INSERT INTO user VALUES(?, ?)", user.ID, user.UserName)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

// get user using id
func GetUser(db *sql.DB, id int) User {
	data, err := db.Query("SELECT * FROM user WHERE id = ? LIMIT 1", id)
	if err != nil {
		panic(err.Error())
	}

	var user User
	for data.Next() {
		err = data.Scan(&user.ID, &user.UserName)
		if err != nil {
			panic(err.Error())
		}
	}

	return user
}

// update user using id
func UpdateUser(db *sql.DB, user User) User {
	_, err := db.Query("UPDATE user SET user.user_name = ? WHERE id = ?", user.UserName, user.ID)
	if err != nil {
		panic(err.Error())
	}
	
	return user
}

// delete user using id
func DeleteUser(db *sql.DB, id int) {
	_, err := db.Query("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
}

// Paste struct
type Paste struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

// CRUD application methods for paste table
// create paste
func CreatePaste(db *sql.DB, paste Paste) {
	insert, err := db.Query("INSERT INTO paste VALUES (?, ?, ?, ?)", paste.ID, paste.UserID, paste.Title, paste.Content)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

// update paste
func UpdatePaste(db *sql.DB, paste Paste) Paste {
	_, err := db.Query("UPDATE paste SET title = ?, content = ? WHERE id = ? AND user_id = ?", paste.Title, paste.Content, paste.ID, paste.UserID)
	if err != nil {
		panic(err.Error())
	}

	return paste
}

// get paste
func GetPaste(db *sql.DB, id int) Paste {
	val, err := db.Query("SELECT * FROM paste WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	var paste Paste
	for val.Next() {
		err = val.Scan(&paste.ID, &paste.UserID, &paste.Title, &paste.Content)
		if err != nil {
			panic(err.Error())
		}
	}

	return paste
}

// delete post
func DeletePaste(db *sql.DB, id int) {
	_, err := db.Query("DELETE FROM paste WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	// try connecting to database
	db, err := sql.Open("mysql", "pasteit:pasteit@tcp(127.0.0.1:3306)/pasteit_db")
	if err != nil {
		panic(err.Error())
	}

	// trying to ping database
	err = db.Ping();
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("database connection successfull!!!")
}