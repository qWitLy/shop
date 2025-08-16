package urls

import (
	"html/template"
	"net/http"
	sqlR "www/sql"
	//st "www/structs"
)

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
