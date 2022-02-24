package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
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

func indexHtml(w http.ResponseWriter, i *http.Request) {
	var data Data
	data.Title = "我的标题"
	data.Body = "我的文本"
	t := template.New("index.html")
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, data)
}

func main() {
	//程序入口
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	//路由
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
