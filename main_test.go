package main

import (
	//"database/sql"
	//"fmt"
	//_ "github.com/go-sql-driver/mysql"
	"fmt"
	"logistics"
)

//type Role struct {
//	roleId           int
//	rolename, bz string
//}
//
//var r Role

func main() {
	fmt.Println("asdf")
	logistics.LogisticsUpdate()
	////打开mysql连接
	//db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/leshop?charset=utf8")
	//if err != nil {
	//panic(err.Error())
	//}
	//
	////连接mysql连接
	//err = db.Ping()
	//if err != nil {
	//panic(err.Error())
	//}
	//
	////建立sql语句模型
	//stmt, err := db.Prepare("INSERT role SET rolename=?,bz=?")
	//if err != nil {
	//panic(err.Error())
	//}
	//
	////建立sql模型数据
	//res, err := stmt.Exec("liangxu", "huang")
	//if err != nil {
	//panic(err.Error())
	//}
	//
	////获取最后更新的id
	//id, err := res.LastInsertId()
	//if err != nil {
	//panic(err.Error())
	//}
	//
	////输出最后更新的id
	//fmt.Println(id)
	//
	////查询数据
	//rows, err := db.Query("Select roleid, rolename from role")
	//if err != nil {
	//panic(err.Error())
	//}
	//
	////打印各列的字段名
	//fmt.Println(rows.Columns())
	//
	//roleInfo := make(map[interface{}] interface{})
	//
	//for rows.Next() {
	//    err := rows.Scan(&r.roleId, &r.rolename)
	//    if err != nil {
	//        panic(err.Error())
	//    }
	//
	//    roleInfo[r.roleId] = r
	//
	//}
	//
	//fmt.Println(roleInfo)

}
