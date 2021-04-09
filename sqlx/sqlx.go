package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	Id   int
	Name string
	Age  int
}

// 更新数据
func updateRowDemo() {
	var sqlStr = "update user set age = ? where id = ?"
	ret, err := db.Exec(sqlStr, 88, 2)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("Get rowaffected faile, err: %v\n", err)
		return
	}
	fmt.Printf("Update DB success, RowsAfffected: %d\n", n)
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values(?, ?)"
	ret, err := db.Exec(sqlStr, "zhoulei", 11)
	if err != nil {
		fmt.Printf("Insert data failed, err: %v\n", err)
		return
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("Get last insert id failed, err: %v\n", err)
		return
	}
	fmt.Printf("Insert success, the id is: %d\n", theId)
}

// DB多行查询
func queryMultiRowDemo() {
	var sqlStr = "select name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("Query DB Failed, err: %v\n", err)
		return
	}
	fmt.Printf("users: %#v\n", users)
}

// DB 单行查询
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var u user
	err := db.Get(&u, sqlStr, 0)
	if err != nil {
		fmt.Printf("db.Get failed, err: %v\n", err)
		return
	}
	fmt.Printf("id: %d, name: %s, age: %d\n", u.Id, u.Name, u.Age)
}

// DB初始化
func initDB() (err error) {
	dsn := "root:@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("Connect mysql failed, err: %v\n", err)
		return err
	}
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(5)
	return nil
}

func main() {
	initDB()
	queryRowDemo()
	fmt.Println("======================")
	queryMultiRowDemo()
	insertRowDemo()
	updateRowDemo()
	queryMultiRowDemo()
}
