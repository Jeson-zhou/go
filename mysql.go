package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

type user struct {
	id int
	age int
	name string
}

// 预处理插入
func prepareInsertDemo() {
	sqlStr := "insert into user(name, age) values(?, ?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err: %v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("zhoulei", 26)
	if err != nil {
		fmt.Printf("Insert failed, err: %v\n", err)
		return
	}
	fmt.Printf("Insert success!!!")
}

// 预处理查询
func prepareQueryDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err: %v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
		return
	}
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("id: %d, name: %s, age: %d\n", u.id, u.name, u.age)
	}
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("delete from user success, affected rows: %d\n", n)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age = ? where id = ?"
	ret, err := db.Exec(sqlStr, 25, 1)
	if err != nil {
		fmt.Printf("Update DB failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("Update DB Success, affected rows: %d\n", n)
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values(?, ?)"
	ret, err := db.Exec(sqlStr, "zhou", 23)
	if err != nil {
		fmt.Printf("Insert failed, err: %v\n", err)
		return
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("Get lastInsert ID failed, err: %v\n", err)
		return
	}
	fmt.Printf("insert scuuess, the id is %d\n", theId)
}

// 多行查询
func QueryMultiRowDemo() {
	sqlStr := "SELECT id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("Query failed, err: %v\n", err)
		return
	}
	// 非常重要，关闭rows释放持有的数据库连接
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("Scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("id: %d, name: %s, age: %d\n", u.id, u.name, u.age)
	}
}

// 单行查询
func QueryRow() {
	sqlStr := "SELECT id, name, age from user where id = ?"
	var u user
	// 非常重要，确保QueryRow之后调用Scan方法，否则持有的数据库连接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("Scan failed, err: %v\n", err)
		return
	}
	fmt.Printf("id: %d name: %s age: %d\n", u.id, u.name, u.age)
}


// 使用MySQL驱动
func initDB() (err error) {
	// DSN：Data Source Name
	dsn := "root:@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Open DB failed", err)
		return err
	}
	// 尝试与DB建了连接(校验DSN是否正确)
	err = db.Ping()
	if err != nil {
		fmt.Println("Username or Password is error!!!")
		return err
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(2)
	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Init DB failed, err: %v\n", err)
		return
	}
	QueryRow()
	QueryMultiRowDemo()
	insertRowDemo()
	QueryMultiRowDemo()
	updateRowDemo()
	QueryMultiRowDemo()
	deleteRowDemo()
	QueryMultiRowDemo()
	prepareQueryDemo()
	prepareInsertDemo()
	QueryMultiRowDemo()
}
