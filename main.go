package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/toddlerya/luckyRoll/managedata"
)

func main() {
	managedata.LoadData()
	names, number := managedata.GetStudentsByClass("2017", "1")
	data, err := json.Marshal(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
	fmt.Println(number)
}
