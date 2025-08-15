package urls

import (
	"html/template"
	"net/http"
	sqlR "www/sql"
	//st "www/structs"
)

func HandleRequest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":3030", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	var p = sqlR.Getproducts()
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, p)
}
