package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	data := []byte(`{"ID": 3, "Title": "Евгений Онегин", "Author": "Пушкин"}`) //создаем JSON и переводим его в слайс байт - тк только байты можно передать

	r := bytes.NewReader(data)                                                   //из слайса байт делаем переменную типа Reader (она нужна в методе Pots пакета http)
	resp, err := http.Post("http://127.0.0.1:8080/books", "application/json", r) //вызываем метод POST пакета http, передаем URL, тип передаваемых данных (JSON - почитать про ЗАГОЛОВКИ HTTP) и сами данные
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Сервер мне ответил! Он сказал: ", resp.Body)
}
