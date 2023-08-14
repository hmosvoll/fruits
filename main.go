package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")

	type Product struct {
		Name  string
		Emoji string
		Price int
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("index.html"))

		products := []Product{
			{Name: "Apple", Emoji: "🍎", Price: 10},
			{Name: "Banana", Emoji: "🍌", Price: 5},
			{Name: "Orange", Emoji: "🍊", Price: 15},
			{Name: "Pear", Emoji: "🍐", Price: 20},
			{Name: "Pineapple", Emoji: "🍍", Price: 25},
			{Name: "Strawberry", Emoji: "🍓", Price: 30},
			{Name: "Watermelon", Emoji: "🍉", Price: 35},
			{Name: "Kiwi", Emoji: "🥝", Price: 40},
			{Name: "Avocado", Emoji: "🥑", Price: 45},
			{Name: "Tomato", Emoji: "🍅", Price: 50},
			{Name: "Eggplant", Emoji: "🍆", Price: 55},
			{Name: "Potato", Emoji: "🥔", Price: 60},
			{Name: "Carrot", Emoji: "🥕", Price: 65},
			{Name: "Corn", Emoji: "🌽", Price: 70},
			{Name: "Hot Pepper", Emoji: "🌶", Price: 75},
			{Name: "Cucumber", Emoji: "🥒", Price: 80},
			{Name: "Mushroom", Emoji: "🍄", Price: 85},
			{Name: "Peanuts", Emoji: "🥜", Price: 90},
			{Name: "Chestnut", Emoji: "🌰", Price: 95},
		}

		template.Execute(w, products)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
