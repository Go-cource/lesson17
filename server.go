package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
GET /books - получить все книги
GET /books/{id} - получить одну книгу с идентификатором

POST /books - отправить (создать) новую книгу
*/
func BooksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Books")
}
func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Add Book")
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", BooksHandler)
	r.HandleFunc("/books/{id}", AddBookHandler)
	fmt.Println("СТАРТУЕМ! Я сказал стартуем...")
	http.ListenAndServe("127.0.0.1:8080", r)
}
