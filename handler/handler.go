package handler

import (
	"net/http"
	urls "www/handler/URLS"
)

func HandleRequest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/product/static/", http.StripPrefix("/product/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", urls.HomePage)
	http.HandleFunc("/product/", urls.ProductPage)
	http.ListenAndServe(":3030", nil)
}
