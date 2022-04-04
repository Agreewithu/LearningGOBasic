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

//定义一个结构体，查询过来
type User struct {
	id       int
	username string
	password string
}

//查询单行
func queryOneRow() {
	s := "select * from user_tb1 where id = ?"
	var u User
	err := db.QueryRow(s, 1).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Printf("err : %v\n", err)
	}
	fmt.Printf("u: %v\n", u)
}

//查询多行
func queryManyRow() {
	s := "select * from user_tb1"
	r, err := db.Query(s)
	var u User
	defer r.Close() //r返回来是一个迭代器
	if err != nil {
		fmt.Printf("err : %v\n", err)
	}
	for r.Next() {
		r.Scan(&u.id, &u.username, &u.password)
		fmt.Printf("u: %v\n", u)
	}
}

func main() {
	//先进行初始化，相当于把数据库连接拿到了
	err := initDB()
	if err != nil {
		fmt.Printf("err : %v\n", err)
	}
	fmt.Println("连接成功~")

	queryOneRow()
	queryManyRow()
}
