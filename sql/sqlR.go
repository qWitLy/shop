package sqlR

import (
	"database/sql"
	"log"
	st "www/structs"

	_ "github.com/go-sql-driver/mysql"
)

// выборка данных
func Getproducts() []st.Product {
	query := "SELECT product.id, `name`, `description`, `price`, `count`, `link` FROM `shop`.`product`INNER JOIN `shop`.`links` ON product.id = links.id"
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:8080)/shop")
	if err != nil {
		log.Fatal("не удалось подключиться к базе данных: ", err.Error())
	}
	log.Println("Подключение успешное")
	res, err := db.Query(query)
	if err != nil {
		log.Fatal("Не удалось выполнить запрос: ", err.Error())
	}
	log.Println("Запрос на выборку данных выполнился")
	var products []st.Product
	for res.Next() {
		var p st.Product
		err = res.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Count, &p.Link)
		if err != nil {
			log.Fatal("Не удалось получить данные: ", err.Error())
		}
		products = append(products, p)
	}
	defer func() {
		db.Close()
		log.Println("Отключился от бд")
	}()
	defer res.Close()
	return products
}
