package managedata

import (
	"fmt"
	"math/rand"
)

type Students struct {
	Name string
}

type ClassInfo struct {
	Class string
}

type GradeInfo struct {
	Grade      string
	ClassChild []ClassInfo
}

func GetAllGradeClassInfo() {
	gradeClassMapArray := QueryGradeClassInfo()
	fmt.Println(gradeClassMapArray)
}

func GetStudentsByClass(grade, class string) ([]Students, int) {
	//func GetStudentsByClass(grade, class string) ([]byte, int) {
	studentMaps, count := QueryStudentsInfo(grade, class)
	// 洗牌
	rand.Shuffle(len(studentMaps), func(i, j int) { studentMaps[i], studentMaps[j] = studentMaps[j], studentMaps[i] })
	var transData []Students
	for _, each := range studentMaps {
		temp := Students{
			Name: each["stu_name"],
		}
		transData = append(transData, temp)
	}
	return transData, count
}
