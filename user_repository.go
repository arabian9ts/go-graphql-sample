package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	conn *sql.DB
}

func NewUserRepository() *UserRepository {
	userName := "graphql"
	password := "graphql"
	host := "localhost"
	port := 3306
	database := "go_graphql_sample"
	uri := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		userName,
		password,
		host,
		port,
		database,
	)

	conn, err := sql.Open("mysql", uri)
	if err != nil {
		panic(err.Error())
	}
	return &UserRepository{
		conn: conn,
	}
}

func (repo *UserRepository) FindAll() Users {
	rows, err := repo.conn.Query("SELECT * FROM `users`")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	users := Users{}

	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			continue
		}
		users = append(users, user)
	}

	return users
}

func (repo *UserRepository) Find(id int64) (*User, error) {
	rows, err := repo.conn.Query("SELECT * FROM `users` WHERE `id` = ?", id)
	if err != nil {
		return nil, err
	}

	user := &User{}
	rows.Next()

	err = rows.Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) Create(user User) int64 {
	result, err := repo.conn.Exec("INSERT INTO `users` (`name`) VALUES (?)", user.Name)
	if err != nil {
		return 0
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0
	}

	return int64(id)
}

func (repo *UserRepository) Delete(id int64) int64 {
	result, err := repo.conn.Exec("DELETE FROM `users` WHERE `id` = ?", id)
	if err != nil {
		return 0
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0
	}

	return int64(count)
}
