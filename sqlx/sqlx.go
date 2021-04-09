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

//DB.NamedQuery
func namedQuery() {
	sqlStr := "select * from user where name = :name"
	//	使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{
		"name": "zhou",
	})
	if err != nil {
		fmt.Printf("NamedQuery failed, err: %v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("rows.StructScan failed, err: %v\n", err)
			return
		}
		fmt.Printf("user: %#v\n", u)
	}

	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
	u := user{
		Name: "zhou",
	}
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery Failed, err: %v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("StructScan failed, err: %v\n", err)
			return
		}
		fmt.Printf("user: %#v\n", u)
	}
}

// DB.NamedExec方法用来绑定SQL语句与结构体或map中的同名字段。
func insertUserDemo() (err error) {
	sqlStr := "insert into user(name, age) values(:name, :age)"
	ret, err := db.NamedExec(sqlStr, map[string]interface{}{
		"name": "anlina",
		"age":  23,
	})
	if err != nil {
		fmt.Printf("Insert user failed, err: %v\n", err)
		return err
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("Get lasteinsertid failed, err: %v\n", err)
		return err
	}
	fmt.Printf("Insert user success, last insert id is: %d\n\n", theId)
	return nil
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 4)
	if err != nil {
		fmt.Printf("Delete failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowAffected Failed, err: %v\n", err)
		return
	}
	fmt.Printf("Delete success, affedted rows: %d\n", n)
}

// 更新数据
func updateRowDemo() {
	var sqlStr = "update user set age = ? where id = ?"
	ret, err := db.Exec(sqlStr, 88, 2)
	if err != nil {
		fmt.Printf("Update failed, err: %v\n", err)
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
	deleteRowDemo()
	insertUserDemo()
	namedQuery()
}
