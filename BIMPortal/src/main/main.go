package main

import (
"net/http"
	"log"
	"html/template"
)


func IndexHandler(w http.ResponseWriter, r *http.Request){

    t, err := template.ParseFiles("resources/html/index.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}

func main() {
    http.Handle("/css/", http.FileServer(http.Dir("resources/bootstrap-4.3.1")))
    http.Handle("/js/", http.FileServer(http.Dir("resources/bootstrap-4.3.1")))

    http.HandleFunc("/",IndexHandler)
    http.ListenAndServe(":8181", nil)
}