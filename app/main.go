package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/RahulKraken/Paste-it/database"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

// database connection
var db *sql.DB
var err error

// handlers
// create new user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit", "/api/user", "Method:", r.Method)
	decoder := json.NewDecoder(r.Body)
	var user database.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err.Error())
	}

	// if everything goes fine then persist user into database
	log.Println("sending to db:", user)
	database.CreateUser(db, user)
}

// update existing user
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user database.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err.Error())
	}

	// if everything goes fine then update user in database
	log.Println(user)
	database.UpdateUser(db, user)
}

// get user with id
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("id:", vars["id"])
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	val := database.GetUser(db, id)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(val)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
}

// delete user with id
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("id:", vars["id"])
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}
	
	database.DeleteUser(db, id)
}

// create new paste
func createPasteHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var paste database.Paste
	err := decoder.Decode(&paste); if err != nil {
		panic(err.Error())
	}

	log.Println(paste)
	log.Println("inserting into db")
	database.CreatePaste(db, paste)
}

// update existing paste
func updatePasteHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var paste database.Paste
	err := decoder.Decode(&paste); if err != nil {
		panic(err.Error())
	}

	log.Println(paste)
	log.Println("updating in db")
	database.UpdatePaste(db, paste)
}

// get paste with id
func getPasteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("paste id:", vars["id"])
	id, err := strconv.Atoi(vars["id"]); if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	} 

	val := database.GetPaste(db, id)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(val); if err != nil {
		log.Println(err)
		http.Error(w, "Error parsing result", http.StatusInternalServerError)
		return
	}
}

// delete paste with id
func deletePasteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("paste id:", vars["id"])
	id, err := strconv.Atoi(vars["id"]); if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	}

	log.Println("deleting from db")
	database.DeletePaste(db, id)
}

func main() {
	fmt.Println("Hello from main!!!")
	// try connecting to database
	db, err = sql.Open("mysql", "pasteit:pasteit@tcp(127.0.0.1:3306)/pasteit_db")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	log.Println("Database connection successfull...")

	defer db.Close()

	fmt.Println("testing db here ->", db)

	// mux router
	r := mux.NewRouter()

	// user handlers
	r.HandleFunc("/api/user", createUserHandler).Methods("POST")
	r.HandleFunc("/api/user", updateUserHandler).Methods("PUT")
	r.HandleFunc("/api/user/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/api/user/{id}", deleteUserHandler).Methods("DELETE")

	// paste handlers
	r.HandleFunc("/api/paste", createPasteHandler).Methods("POST")
	r.HandleFunc("/api/paste", updatePasteHandler).Methods("PUT")
	r.HandleFunc("/api/paste/{id}", getPasteHandler).Methods("GET")
	r.HandleFunc("/api/paste/{id}", deletePasteHandler).Methods("DELETE")

	// listen and serve
	err = http.ListenAndServe(":5000", r)
	fmt.Println("Server started on port: 5000")
	if err != nil {
		panic(err.Error())
	}
}