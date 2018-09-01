package managedata

import (
	"math/rand"
	"time"
)

func GetStudentsByClass(grade, class string) ([]map[string]string, int) {
	studentMaps, count := QueryStudentsInfo(grade, class)
	return shuffle(studentMaps), count
}

// 随机洗牌
func shuffle(rawData []map[string]string) []map[string]string {
	rand.Seed(time.Now().UnixNano())
	// TODO 需要学习下Shuffle的用法
	rand.Shuffle(len(rawData), func(i, j int) { rawData[i], rawData[j] = rawData[j], rawData[i] })
	return rawData
}
