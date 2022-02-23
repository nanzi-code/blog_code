package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Data struct {
	Title string `json:title`
	Body  string `json:body`
}

func index(w http.ResponseWriter, i *http.Request) {
	var data Data
	data.Title = "我的标题"
	data.Body = "我的文本"
	jstr, _ := json.Marshal(data)
	w.Write(jstr)
}

func main() {
	//程序入口
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
