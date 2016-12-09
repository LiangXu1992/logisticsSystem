package databases

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//定义数据库连接结构
type DatabaseConfig struct {
	driver   string
	user     string
	passWord string
	dbName   string
	host     string
	port     string
}

//定义mysql配置变量
var mysql = DatabaseConfig{
	driver:   "mysql",
	user:     "root",
	passWord: "root",
	dbName:   "leshop",
	host:     "127.0.0.1",
	port:     ":3306",
}

//初始化连接数据库
func InitDb(c *gin.Context) {
	var err error
	c.Db, err = sql.Open(mysql.driver, mysql.user+":"+mysql.passWord+"@tcp("+mysql.host+mysql.port+")/"+mysql.dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	err = c.Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDb(c *gin.Context) {
	c.Db.Close()
}
