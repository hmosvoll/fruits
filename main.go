package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("hello world")

	type Product struct {
		Name string
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("index.html"))

		products := []Product{{Name: "Apple"}, {Name: "Banana"}, {Name: "Orange"}}

		template.Execute(w, products)
	})

	http.HandleFunc("/fruits", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)

		fruit := r.PostFormValue("name")

		template := template.Must(template.ParseFiles("index.html"))
		template.ExecuteTemplate(w, "fruit", Product{Name: fruit})
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
