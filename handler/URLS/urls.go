package urls

import (
	"html/template"
	"net/http"
	"strconv"
	coockieFile "www/handler/coockie"
	sqlR "www/sql"
	st "www/structs"
)

var loginedUser st.User

/*
	 func Redirect(w http.ResponseWriter, r *http.Request, u st.User) {
		user := st.User{}
		if u == user {
			http.Redirect(w, r, "/signin/", http.StatusFound)
		} else {
			log.Println("Пользователь не пустой")
		}
	}
*/
func HomePage(w http.ResponseWriter, r *http.Request) {
	coockieFile.CheckCoockie(w, r)
	if r.Method == "POST" {
		num := r.URL.Query().Get("id")
		sqlR.AddInCart(num, strconv.Itoa(loginedUser.Id))
		http.Redirect(w, r, "/shop/", http.StatusFound)
	}
	var p = sqlR.Getproducts()
	tmpl, _ := template.ParseFiles("templates/home_page.html", "templates/footer.html", "templates/header.html")
	tmpl.Execute(w, p)
}
func ProductPage(w http.ResponseWriter, r *http.Request) {
	num := r.URL.Query().Get("id")
	var p = sqlR.GetproductById(num)
	tmpl, _ := template.ParseFiles("templates/product_page.html", "templates/footer.html", "templates/header.html")
	tmpl.Execute(w, p)
}
func SignIn(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/signin.html", "templates/footer.html", "templates/header.html")
	message := ""
	if r.Method == "POST" {
		data := st.User{
			Login:    r.FormValue("login"),
			Password: r.FormValue("password"),
		}
		if user, result := sqlR.GetUser(data); result {
			loginedUser = user
			coockieFile.GetCoockie(w, r)
			http.Redirect(w, r, "/shop/", http.StatusFound)
		} else {
			message = "Такого пользователя не существует"
		}

	}
	tmpl.Execute(w, message)
}
func SignUp(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/signup.html", "templates/footer.html", "templates/header.html")
	if r.Method == "POST" {
		data := st.User{
			Login:    r.FormValue("login"),
			Password: r.FormValue("password"),
		}
		sqlR.RegistrUser(data)
		http.Redirect(w, r, "/signin/", http.StatusFound)
	}
	tmpl.Execute(w, nil)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	coockieFile.CheckCoockie(w, r)
	tmpl, _ := template.ParseFiles("templates/profile.html", "templates/footer.html", "templates/header.html")
	tmpl.Execute(w, loginedUser)
}

func Exit(w http.ResponseWriter, r *http.Request) {
	user := st.User{}
	loginedUser = user
	http.Redirect(w, r, "/signin/", http.StatusFound)
}

func Cart(w http.ResponseWriter, r *http.Request) {
	coockieFile.CheckCoockie(w, r)
	if r.Method == "POST" {
		id := r.URL.Query().Get("id")
		sqlR.DeletProdInCart(strconv.Itoa(loginedUser.Id), id)
		http.Redirect(w, r, "/cart/", http.StatusFound)
	}
	p, _ := sqlR.ProdInCart(strconv.Itoa(loginedUser.Id))
	tmpl, _ := template.ParseFiles("templates/cart.html", "templates/footer.html", "templates/header.html")
	tmpl.Execute(w, p)
}
func Buy(w http.ResponseWriter, r *http.Request) {
	coockieFile.CheckCoockie(w, r)
	if r.Method == "POST" {
		var sum float64
		p, _ := sqlR.ProdInCart(strconv.Itoa(loginedUser.Id))
		for _, s := range p {
			sum += s.Price
		}
		if sum <= loginedUser.Money {
			loginedUser.Money = loginedUser.Money - sum
			sqlR.ChangeMoney(loginedUser.Money, strconv.Itoa(loginedUser.Id))
			for _, prod := range p {
				sqlR.Buy(true, strconv.Itoa(loginedUser.Id), strconv.Itoa(prod.Id))
				sqlR.ChangeCountProd(prod.Count-1, prod.Id)
			}
			http.Redirect(w, r, "/cart/", http.StatusFound)
		} else {
			http.Redirect(w, r, "/cart/", http.StatusFound)
		}
	}
}

func Replenishment(w http.ResponseWriter, r *http.Request) {
	coockieFile.CheckCoockie(w, r)
	if r.Method == "POST" {
		money := r.FormValue("money")
		m, _ := strconv.ParseFloat(money, 64)
		loginedUser.Money += m
		sqlR.ChangeMoney(loginedUser.Money, strconv.Itoa(loginedUser.Id))
		http.Redirect(w, r, "/shop/", http.StatusFound)
	}
	tmpl, _ := template.ParseFiles("templates/replenushment.html", "templates/footer.html", "templates/header.html")
	tmpl.Execute(w, nil)
}
