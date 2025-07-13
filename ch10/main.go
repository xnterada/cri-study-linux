package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/lib/pq"
)

type Application struct {
    db *sql.DB
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type NewUserRequest struct {
	Username string `json:"username"`
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("usage: ./server <http_port_number> <db_hostname>")
		os.Exit(1)
	}
	portStr := os.Args[1]	
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port number: %s\n", portStr)
		os.Exit(1)
	}
	dbHost := os.Args[2]

	db, err := initDB(dbHost)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	app := &Application{db: db}

	http.HandleFunc("/users", app.usersHandler)

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on port %s...", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initDB(dbHost string) (*sql.DB, error) {
	dbPort := "5432"
	dbName := "user_db"
	dbUser := "postgres"
	sslMode := "disable"

	connStr := fmt.Sprintf("user=%s host=%s port=%s dbname=%s sslmode=%s", dbUser, dbHost, dbPort, dbName, sslMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("Successfully connected to the database.")
	return db, nil
}

func (app *Application) usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		app.createUser(w, r)
	default:
		sendJSONResponse(w, http.StatusMethodNotAllowed, nil, "Method Not Allowed")
	}
}

func (app *Application) createUser(w http.ResponseWriter, r *http.Request) {
	var req NewUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSONResponse(w, http.StatusBadRequest, nil, "Invalid JSON payload")
		return
	}

	if req.Username == "" {
		sendJSONResponse(w, http.StatusBadRequest, nil, "Username is required")
		return
	}

	var newUser User
	sqlStatement := `INSERT INTO users (username) VALUES ($1) RETURNING id, username, created_at`
	err := app.db.QueryRow(sqlStatement, req.Username).Scan(&newUser.ID, &newUser.Username, &newUser.CreatedAt)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code.Name() == "unique_violation" {
			sendJSONResponse(w, http.StatusConflict, nil, fmt.Sprintf("Username '%s' already exists", req.Username))
			return
		}
		log.Printf("Error inserting user: %v", err)
		sendJSONResponse(w, http.StatusInternalServerError, nil, "Failed to create user")
		return
	}

	sendJSONResponse(w, http.StatusCreated, newUser, "User created successfully")
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := make(map[string]interface{})

	if statusCode >= 200 && statusCode < 300 {
		response["message"] = msg
		if data != nil {
			response["data"] = data
		}
	} else {
		response["error"] = msg
	}

    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Printf("Failed to encode JSON response: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
