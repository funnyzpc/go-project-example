package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", index)
	log.Println("请访问:", "http://127.0.0.1:2222")
	error := http.ListenAndServe(":2222", nil)
	if error != nil {
		log.Fatal("错误：", error)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		return
	}
	fmt.Printf("[%s|%s] -> http://%s%s \n", r.Method, r.Proto, r.Host, r.RequestURI)
	dateTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintln(w, "\n")
	fmt.Fprintln(w, "\tHello youth !")
	fmt.Fprintln(w, "\tRequest by : ", dateTime)
}
