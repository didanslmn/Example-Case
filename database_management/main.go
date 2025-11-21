package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

// repository pattern

type UserRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user User) error {
	script := "INSERT INTO users(name, email) VALUES($1, $2)"
	_, err := r.db.Exec(script, user.Name, user.Email)
	return err
}

func (r *UserRepository) GetByID(id int) (*User, error) {
	script := "SELECT id,name,email,created_at FROM users WHERE id = $1"
	row := r.db.QueryRow(script, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func main() {
	// koneksi ke postgres
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// test koneksi
	if err := db.Ping(); err != nil {
		panic(err)

	}
	fmt.Println("koneksi berhasil")

	// initial repository
	repo := NewRepository(db)

	// Contoh penggunaan
	user := User{Name: "budi", Email: "budi@email.com"}
	if err := repo.Create(user); err != nil {
		log.Println("Error create:", err)
	}

	found, err := repo.GetByID(1)
	if err != nil {
		log.Println("Error get:", err)
	} else {
		fmt.Printf("User: %+v\n", found)
	}
}
