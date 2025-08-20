package urls

import (
	"html/template"
	"net/http"
	sqlR "www/sql"
	st "www/structs"
)

var loginedUser st.User

func Redirect(w http.ResponseWriter, r *http.Request, u st.User) {
	user := st.User{}
	if u == user {
		http.Redirect(w, r, "/signin/", http.StatusFound)
	}
}
func HomePage(w http.ResponseWriter, r *http.Request) {
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
		http.Redirect(w, r, "/login/", http.StatusFound)
	}
	tmpl.Execute(w, nil)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/profile.html", "templates/footer.html", "templates/header.html")
	Redirect(w, r, loginedUser)
	tmpl.Execute(w, loginedUser)
}
