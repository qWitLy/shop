package main

import (
	_ "database/sql"
	_ "fmt"
	hand "www/handler"
	_ "www/structs"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//handleRequest()
	hand.HandleRequest()

}

/* db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:8080)/user")
if err != nil {
	fmt.Print("Ошибка")
	panic(err)
}

defer db.Close() */
/* fmt.Print("Подключение установлено")
insert, err3 := db.Query("INSERT INTO `user`.`user` (`name`, `age`) VALUES ('Alex',25)")
if err3 != nil {
	fmt.Print("hwllo")
	panic(err3)
}
defer insert.Close() */

//выборка данных
/* res, err := db.Query("SELECT `name`, `age` FROM `user`.`user`")
if err != nil {
	fmt.Print("Ошибка")
	panic(err)
}
for res.Next() {
	var user structer.User
	err = res.Scan(&user.Name, &user.Age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User: %s with age: %d \n", user.Name, user.Age)
} */
