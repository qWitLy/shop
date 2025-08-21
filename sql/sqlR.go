package sqlR

import (
	"database/sql"
	"log"
	st "www/structs"

	_ "github.com/go-sql-driver/mysql"
)

const connectionString = "root:admin@tcp(:8080)/shop"
const nameDb = "mysql"

// выборка данных
func Getproducts() []st.Product {
	query := "SELECT product.id, `name`, `description`, `price`, `count`, `link` FROM `shop`.`product`INNER JOIN `shop`.`links` ON product.id = links.id"
	db, err := sql.Open(nameDb, connectionString)
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

func GetproductById(id string) st.Product {
	query := "SELECT product.id, `name`, `description`, `price`, `count`, `link` FROM `shop`.`product`INNER JOIN `shop`.`links` ON product.id = links.id_prod where product.id =?"
	db, err := sql.Open(nameDb, connectionString)
	if err != nil {
		log.Fatal("не удалось подключиться к базе данных: ", err.Error())
	}
	log.Println("Подключение успешное")
	res, err := db.Query(query, id)
	if err != nil {
		log.Fatal("Не удалось выполнить запрос: ", err.Error())
	}
	log.Println("Запрос на выборку данных выполнился")
	var p st.Product
	for res.Next() {
		err = res.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Count, &p.Link)
		if err != nil {
			log.Fatal("Не удалось получить данные: ", err.Error())
		}
	}
	defer func() {
		db.Close()
		log.Println("Отключился от бд")
	}()
	defer res.Close()
	return p
}

func GetUser(u st.User) (st.User, bool) {
	query := "SELECT * FROM `shop`.`user` where `user`.`name` = ? and `user`.`password` = ?"
	db, err := sql.Open(nameDb, connectionString)
	if err != nil {
		log.Fatal("не удалось подключиться к базе данных: ", err.Error())
	}
	log.Println("Подключение успешное")
	res := db.QueryRow(query, u.Login, u.Password)
	var user st.User
	check := st.User{}
	err = res.Scan(&user.Id, &user.Login, &user.Password, &user.Money)
	if err != nil {
		user = check
		return user, false
	}
	defer func() {
		db.Close()
		log.Println("Отключился от бд")
	}()
	return user, true
}

func RegistrUser(u st.User) {
	query := "INSERT INTO `shop`.`user` (`name`, `money`, `password`) VALUES (?, ?, ?)"
	db, err := sql.Open(nameDb, connectionString)
	if err != nil {
		log.Fatal("не удалось подключиться к базе данных: ", err.Error())
	}
	log.Println("Подключение успешное")
	res, err := db.Query(query, u.Login, 0, u.Password)
	if err != nil {
		log.Fatal("Не удалось выполнить запрос: ", err.Error())
	}
	defer func() {
		db.Close()
		log.Println("пользователь добавлен")
		log.Println("Отключился от бд")
	}()
	defer res.Close()
}

func AddInCart(prodId, userId string) bool {
	query := "INSERT INTO `shop`.`produs` (`product_id`,`user_id`) VALUES (?,?)"
	db, err := sql.Open(nameDb, connectionString)
	if err != nil {
		log.Fatal("не удалось подключиться к базе данных: ", err.Error())
	}
	log.Println("Подключение успешное")
	res, err := db.Query(query, prodId, userId)
	if err != nil {
		log.Fatal("Не удалось выполнить запрос: ", err.Error())
	}
	log.Println("Запрос на выборку данных выполнился")
	defer func() {
		db.Close()
		log.Println("Отключился от бд")
	}()
	defer res.Close()
	return true
}

func ProdInCart(userId string) ([]st.Product, bool) {
	query := "SELECT product.id, `name`, `description`, `price`, `count`, `link` FROM `shop`.`product`INNER JOIN `shop`.`links` ON product.id = links.id_prod INNER JOIN `shop`.`produs` ON product.id = produs.product_id where `user_id` = ?"
	db, err := sql.Open(nameDb, connectionString)
	if err != nil {
		log.Fatal("не удалось подключиться к базе данных: ", err.Error())
	}
	log.Println("Подключение успешное")
	res, err := db.Query(query, userId)
	if err != nil {
		log.Fatal("Не удалось выполнить запрос: ", err.Error())
	}
	log.Println("Запрос на выборку данных выполнился")
	var products []st.Product
	for res.Next() {
		var p st.Product
		err = res.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Count, &p.Link)
		if err != nil {
			return products, false
		}
		products = append(products, p)
	}
	defer func() {
		db.Close()
		log.Println("Отключился от бд")
	}()
	defer res.Close()
	return products, true
}
