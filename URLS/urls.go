package urls

import (
	"html/template"
	"net/http"
	sqlR "www/sql"
	//st "www/structs"
)

func HandleRequest() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":3030", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	var p = sqlR.Getproducts()
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, p)
}
