package handler

import (
	"net/http"
	urls "www/handler/URLS"
)

func HandleRequest() {
	http.Handle("/shop/static/", http.StripPrefix("/shop/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/product/static/", http.StripPrefix("/product/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/signup/static/", http.StripPrefix("/signup/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/signin/static/", http.StripPrefix("/signin/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/shop/", urls.HomePage)
	http.HandleFunc("/profile/", urls.Profile)
	http.HandleFunc("/product/", urls.ProductPage)
	http.HandleFunc("/signup/", urls.SignUp)
	http.HandleFunc("/signin/", urls.SignIn)
	http.ListenAndServe(":3030", nil)
}
