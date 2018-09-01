package managedata

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var dbPath string = "data/data.db"

func InitSqlite() {
	// 连接数据库
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 建表
	createStudentsBaseInfoSql := `CREATE TABLE IF NOT EXISTS students_base_info (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		stu_code VARCHAR(64),
		stu_name VARCHAR(64),
		stu_sex VARCHAR(10),
		stu_grade VARCHAR(64),
		stu_class VARCHAR(64)
	)`

	_, err = db.Exec(createStudentsBaseInfoSql)
	if err != nil {
		log.Panicf("%q: %s\n", err, createStudentsBaseInfoSql)
		return
	}

	createXlsxInfoSql := `CREATE TABLE IF NOT EXISTS xlsx_info (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		xlsx_name VARCHAR(64),
		xlsx_md5 VARCHAR(10),
		xlsx_date VARCHAR(64),
		xlsx_size VARCHAR(64)
	)`

	_, err = db.Exec(createXlsxInfoSql)
	if err != nil {
		log.Panicf("%q: %s\n", err, createXlsxInfoSql)
		return
	}
}

func InsertData2db(data []map[string]string) {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into students_base_info(stu_code, stu_name, stu_sex, stu_class, stu_grade) values(?, ?, ?, ?, ?)")
	defer stmt.Close()

	for _, row := range data {
		_, err = stmt.Exec(row["stu_code"], row["stu_name"], row["stu_sex"], row["stu_class"], row["stu_grade"])
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}
