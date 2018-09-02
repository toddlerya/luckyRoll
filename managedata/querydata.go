package managedata

import (
	"math/rand"
)

type Students struct {
	Name string
}

type GradeClassInfo struct {
	Grade string
	Class string
}

func GetAllGradeClassInfo() []GradeClassInfo {
	gradeClassMapArray := QueryGradeClassInfo()
	var allDataArray []GradeClassInfo
	for _, each := range gradeClassMapArray {
		temp := GradeClassInfo{
			Grade: each["stu_grade"],
			Class: each["stu_class"],
		}
		allDataArray = append(allDataArray, temp)
	}
	return allDataArray
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
