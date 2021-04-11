package main

import (
	"database/sql/driver"
	"errors"
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

type User struct {
	Name string
	Age  int
}

// 使用sqlx.in实现批量插入，前提需要我们的结构体实现driver.Value接口
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

func BatchInsertUsers2(users []interface{}) error {
	sqlStr := "insert into user(name, age) values(?), (?), (?)"
	queryStr, args, _ := sqlx.In(sqlStr, users...)
	fmt.Println(queryStr)
	fmt.Println(args)
	_, err := db.Exec(queryStr, args...)
	return err
}

// sqlx 事务操作
func transactionDemo2() (err error) {
	// 开启事务
	tx, err := db.Beginx()
	if err != nil {
		fmt.Printf("Begin trans failed, err: %v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Println("31: Rollback...")
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				fmt.Printf("Rollback failed, err: %v\n", err)
				panic(err)
			}
			fmt.Println("Commit success...")
		}
	}()
	sqlStr1 := "Update user set age = 3 0 where id = ?"
	ret, err := db.Exec(sqlStr1, 1)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}

	sqlStr2 := "Update user set age = 50 where id = ?"
	ret, err = db.Exec(sqlStr2, 2)
	if err != nil {
		return err
	}
	n, err = ret.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr2 failed")
	}
	return nil
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
	transactionDemo2()
	fmt.Println("--------")
	u1 := User{Name: "x", Age: 1}
	u2 := User{Name: "xx", Age: 2}
	u3 := User{Name: "xxx", Age: 3}
	user2 := []interface{}{u1, u2, u3}
	BatchInsertUsers2(user2)
}
