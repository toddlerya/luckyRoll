package managedata

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func calcMd5(fileName string) string {
	fileObj, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	md5hash := md5.New()
	io.Copy(md5hash, fileObj)
	fileMd5 := hex.EncodeToString(md5hash.Sum(nil))
	return fileMd5
}

// 获取指定目录下所有的xlsx文件名
func getAllXlsx(dirpath string) ([]string, error) {
	var files []string
	_, err := os.Stat(dirpath)
	if err != nil {
		pathErr := fmt.Sprintf("%s目录不存在!\n", dirpath)
		return files, errors.New(pathErr)
	}
	xlsxPath := fmt.Sprintf("%s/*.xlsx", dirpath)
	files, err = filepath.Glob(xlsxPath)
	return files, err
}

// 读取一个学生名单
func readXlsx(xlsxPath string) ([][]string, error) {
	var xlsxArray [][]string
	xlsx, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		fmt.Println(err)
		return xlsxArray, err
	}
	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		if rows[0][0] != "学号" && rows[0][1] != "姓名" && rows[0][2] != "性别" {
			xlsxTitleErr := fmt.Sprintf("%s文件Sheet1的表头应该为: [学号 姓名 性别]")
			return xlsxArray, errors.New(xlsxTitleErr)
		}
		if row[0] == "学号" && row[1] == "姓名" && row[2] == "性别" {
			continue
		}
		xlsxArray = append(xlsxArray, row)
	}
	return xlsxArray, nil
}

// 解析所有数据，进行预处理, 应该是根据每个文件的md5判断, 若文件发生变更则根据此文件更新数据库
func parseXlsxData(dataPath string) ([]map[string]string, error) {
	var datas []map[string]string
	// 获取所有的文件
	allFiles, err := getAllXlsx(dataPath)
	if err != nil {
		errMsg := fmt.Sprintf("从%s目录下获取所有xlsx文件失败: %s", dataPath, err)
		return datas, errors.New(errMsg)
	}

	// 依次读取所有文件
	for _, filePathName := range allFiles {
		// 获取文件去除路径后的名字
		xlsxName := filepath.Base(filePathName)
		xlsxName = strings.TrimSpace(xlsxName)
		// 校验文件是否有改变, 若无改变则不在重新读取
		fileInfoMap, num := QueryXlsxInfo(xlsxName)
		fmt.Println("===>", num)
		fmt.Println(fileInfoMap)
		if num == 1 {
			oldFileMd5 := fileInfoMap["xlsx_md5"]
			newFileMd5 := calcMd5(filePathName)
			if newFileMd5 == oldFileMd5 {
				fmt.Printf("%s文件没有变化, 跳过, md5: %s", filePathName, oldFileMd5)
				continue
			}
		} else {
			// 数据库没有记录此文件, 获取文件信息入库
			m := make(map[string]string)
			xlsxMd5 := calcMd5(xlsxName)
			xlsxBaseInfo, err := os.Stat(filePathName)
			if err != nil {
				log.Fatal(err)
			}
			xlsxDate := xlsxBaseInfo.ModTime()
			xlsx_unix_sec, _ := time.ParseInLocation("2006-01-02 15:04:05", xlsxDate, time.Local)
			xlsxSize := xlsxBaseInfo.Size()
			fmt.Println(xlsxDate)
			fmt.Println(xlsx_unix_sec)
			fmt.Println(xlsxSize)
			m["xlsx_name"] = xlsxName
			m["xlsx_md5"] = xlsxMd5
			m["xlsx_date"] = fmt.Sprintf("%s", xlsx_unix_sec)
			m["xlsx_size"] = fmt.Sprintf("%s", xlsxSize)
			InsertXlsxInfo(m)
		}
		xlsxArray, err := readXlsx(filePathName)
		if err != nil {
			log.Printf("读取%s错误: %s", filePathName, err)
		}
		// 获取班级信息
		xlsxSplitArray := strings.Split(xlsxName, ".xlsx")[0]
		tmpArray := strings.Split(xlsxSplitArray, "级")
		grade := tmpArray[0]
		class := strings.Split(tmpArray[1], "班")[0]
		_, err = strconv.Atoi(grade)
		_, err = strconv.Atoi(class)
		if err != nil {
			log.Fatalf("请检查xlsx文件名是否为XXXX级Y班, XXXX为年份(比如2018代表2018级), Y为班级序号(比如1代表一班)")
		}
		for _, item := range xlsxArray {
			m := make(map[string]string)
			m["stu_code"] = item[0]
			m["stu_name"] = item[1]
			m["stu_sex"] = item[2]
			m["stu_grade"] = grade
			m["stu_class"] = class
			datas = append(datas, m)
		}
	}
	return datas, nil
}

func LoadData() {
	InitSqlite()
	xlsxDatas, err := parseXlsxData("data")
	if err != nil {
		log.Fatal(err)
	}
	InsertXlsxData2db(xlsxDatas)
}
