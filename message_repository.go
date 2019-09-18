package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MessageRepository struct {
	conn *sql.DB
}

func NewMessageRepository() *MessageRepository {
	messageName := "graphql"
	password := "graphql"
	host := "localhost"
	port := 3306
	database := "go_graphql_sample"
	uri := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		messageName,
		password,
		host,
		port,
		database,
	)

	conn, err := sql.Open("mysql", uri)
	if err != nil {
		panic(err.Error())
	}
	return &MessageRepository{
		conn: conn,
	}
}

func (repo *MessageRepository) FindAll() Messages {
	rows, err := repo.conn.Query("SELECT * FROM `messages`")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	messages := Messages{}

	for rows.Next() {
		message := &Message{}
		err := rows.Scan(&message.ID, &message.UserID, &message.Text)
		if err != nil {
			continue
		}
		messages = append(messages, message)
	}

	return messages
}

func (repo *MessageRepository) Find(userID int64) (*Message, error) {
	rows, err := repo.conn.Query(
		"SELECT * FROM `messages` WHERE `messages`.id = ?",
		userID,
	)
	if err != nil {
		return nil, err
	}

	message := &Message{}
	rows.Next()

	err = rows.Scan(&message.ID, &message.UserID, &message.Text)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (repo *MessageRepository) Create(message Message) int64 {
	result, err := repo.conn.Exec("INSERT INTO `messages` (`name`) VALUES (?,?)", message.UserID, message.Text)
	if err != nil {
		return 0
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0
	}

	return int64(id)
}

func (repo *MessageRepository) Delete(id int64) int64 {
	result, err := repo.conn.Exec("DELETE FROM `messages` WHERE `id` = ?", id)
	if err != nil {
		return 0
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0
	}

	return int64(count)
}
