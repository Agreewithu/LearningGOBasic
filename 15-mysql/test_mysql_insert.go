package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:agreewithu@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

//插入
func insert() {
	s := "insert into user_tb1 (username,password) values(?,?)" //插入表的sql语句
	r, err := db.Exec(s, "zhangsan", "zs123")
	if err != nil {
		fmt.Printf("err : %v\n", err)
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("i: %v\n", i)
	}
}

func insertwithpara(username, password string) {
	s := "insert into user_tb1 (username,password) values(?,?)" //插入表的sql语句
	r, err := db.Exec(s, username, password)
	if err != nil {
		fmt.Printf("err : %v\n", err)
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("i: %v\n", i)
	}
}

func main() {

	err := initDB()
	if err != nil {
		fmt.Printf("err : %v\n", err)
	}
	fmt.Println("连接成功~")
	insert()
	insertwithpara("lisi", "ls123")
}
