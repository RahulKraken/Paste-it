package database

import (
	"fmt"
	"database/sql"
	// import and use as driver
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

// createUser - create new user
func createUser(db *sql.DB, user User) {
	insert, err := db.Query("INSERT INTO user VALUES(?, ?)", user.ID, user.UserName)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

// getUser - get user with id
func getUser(db *sql.DB, id int) User {
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

// updateUser - update user
func updateUser(db *sql.DB, user User) User {
	_, err := db.Query("UPDATE user SET user.user_name = ? WHERE id = ?", user.UserName, user.ID)
	if err != nil {
		panic(err.Error())
	}
	
	return user
}

// deleteUser - delete user with id
func deleteUser(db *sql.DB, id int) {
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

// createPaste - create new paste
func createPaste(db *sql.DB, paste Paste) {
	insert, err := db.Query("INSERT INTO paste VALUES (?, ?, ?, ?)", paste.ID, paste.UserID, paste.Title, paste.Content)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

// updatePaste - update existing paste
func updatePaste(db *sql.DB, paste Paste) Paste {
	_, err := db.Query("UPDATE paste SET title = ?, content = ? WHERE id = ? AND user_id = ?", paste.Title, paste.Content, paste.ID, paste.UserID)
	if err != nil {
		panic(err.Error())
	}

	return paste
}

// getPaste - get paste with id
func getPaste(db *sql.DB, id int) Paste {
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

// deletePaste - delete paste with id
func deletePaste(db *sql.DB, id int) {
	_, err := db.Query("DELETE FROM paste WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
}

// exposed CRUD functions
// User functions

// CreateUser - new user
func CreateUser(user User) {
	createUser(db, user)
}

// UpdateUser - update existing user
func UpdateUser(user User) {
	updateUser(db, user)
}

// GetUser - get user with id
func GetUser(id int) User {
	return getUser(db, id)
}

// DeleteUser - delete user with id
func DeleteUser(id int) {
	deleteUser(db, id)
}

// Paste functions

// CreatePaste - create new paste
func CreatePaste(paste Paste) {
	createPaste(db, paste)
}

// UpdatePaste - update paste
func UpdatePaste(paste Paste) {
	updatePaste(db, paste)
}

// GetPaste - get paste with id
func GetPaste(id int) Paste {
	return getPaste(db, id)
}

// DeletePaste - delete paste with id
func DeletePaste(id int) {
	deletePaste(db, id)
}

func init() {
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