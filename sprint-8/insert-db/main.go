package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

func NewDB() (DB, error) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		return DB{}, nil
	}

	return DB{db}, nil
}

type Client struct {
	FIO      string
	Login    string
	Birthday string
	Email    string
}

func (d DB) insertClient(client Client) error {
	res, err := d.Exec("INSERT INTO clients (fio, login, birthday, email) VALUES (:fio, :login, :birthday, :email)",
		sql.Named("fio", client.FIO),
		sql.Named("login", client.Login),
		sql.Named("birthday", client.Birthday),
		sql.Named("email", client.Email))
	if err != nil {
		err = fmt.Errorf("db exec error: %w", err)
		return err
	}

	log.Println(res.LastInsertId())
	log.Println(res.RowsAffected())

	return nil
}

func (d DB) updateClientLogin(id int, login string) error {
	res, err := d.Exec("UPDATE clients SET login = :login WHERE id = :id",
		sql.Named("id", id),
		sql.Named("login", login))
	if err != nil {
		err = fmt.Errorf("db exec error: %w", err)
		return err
	}

	log.Println(res.LastInsertId())
	log.Println(res.RowsAffected())

	return nil
}

func (d DB) deleteClient(id int) error {
	res, err := d.Exec("DELETE FROM clients WHERE id = :id",
		sql.Named("id", id))
	if err != nil {
		err = fmt.Errorf("db exec error: %w", err)
		return err
	}

	log.Println(res.LastInsertId())
	log.Println(res.RowsAffected())

	return nil
}

func (d DB) selectClient(id int) (Client, error) {
	var client Client

	row := d.QueryRow("SELECT fio, login, birthday, email FROM clients WHERE id = :id",
		sql.Named("id", id))

	err := row.Scan(&client.FIO, &client.Login, &client.Birthday, &client.Email)
	if err != nil {
		return Client{}, err
	}

	return client, nil
}

func main() {
	db, err := NewDB()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	newClient := Client{
		FIO:      "Сергей Косовский",
		Login:    "skosovsky2",
		Birthday: "19800721",
		Email:    "skosovsky@gmail.com",
	}

	err = db.insertClient(newClient)
	if err != nil {
		log.Println(err)
	}

	// TODO: Add method from tasks

	// test func
	err = insertProduct(db.DB, "Облачное хранилище", 300)
	if err != nil {
		log.Println(err)
	}

	err = updatePriceProduct(db.DB, 1, 700)
	if err != nil {
		log.Println(err)
	}

	err = deleteClient(db.DB, 3)
	if err != nil {
		log.Println(err)
	}
}

func insertProduct(db *sql.DB, product string, price int) error {
	res, err := db.Exec("INSERT INTO products (product, price) VALUES (:product, :price)",
		sql.Named("product", product),
		sql.Named("price", price))
	if err != nil {
		err = fmt.Errorf("db exec error: %w", err)
		return err
	}

	log.Println(res.LastInsertId())
	log.Println(res.RowsAffected())

	return nil
}

func updatePriceProduct(db *sql.DB, productID, price int) error {
	res, err := db.Exec("UPDATE products SET price = :price WHERE id = :product_id",
		sql.Named("product_id", productID),
		sql.Named("price", price))
	if err != nil {
		err = fmt.Errorf("db exec error: %w", err)
		return err
	}

	log.Println(res.LastInsertId())
	log.Println(res.RowsAffected())

	return nil
}

func deleteClient(db *sql.DB, clientID int) error {
	res, err := db.Exec("DELETE FROM clients WHERE  id = :client_id", sql.Named("client_id", clientID))
	if err != nil {
		err = fmt.Errorf("db exec error: %w", err)
		return err
	}

	log.Println(res.LastInsertId())
	log.Println(res.RowsAffected())

	return nil
}
