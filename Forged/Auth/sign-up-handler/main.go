package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB *sql.DB
}

type signupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requests signupRequest
	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}
	if requests.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Input your email!",
		})
		return
	}

	if len(requests.Password) < 8 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "password must at least eight(8) characters!",
		})
		return
	}
	hash, err := HashPassword(requests.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "",
		})
		return
	}
	_, err = h.DB.Exec(
		"INSERT INTO users (email, password_hash) VALUES ($1, $2)",
		requests.Email, hash,
	)
	if err != nil {
		log.Println("insert error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to insert ROw",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "account created"})

}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on real environment variables")
	}

	connstr := os.Getenv("DATABASE_URL")
	if connstr == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schema := `CREATE TABLE IF NOT EXISTS users (
    	id SERIAL PRIMARY KEY,
   		email TEXT UNIQUE NOT NULL,
   		password_hash TEXT NOT NULL,
    	verified BOOLEAN NOT NULL DEFAULT FALSE,
    	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);`
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	handler := &AuthHandler{DB: db}

	http.HandleFunc("/signup", handler.Signup)

	log.Println("listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
