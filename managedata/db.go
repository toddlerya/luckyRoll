package managedata

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var dbPath string = "data/data.db"

func initSqlite() {
	// 连接数据库
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 建表
	createSql := `CREATE TABLE IF NOT EXISTS students_base_info (
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		stu_code VARCHAR(64),
		stu_name VARCHAR(64),
		stu_sex VARCHAR(10),
		stu_grade VARCHAR(64),
		stu_class VARCHAR(64),
	)`

	_, err = db.Exec(createSql)
	if err != nil {
		log.Panicf("%q: %s\n", err, createSql)
		return
	}
}

//func loadSqlite(data []map[string]string) {
//	db, err := sql.Open("sqlite3", dbPath)

//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()

//	tx, err := db.Begin()
//	if err != nil {
//		log.Fatal(err)
//	}
//	stmt, err := tx.Prepare("insert into students_base_info(stu_code, stu_name, stu_sex, stu_class, stu_grade) values(?, ?, ?, ?, ?)")
//	defer stmt.Close()

//	for _, row := range data {
//		_, err = stmt.Exec(row[row["stu_code"], "stu_name"], row["stu_sex"], row["stu_class"], row["stu_grade"])
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//	tx.Commit()
//}
