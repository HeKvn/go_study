package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//1.链接数据库
	connStr := "hk01:hk01@tcp(xxx.xxx.xx.xxx:3306)/test"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//创建一张数据表
	//person:id,name,age
	// _, err1 := db.Exec("create table person(" +
	// 	"id int auto_increment primary key," +
	// 	"name varchar(12) not null," +
	// 	"age int default 1" +
	// 	");")

	// if err1 != nil {
	// 	log.Fatal(err.Error())
	// 	return
	// } else {
	// 	fmt.Println("完成")
	// }

	//插入数据到数据库表中
	_, err1 := db.Exec("insert into person(name,age)"+
		"values(?,?);", "davie", 28)

	if err1 != nil {
		log.Fatal(err1.Error())
		return
	} else {
		fmt.Println("插入成功")
	}

	//查询数据库
	rows, err := db.Query("select id,name,age from person")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
scan:
	if rows.Next() {
		person := new(Person)
		err := rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Name, person.Age, person.Id)
		goto scan
	}
}

type Person struct {
	Id   int
	Name string
	Age  int
}
