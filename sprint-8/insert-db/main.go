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
		err = fmt.Errorf("db open error: %w", err)
		return DB{DB: nil}, err
	}

	return DB{DB: db}, nil
}

type Client struct {
	ID       int
	FIO      string
	Login    string
	Birthday string
	Email    string
}

func (c Client) String() string {
	return fmt.Sprintf("ID: %d, FIO: %s, Login %s, Birthday: %s, Email %s",
		c.ID, c.FIO, c.Login, c.Birthday, c.Email)
}

func (d DB) insertClient(client Client) (int, error) {
	res, err := d.Exec("INSERT INTO clients (fio, login, birthday, email) VALUES (:fio, :login, :birthday, :email)",
		sql.Named("fio", client.FIO),
		sql.Named("login", client.Login),
		sql.Named("birthday", client.Birthday),
		sql.Named("email", client.Email))
	if err != nil {
		err = fmt.Errorf("db exec error: %w", err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		err = fmt.Errorf("get last insert id error %w: ", err)
		return 0, err
	}

	return int(id), nil
}

func (d DB) updateClientLogin(id int, login string) error {
	_, err := d.Exec("UPDATE clients SET login = :login WHERE id = :id",
		sql.Named("id", id),
		sql.Named("login", login))
	if err != nil {
		err = fmt.Errorf("db exec error: %w", err)
		return err
	}

	return nil
}

func (d DB) deleteClient(id int) error {
	_, err := d.Exec("DELETE FROM clients WHERE id = :id",
		sql.Named("id", id))
	if err != nil {
		err = fmt.Errorf("db exec error: %w", err)
		return err
	}

	return nil
}

func (d DB) selectClient(id int) (Client, error) {
	var client Client

	row := d.QueryRow("SELECT id, fio, login, birthday, email FROM clients WHERE id = :id",
		sql.Named("id", id))

	err := row.Scan(&client.ID, &client.FIO, &client.Login, &client.Birthday, &client.Email)
	if err != nil {
		err = fmt.Errorf("row scan error: %w", err)
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
	defer func(db *DB) {
		err = db.Close()
		if err != nil {
			log.Println("db close error")
		}
	}(&db)

	newClient := Client{ //nolint:exhaustruct // ID generated DB
		FIO:      "Сергей Косовский",
		Login:    "skosovsky2",
		Birthday: "19800721",
		Email:    "skosovsky@gmail.com",
	}

	id, err := db.insertClient(newClient)
	if err != nil {
		log.Println(err)
		return
	}

	selected, err := db.selectClient(id)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(selected) //nolint:forbidigo // it's learning code

	err = db.updateClientLogin(id, "newLogin")
	if err != nil {
		log.Println(err)
		return
	}

	selected, err = db.selectClient(id)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(selected) //nolint:forbidigo // it's learning code

	err = db.deleteClient(id)
	if err != nil {
		log.Println(err)
		return
	}

	selected, err = db.selectClient(id)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(selected) //nolint:forbidigo // it's learning code

	// test func
	err = insertProduct(db.DB, "Облачное хранилище", 300) //nolint:gomnd // it's learning code
	if err != nil {
		log.Println(err)
	}

	err = updatePriceProduct(db.DB, 1, 700) //nolint:gomnd // it's learning code
	if err != nil {
		log.Println(err)
	}

	err = deleteClient(db.DB, 3) //nolint:gomnd // it's learning code
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
