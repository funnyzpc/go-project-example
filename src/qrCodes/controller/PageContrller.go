package PageContrller

import (
	"html/template"
	"log"
	"net/http"
)

func IndexAction(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("template/page/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(writer, nil)
}
