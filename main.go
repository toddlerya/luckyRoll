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

	// 启动静态文件服务
	h := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", h))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("web/index.html")
		if err != nil {
			fmt.Fprintf(w, "index.html ParseFiles: %v", err)
			return
		}
		// 根据条件读取数据
		gradeClassArray := managedata.GetAllGradeClassInfo()
		err = tmpl.Execute(w, gradeClassArray)
		if err != nil {
			fmt.Fprintf(w, "Excute index.html: %v", err)
			return
		}
	})

	http.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("web/roll.html")
		if err != nil {
			fmt.Fprintf(w, "roll.html ParseFiles: %v", err)
			return
		}
		// 根据条件读取数据
		grade := r.FormValue("grade")
		class := r.FormValue("class")
		StudentsArray, _ := managedata.GetStudentsByClass(grade, class)
		err = tmpl.Execute(w, StudentsArray)
		if err != nil {
			fmt.Fprintf(w, "Excute roll.html: %v", err)
			return
		}
	})
	port := "9000"
	log.Printf("Starting Server at 0.0.0.0:%s \n请使用浏览器打开 http://127.0.0.1:%s", port, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
