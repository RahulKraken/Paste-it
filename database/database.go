package database

import (
	"database/sql"
	"log"
)

// User struct
type User struct {
	ID		 int 	`json:"id"`
	UserName string `json:"user_name"`
}

// Paste struct
type Paste struct {
	ID 		int 	`json:"id"`
	UserID 	int 	`json:"user_id"`
	Title 	string 	`json:"title"`
	Content string 	`json:"content"`
}

// Mapping struct
type Mapping struct {
	ID		int 	`json:"paste_id"`
	Hash	string	`json:"paste_hash"`
}

// CRUD application methods for user table

// ListUsers - list all available users
func ListUsers(db *sql.DB) []User {
	results, err := db.Query("SELECT * FROM user")
	if err != nil {
		// could use log.Panicln() but still
		log.Println("Error fetching user list:", err)
		panic(err.Error())
	}

	// to hold the resutls
	var users []User
	var user User
	for results.Next() {
		// scan tuple into the user var
		err = results.Scan(&user.ID, &user.UserName)
		if err != nil {
			log.Println("Error parsing user:", err)
			// return whatever's parsed so far
			return users
		}

		// append to users slice
		users = append(users, user)
	}

	return users
}

// CreateUser - create new user
func CreateUser(db *sql.DB, user User) {
	insert, err := db.Query("INSERT INTO user VALUES(?, ?)", user.ID, user.UserName)
	if err != nil {
		log.Println("db error:", err)
		panic(err.Error())
	}
	log.Println("Inserting into db:", user)
	insert.Close()
}

// GetUser - get user with id
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

// UpdateUser - update user
func UpdateUser(db *sql.DB, user User) User {
	_, err := db.Query("UPDATE user SET user.user_name = ? WHERE id = ?", user.UserName, user.ID)
	if err != nil {
		panic(err.Error())
	}
	
	return user
}

// DeleteUser - delete user with id
func DeleteUser(db *sql.DB, id int) {
	_, err := db.Query("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
}

// CRUD application methods for paste table

// CreatePaste - create new paste
func CreatePaste(db *sql.DB, paste Paste) {
	insert, err := db.Query("INSERT INTO paste VALUES (?, ?, ?, ?)", paste.ID, paste.UserID, paste.Title, paste.Content)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

// UpdatePaste - update existing paste
func UpdatePaste(db *sql.DB, paste Paste) Paste {
	_, err := db.Query("UPDATE paste SET title = ?, content = ? WHERE id = ? AND user_id = ?", paste.Title, paste.Content, paste.ID, paste.UserID)
	if err != nil {
		panic(err.Error())
	}

	return paste
}

// GetPaste - get paste with id
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

// DeletePaste - delete paste with id
func DeletePaste(db *sql.DB, id int) {
	_, err := db.Query("DELETE FROM paste WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
}

// CRUD application methods for mapping table

// CreateMapping - create a new mapping for a paste
func CreateMapping(db *sql.DB, mapping Mapping) (string, error) {
	_, err := db.Query("INSERT INTO mapping VALUES (?, ?)", mapping.ID, mapping.Hash)
	if err != nil {
		// log the error details
		log.Println("Error inserting mapping for:", mapping.ID, "with hash:", mapping.Hash, "->", err)
	}

	return mapping.Hash, err
}

// GetMapping - get mapping for a paste
func GetMapping(db *sql.DB, pasteID int) (string, error) {
	val, err := db.Query("SELECT paste_hash FROM mapping WHERE paste_id = ? LIMIT 1", pasteID)
	if err != nil {
		// log the error details
		log.Println("Error getting mapping for:", pasteID, "->", err)
	}

	var hash string
	for val.Next() {
		err = val.Scan(&hash); if err != nil {
			// log the error
			log.Println("Error parsing hash from results: err ->", err, "\nresult ->", val)
		}
	}

	return hash, err
}

// DeleteMapping - delete a mapping in table
func DeleteMapping(db *sql.DB, pasteID int) error {
	_, err := db.Query("DELETE FROM mapping WHERE paste_id = ?", pasteID)
	if err != nil {
		// log the error
		log.Println("Error deleting entry for paste_id =", pasteID)
	}

	return err
}

// ExistsMapping - check if a hash exists already
func ExistsMapping(db *sql.DB, h string) bool {
	res, err := db.Query("SELECT * FROM mapping WHERE paste_hash = ?", h)
	if err != nil || !res.Next() {
		// log `h` is available
		log.Println(h, "is available")
		return false
	}

	// log `h` is already in use
	log.Println(h, "is already being used.")
	return true
}