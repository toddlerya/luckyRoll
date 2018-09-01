package managedata

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var dbPath string = "data/data.db"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InitSqlite() {
	// 连接数据库
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		checkErr(err)
	}

	// 建表
	createStudentsBaseInfoSql := `CREATE TABLE IF NOT EXISTS students_base_info (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		stu_code VARCHAR(64),
		stu_name VARCHAR(64),
		stu_sex VARCHAR(10),
		stu_grade VARCHAR(64),
		stu_class VARCHAR(64)
	)`
	createUniqIndexStudentsBaseInfoSql := `
	CREATE UNIQUE INDEX IF NOT EXISTS stu_idx ON students_base_info(stu_code, stu_name)`

	_, err = db.Exec(createStudentsBaseInfoSql)
	checkErr(err)
	_, err = db.Exec(createUniqIndexStudentsBaseInfoSql)
	checkErr(err)
	//	if err != nil {
	//		log.Panicf("%q: %s\n", err, createStudentsBaseInfoSql)
	//		return
	//	}

	createXlsxInfoSql := `CREATE TABLE IF NOT EXISTS xlsx_info (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		xlsx_name VARCHAR(64),
		xlsx_md5 VARCHAR(10),
		xlsx_date VARCHAR(64),
		xlsx_size VARCHAR(64)
	)`
	createUniqIndexXlsxInfoSql := `
	CREATE UNIQUE INDEX IF NOT EXISTS  xlsx_idx ON xlsx_info (xlsx_md5)`

	_, err = db.Exec(createXlsxInfoSql)
	checkErr(err)
	_, err = db.Exec(createUniqIndexXlsxInfoSql)
	checkErr(err)
	//	if err != nil {
	//		log.Panicf("%q: %s\n", err, createXlsxInfoSql)
	//		return
	//	}
}

func InsertXlsxData2db(data []map[string]string) {
	db, err := sql.Open("sqlite3", dbPath)

	checkErr(err)
	defer db.Close()

	tx, err := db.Begin()
	checkErr(err)
	stmt, err := tx.Prepare("insert or replace into students_base_info(stu_code, stu_name, stu_sex, stu_class, stu_grade) values (?, ?, ?, ?, ?)")
	defer stmt.Close()

	for _, row := range data {
		_, err = stmt.Exec(row["stu_code"], row["stu_name"], row["stu_sex"], row["stu_class"], row["stu_grade"])
		checkErr(err)
	}
	tx.Commit()
}

func InsertXlsxInfo(xlsMap map[string]string) {
	db, err := sql.Open("sqlite3", dbPath)
	checkErr(err)
	defer db.Close()

	tx, err := db.Begin()
	checkErr(err)

	stmt, err := tx.Prepare("insert or replace into xlsx_info (xlsx_name, xlsx_md5, xlsx_date, xlsx_size ) values (?, ?, ?, ?)")
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(xlsMap["xlsx_name"], xlsMap["xlsx_md5"], xlsMap["xlsx_date"], xlsMap["xlsx_size"])
	checkErr(err)

	tx.Commit()
}

func QueryXlsxInfo(xlsxName string) (map[string]string, int) {
	db, err := sql.Open("sqlite3", dbPath)
	checkErr(err)
	defer db.Close()

	var xlsxMd5 string
	var xlsxDate string
	var xlsxSize string
	m := make(map[string]string)
	err = db.QueryRow("select xlsx_md5, xlsx_date, xlsx_size from xlsx_info where xlsx_name=?", xlsxName).Scan(&xlsxMd5, &xlsxDate, &xlsxSize)
	switch {
	case err == sql.ErrNoRows:
		return m, 0
	case err != nil:
		checkErr(err)
	}
	m["xlsx_name"] = xlsxName
	m["xlsx_md5"] = xlsxMd5
	m["xlsx_date"] = xlsxDate
	m["xlsx_size"] = xlsxSize
	return m, 1
}
