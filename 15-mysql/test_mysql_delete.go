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

//删除
func delete() {
	s := "delete from user_tb1 where id=?"
	r, err := db.Exec(s, 1)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	i, _ := r.RowsAffected()
	fmt.Printf("i: %v\n", i)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("err : %v\n", err)
	}
	fmt.Println("连接成功~")
	delete()
}
