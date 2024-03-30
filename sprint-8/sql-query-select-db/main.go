package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date: %s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		err = fmt.Errorf("db open error: %w", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT product, volume, date FROM sales WHERE client = :client", sql.Named("client", client))
	if err != nil {
		err = fmt.Errorf("db query error: %w", err)
		return nil, err
	}
	defer rows.Close()

	var sales []Sale
	for rows.Next() {
		sale := Sale{
			Product: 0,
			Volume:  0,
			Date:    "",
		}

		err = rows.Scan(&sale.Product, &sale.Volume, &sale.Date)
		if err != nil {
			err = fmt.Errorf("rows scan error: %w", err)
			return nil, err
		}

		sales = append(sales, sale)
	}

	if err = rows.Err(); err != nil {
		err = fmt.Errorf("rows next error: %w", err)
		return sales, err
	}

	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		log.Println(err)
		return
	}

	for _, sale := range sales {
		log.Println(sale)
	}
}
