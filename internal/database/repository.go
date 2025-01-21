package database

import (
	"database/sql"

	"github.com/Yandex-Practicum/final-project-encoding-go/internal"
	_ "modernc.org/sqlite"
)

// InsertClient метод для добавления нового клиента в базу
func InsertClient(db *sql.DB, client *internal.Client) (int64, error) {
	// напишите здесь код для добавления новой записи в таблицу clients
	result, err := db.Exec("INSERT INTO clients (fio, login, birthday, email) VALUES (:fio, :login, :birthday, :email)",
		sql.Named("fio", client.FIO),
		sql.Named("login", client.Login),
		sql.Named("birthday", client.Birthday),
		sql.Named("email", client.Email))

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateClientLogin метода для обновления логина по клиенту в базе
func UpdateClientLogin(db *sql.DB, login string, id int64) error {
	// напишите здесь код для обновления поля login в таблице clients у записи с заданным id
	_, err := db.Exec("UPDATE clients SET login = :login WHERE id = :id",
		sql.Named("login", login),
		sql.Named("id", id))

	if err != nil {
		return err
	}
	return nil
}

// DeleteClient метод для удаления клиента по id из базы
func DeleteClient(db *sql.DB, id int64) error {
	// напишите здесь код для удаления записи из таблицы clients по заданному id
	_, err := db.Exec("DELETE FROM clients WHERE id = :id",
		sql.Named("id", id))

	if err != nil {
		return err
	}
	return nil
}

// SelectClient посик клиента в баз по id
func SelectClient(db *sql.DB, id int64) (internal.Client, error) {
	client := internal.Client{}

	row := db.QueryRow("SELECT id, fio, login, birthday, email FROM clients WHERE id = :id", sql.Named("id", id))
	err := row.Scan(&client.ID, &client.FIO, &client.Login, &client.Birthday, &client.Email)

	return client, err
}
