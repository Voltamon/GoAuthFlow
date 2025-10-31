package data

import (
	"database/sql"
	"fmt"
	"time"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) RegisterUser(user *User) error {
	sqlSmt := "INSERT INTO Users (hash, created_at) VALUES (?, ?)"
	user.CreatedAt = time.Now()

	res, err := r.DB.Exec(sqlSmt, user.Hash, user.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to register user: %v", err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last inserted ID: %v", err)
	}

	return nil
}

func (r *UserRepo) GetUserbyHash(user *User) (*User, error) {
	sqlSmt := "SELECT id, hash, created_at FROM Users WHERE hash = ?"
	existingUser := &User{}

	err := r.DB.QueryRow(sqlSmt, user.Hash).Scan(&existingUser.ID, &existingUser.Hash, &existingUser.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user by hash: %v", err)
	}

	return existingUser, nil
}

func (r *UserRepo) GetAllUsers() ([]User, error) {
	sqlSmt := "SELECT id, hash, created_at FROM Users"
	var users []User

	rows, err := r.DB.Query(sqlSmt)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}

	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Hash, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, *user)
	}

	return users, nil
}
