package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/toddlerya/luckyRoll/managedata"
)

var commands = map[string][]string{
	"windows": {"rundll32", "url.dll", "FileProtocolHandler"},
	"darwin":  {"open"},
	"linux":   {"xdg-open"},
}

func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	strRun := strings.Join(run[1:], ",")

	cmd := exec.Command(run[0], strRun, uri)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("请手动打开浏览器，输入:%s  【请不要使用IE浏览器!!!】\n", uri)
		return err
	}
	fmt.Printf("如没有打开浏览器,请手动打开浏览器,输入: %s 【请不要使用IE浏览器!!!】\n", uri)
	return nil
}

func main() {
	// 载入xlsx数据
	managedata.LoadData()

	// 启动静态文件服务
	h := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", h))

	// 展示主页, 选择班级
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

	// 根据班级条件展示大转盘页面
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
	log.Printf("Starting Server at %s port \n", port)
	//	log.Printf("请使用浏览器打开 http://127.0.0.1:%s\n", port)
	Open(fmt.Sprintf("http://127.0.0.1:%s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
