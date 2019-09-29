package database

import (
	"database/sql"
	"log"
)

// User struct
type User struct {
	ID		 int 	`json:"id"`
	Email	 string	`json:"email"`
	UserName string `json:"user_name"`
	Pasword  string `json:"pasword"`
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

// LoginCredentials struct
type LoginCredentials struct {
	Username 	string	 	`json:"username"`
	Pasword 	string		`json:"pasword"`
}

// CRUD application methods for user table

// ListUsers - list all available users
func ListUsers(db *sql.DB) []User {
	results, err := db.Query("SELECT id, user_name FROM user")
	if err != nil {
		// could use log.Panicln() but still
		log.Println("Error fetching user list:", err)
		panic(err.Error())
	}

	// to hold the results
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
func CreateUser(db *sql.DB, user User) int {
	// get count
	res, err := db.Query("SELECT COUNT(id) FROM user")
	if err != nil {
		log.Println("Error creating user")
		panic(err.Error())
	}

	var cnt int
	for res.Next() {
		_ = res.Scan(&cnt)
	}

	log.Println("cnt:", cnt)

	insert, err := db.Query("INSERT INTO user VALUES(?, ?, ?, ?)", cnt + 1, user.Email, user.UserName, user.Pasword)
	if err != nil {
		log.Println("Error creating user")
		panic(err.Error())
	}
	log.Println("inserting into db:", cnt + 1, user)
	_ = insert.Close()

	return cnt + 1
}

// GetUser - get user with id
func GetUser(db *sql.DB, id int) User {
	data, err := db.Query("SELECT id, user_name FROM user WHERE id = ? LIMIT 1", id)
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

// ListPastes - list all pastes with userID = ID
func ListPastes(db *sql.DB, userID int) []Paste {
	results, err := db.Query("SELECT * FROM paste WHERE user_id = ?", userID)
	if err != nil {
		log.Println("Error fetching pastes for user:", userID, "err:", err)
		panic(err.Error())
	}

	// var to hold data
	var pastes []Paste
	var paste Paste

	for results.Next() {
		err = results.Scan(&paste.ID, &paste.UserID, &paste.Title, &paste.Content)
		if err != nil {
			log.Println("Error parsing results:", err)
			panic(err.Error())
		}

		// append to slice
		pastes = append(pastes, paste)
	}

	return pastes
}

// CreatePaste - create new paste
func CreatePaste(db *sql.DB, paste Paste) int {
	// get cnt of pastes
	res, err := db.Query("SELECT COUNT(id) FROM paste")
	if err != nil {
		log.Println("Error creating paste")
		panic(err.Error())
	}

	var cnt int
	for res.Next() {
		_ = res.Scan(&cnt)
	}

	log.Println("cnt:", cnt)

	// insert paste into db
	insert, err := db.Query("INSERT INTO paste VALUES (?, ?, ?, ?)", cnt + 1, paste.UserID, paste.Title, paste.Content)
	if err != nil {
		log.Println("Error creating paste")
		panic(err.Error())
	}

	log.Println("inserting into db:", cnt + 1, paste)
	_ = insert.Close()

	return cnt + 1
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

// GetPasteIdFromHash - get pasteId corresponding to given hash
func GetPasteIdFromHash(db *sql.DB, hash string) int {
	val, err := db.Query("SELECT paste_id FROM mapping WHERE paste_hash = ?", hash)
	if err != nil {
		panic(err.Error())
	}

	var id int
	for val.Next() {
		err = val.Scan(&id); if err != nil {
			panic(err.Error())
		}
	}

	return id
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

// ExistsEmail - check if account with email exists
func ExistsEmail(db *sql.DB, email string) bool {
	res, err := db.Query("SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		log.Println(err)
	}
	if res.Next() {
		return true
	}
	return false
}

// ExistsUsername - check if account with username exists
func ExistsUsername(db *sql.DB, username string) bool {
	res, err := db.Query("SELECT * FROM user WHERE user_name = ?", username)
	if err != nil {
		log.Println(err)
	}
	if res.Next() {
		return true
	}
	return false
}

// MatchCredentials - check if password matches
func MatchCredentials(db *sql.DB, data LoginCredentials) bool {
	res, err := db.Query("SELECT pasword from USER WHERE user_name = ?", data.Username)
	if err != nil {
		log.Println("error verifying pasword", err)
	}

	if res.Next() {
		var psd string
		_ = res.Scan(&psd)
		if psd == data.Pasword {
			return true
		}
	}

	return false
}