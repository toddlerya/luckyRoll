package main

import (
	"fmt"

	"github.com/toddlerya/luckyRoll/managedata"
)

func main() {
	managedata.LoadData()
	names, number := managedata.GetStudentsByClass("2017", "1")
	fmt.Println(names)
	fmt.Println(number)
}
