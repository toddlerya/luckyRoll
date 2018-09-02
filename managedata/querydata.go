package managedata

import (
	"math/rand"
	"time"
)

type Students struct {
	Name string
}

func GetStudentsByClass(grade, class string) ([]Students, int) {
	//func GetStudentsByClass(grade, class string) ([]byte, int) {
	studentMaps, count := QueryStudentsInfo(grade, class)
	shuffledStudentMaps := shuffle(studentMaps)
	//	data, err := json.Marshal(shuffledStudentMaps)
	//	if err != nil {
	//		log.Fatalf("json.Marshal err: %v", err)
	//	}
	//	return data, count
	var transData []Students
	for _, each := range shuffledStudentMaps {
		//		fmt.Println(each)
		temp := Students{
			Name: each["stu_name"],
		}
		transData = append(transData, temp)
	}
	return transData, count
}

// 随机洗牌
func shuffle(rawData []map[string]string) []map[string]string {
	rand.Seed(time.Now().UnixNano())
	// TODO 需要学习下Shuffle的用法
	rand.Shuffle(len(rawData), func(i, j int) { rawData[i], rawData[j] = rawData[j], rawData[i] })
	return rawData
}
