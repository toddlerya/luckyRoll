package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/toddlerya/luckyRoll/managedata"
)

func main() {
	// 载入xlsx数据
	managedata.LoadData()

	h := http.FileServer(http.Dir("web/static"))
	http.Handle("/web/static/", http.StripPrefix("/web/static/", h)) // 启动静态文件服务

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("web/index.html")
		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
			return
		}
		// 根据条件读取数据
		StudentsArray, _ := managedata.GetStudentsByClass("2017", "1")
		err = tmpl.Execute(w, StudentsArray)
		if err != nil {
			fmt.Fprintf(w, "Excute: %v", err)
			return
		}
	})
	log.Print("Starting Server...")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
