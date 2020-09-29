package model

import (
	"fmt"
)

//对应数据库字段
type Person struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

func (p *Person) List() (err error, person []Person) {
	list := []Person{} //相当于 Person[{}]
	//查询所有用户数据
	row, err := MysqlDb.Query("select * from person")
	//是否成功
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	//遍历写入
	for row.Next() {
		item := Person{}
		err = row.Scan(&item.Id, &item.Name, &item.Age, &item.Phone)
		if err != nil {
			fmt.Println("写入失败", err)
			return err, nil
		}
		list = append(list, item)
	}
	return nil, list
}
