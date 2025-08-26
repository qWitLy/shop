package coockieFile

import (
	"log"
	"net/http"
)

func GetCoockie(w http.ResponseWriter, r *http.Request) {
	coockie, err := r.Cookie("qWitLy-coockier")
	if err != nil {
		coockie = &http.Cookie{
			Name:  "qWitLy-coockier",
			Value: "qWitLy",
			Path:  "/",
		}
	}
	http.SetCookie(w, coockie)

}

func CheckCoockie(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("qWitLy-coockier")
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/signin/", http.StatusFound)
	}
}
