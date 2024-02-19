package main

import (
	"database/sql"
	"fmt"

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
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale
	// напишите код здесь
	// Инициализируем пустой слайс для дальнейшего заполнения объектами Sale
	sales = make([]Sale, 0)
	// Подключаемся к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Printf("selectSale sql.Open error %v\n", err)
		return nil, err
	}
	defer db.Close()

	// Создаем и выполняем запрос
	// Также можно использовать функцию  QueryContext, как альтернативный вариант
	rows, err := db.Query("SELECT product, volume, date FROM sales WHERE client = :id", sql.Named("id", client))
	if err != nil {
		fmt.Printf("selectSale db.Query error %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// Считываем результат запроса
	for rows.Next() {
		// Создаем новый объект структуры Sale
		sale := Sale{}
		// Построчно считываем результаты запроса и записываем данные в соответсвующие поля ранее созданной структуры
		err := rows.Scan(&sale.Product, &sale.Volume, &sale.Date)
		if err != nil {
			fmt.Printf("selectSale rows.Scan error %v\n", err)
			return nil, err
		}
		sales = append(sales, sale)
	}
	// Добавляем проверку rows.Err(), она проверяет, что цикл, в котором вызывается  rows.Next() завершился штатно,
	// когда закончились все записи, а не в результате какой-то ошибки
	if err := rows.Err(); err != nil {
		fmt.Printf("selectSale for rows.Next error %v \n", err)
		return nil, err
	}
	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
