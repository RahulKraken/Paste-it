package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/RahulKraken/Paste-it/database"
	"github.com/RahulKraken/Paste-it/hash"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// database connection
var db *sql.DB
var err error

// CORS middleware
type CORSDecorator struct {
	R *mux.Router
}

// CORSDecorator - it will apply the CORS headers to the mux Router
func (c *CORSDecorator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	// set the headers
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Token")
	}

	// if preflight OPTIONS request then stop here
	if r.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(w, r)
}

/**
not being used right now
 */
// auth handlers

// authentication middleware
func handleAuth(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			log.Println("Token:", r.Header["Token"][0])
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (i interface{}, e error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC); if !ok {
					return nil, fmt.Errorf("something wrong happened")
				}

				return hash.MySigningKey, nil
			})

			if err != nil {
				log.Println("Something wrong happened:", err.Error())
				http.Error(w, "Could not authenticate", http.StatusUnauthorized)
				return
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			log.Println("Unauthorized")
			http.Error(w, "Error authenticating", http.StatusUnauthorized)
		}
	})
}

// signUpHandler - handles signup requests
func signUpHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit", "/signup", r.Method)
	var user database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user); if err != nil {
		log.Println("Could not parse request", err)
	}

	log.Println("username:", user.UserName, "; email:", user.Email, "; pasword:", user.Pasword)

	// check if email already exists
	b := database.ExistsEmail(db, user.Email)
	if b {
		// email already exists
		log.Println("User exists")
		http.Error(w, "user with this email already exists", http.StatusBadRequest)
		return
	}

	// check if username exists
	b = database.ExistsUsername(db, user.UserName)
	if b {
		// username already in use
		log.Println("Username is taken")
		http.Error(w, "Username is already in use", http.StatusBadRequest)
		return
	}

	// create user
	_ = database.CreateUser(db, user)

	// generate and send JWT
	token, err := hash.GenerateJWT(user.UserName)
	if err != nil {
		log.Println("Error generating JWT")
		http.Error(w, "Something wrong happened", http.StatusInternalServerError)
	}

	createdUser := database.GetUserWithUsername(db, user.UserName)

	// anonymous struct to send token
	response := struct {
		AuthToken		string		`json:"token"`
		Id				int			`json:"id"`
		Username		string		`json:"username"`
		Email			string		`json:"email"`
	}{
		AuthToken:	token,
		Id:			createdUser.ID,
		Username:	createdUser.UserName,
		Email:		createdUser.Email,
	}

	log.Println(response)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		log.Println("Error sending JWT token")
		http.Error(w, "Something wrong happened", http.StatusInternalServerError)
	}
}

// loginHandler - handles login requests
func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit", "/login", r.Method)
	var data database.LoginCredentials
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data); if err != nil {
		log.Println("Could not parse request", err)
	}

	log.Println("username:", data.Username, "; pasword:", data.Pasword)

	// check if username exists
	b := database.ExistsUsername(db, data.Username)
	if !b {
		log.Println("User does not exist")
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}

	// check if credentials match
	b = database.MatchCredentials(db, data)
	if !b {
		log.Println("Incorrect password")
		http.Error(w, "Incorrect password", http.StatusBadRequest)
		return
	}

	// generate and send jwt
	token, err := hash.GenerateJWT(data.Username)
	if err != nil {
		log.Println("Error generating JWT")
		http.Error(w, "Something wrong happened", http.StatusInternalServerError)
	}

	user := database.GetUserWithUsername(db, data.Username)

	// anonymous struct to send token
	response := struct {
		AuthToken		string		`json:"token"`
		Id				int			`json:"id"`
		Username		string		`json:"username"`
		Email			string		`json:"email"`
	}{
		AuthToken:	token,
		Id:			user.ID,
		Username:	user.UserName,
		Email:		user.Email,
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		log.Println("Error sending JWT token")
		http.Error(w, "Something wrong happened", http.StatusInternalServerError)
	}
}

// user handlers

// list users
func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit", "/api/user", r.Method)
	users := database.ListUsers(db)
	log.Println(users)

	encoder := json.NewEncoder(w)

	err := encoder.Encode(users)
	if err != nil {
		log.Println("Error encoding result:", err)
		http.Error(w, "Something wrong occurred while preparing your result", http.StatusInternalServerError)
		return
	}
}

// create new user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit", "/api/user", r.Method)
	decoder := json.NewDecoder(r.Body)
	var user database.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err.Error())
	}

	// if everything goes fine then persist user into database
	log.Println("sending to db:", user)
	id := database.CreateUser(db, user)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(id)
	if err != nil {
		log.Println("error retrieving id", err)
		http.Error(w, "Error retrieving id", http.StatusInternalServerError)
	}
}

// update existing user
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT: /api/user", r.Method)
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
	log.Println("HIT: /api/user/{id}", r.Method)
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
	log.Println("HIT: /api/user/{id}", r.Method)
	vars := mux.Vars(r)
	fmt.Println("id:", vars["id"])
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}
	
	database.DeleteUser(db, id)
}

// paste handlers

// list paste handler
func listPastesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT: /api/pastes/{id}", r.Method)
	// get the id path variable
	vars := mux.Vars(r)
	log.Println("id:", vars["id"])
	id, err := strconv.Atoi(vars["id"]); if err != nil {
		panic(err.Error())
	}
	pastes := database.ListPastes(db, id)
	log.Println(pastes)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(pastes)
	if err != nil {
		log.Println("Error encoding results:", err)
		http.Error(w, "Error encoding results", http.StatusInternalServerError)
		return
	}
}

// create new paste
func createPasteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT: /api/paste", r.Method)
	decoder := json.NewDecoder(r.Body)
	var paste database.Paste
	err := decoder.Decode(&paste); if err != nil {
		panic(err.Error())
	}

	log.Println(paste)
	log.Println("inserting into db")
	id := database.CreatePaste(db, paste)

	log.Println("creating mapping")
	h := createMapping(id)

	encoder := json.NewEncoder(w)
	details := struct {
		ID		int		`json:"id"`
		Hash	string	`json:"h"`
	} {
		ID:   id,
		Hash: h,
	}

	err = encoder.Encode(details)
	if err != nil {
		log.Println("Error getting paste details", err)
		http.Error(w, "Error getting paste details", http.StatusInternalServerError)
	}
}

// update existing paste
func updatePasteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT: /api/paste", r.Method)
	decoder := json.NewDecoder(r.Body)
	var paste database.Paste
	err := decoder.Decode(&paste); if err != nil {
		log.Panicln("error parsing:", paste)
	}

	log.Println("updating in db:", paste)
	database.UpdatePaste(db, paste)
}

// get paste with id
func getPasteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT: /api/paste/{id}", r.Method)
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

// get paste with hash
func fetchPasteWithHash(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT: /paste/{hash}", r.Method)
	vars := mux.Vars(r)
	log.Println("paste hash:", vars["hash"])

	// get id associated with given hash
	id := database.GetPasteIdFromHash(db, vars["hash"])

	// fetch the paste from db
	paste := database.GetPaste(db, id)

	// return the content as text
	_, err = fmt.Fprintf(w, paste.Content)
	if err != nil {
		log.Println("Error sending response:", err)
		http.Error(w, "Error fetching content", http.StatusInternalServerError)
	}
}

// delete paste with id
func deletePasteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT: /api/paste/{id}", r.Method)
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

// create mapping
func createMapping(id int) string {
	created := true
	var h string
	for created {
		h = hash.Hash()
		created = database.ExistsMapping(db, h)
		if !created {
			_, err = database.CreateMapping(db, database.Mapping{
				ID: id,
				Hash: h,
			})
			if err != nil {
				log.Println("Error saving mapping to db for pasteID:", id)
			}
			created = false
		}
	}

	return h
}

func main() {
	fmt.Println("Hello from main!!!")
	// try connecting to database
	db, err = sql.Open("mysql", "pasteit:pasteit@tcp(127.0.0.1:3306)/pasteit")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	log.Println("Database connection successful...")

	defer db.Close()

	// mux router
	r := mux.NewRouter()

	// auth handlers
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/signup", signUpHandler).Methods("POST")

	// user handlers
	r.HandleFunc("/api/user", listUsersHandler).Methods("GET")
	r.HandleFunc("/api/user", createUserHandler).Methods("POST")
	r.HandleFunc("/api/user", updateUserHandler).Methods("PUT")
	r.HandleFunc("/api/user/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/api/user/{id}", deleteUserHandler).Methods("DELETE")

	// paste handlers
	r.HandleFunc("/paste/{hash}", fetchPasteWithHash).Methods("GET")
	r.HandleFunc("/api/pastes/{id}", listPastesHandler).Methods("GET")
	r.HandleFunc("/api/paste", createPasteHandler).Methods("POST")
	r.HandleFunc("/api/paste", updatePasteHandler).Methods("PUT")
	r.HandleFunc("/api/paste/{id}", getPasteHandler).Methods("GET")
	r.HandleFunc("/api/paste/{id}", deletePasteHandler).Methods("DELETE")

	fmt.Println("Server started on port: 5000")

	// listen and serve
	// using CORS middleware on the router
	err = http.ListenAndServe(":5000", &CORSDecorator{r})
	if err != nil {
		panic(err.Error())
	}
}