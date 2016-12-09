package models

import (
	//"common"
	//"time"
	//"strconv"

	"common"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
)

//物流员结构体
type Shop_logistics struct {
	Logistics_id   int `gorm:"primary_key"`
	Logistics_name string
	Token          string
	Client_type    string
	Password       string
	Login_time     int
	App_token      string
	Db             *sql.DB
}

//物流单结构体
type Shop_v_dispatch struct {
	Dispatch_id int
	Order_sn    string
	Db          *sql.DB
}

func (l *Shop_logistics) GetLogisticsInfo() {
	//查找物流员的信息

	err := l.Db.QueryRow("SELECT logistics_id, token FROM shop_logistics where logistics_name = ? AND logistics_pwd = ?", l.Logistics_name, common.STMd5(l.Password)).Scan(&l.Logistics_id, &l.Token)
	if err != nil {
		log.Fatalf("Query logistics error %v", err)
	}
	//更新token，apptoken
	l.UpdateToken()

}

func (l *Shop_logistics) UpdateToken() {
	//更新物流员登录token
	randomInt := strconv.FormatInt(time.Now().Unix(), 10) + common.RandNum() //randomInt由时间戳和一个随机数组成的一个int
	randomStr := l.Logistics_name + randomInt                                //物流员名字加上随机int产生的一个str
	upSql, err := l.Db.Prepare("UPDATE shop_logistics set token = ?, login_time = ?, app_token = ? WHERE logistics_id = ? LIMIT 1")
	if err != nil {
		log.Fatalf("Build update token prepare token err %v", err)
	}
	res, err := upSql.Exec(common.STMd5(randomStr), int(time.Now().Unix()), l.App_token, l.Logistics_id)
	if err != nil {
		log.Fatalf("Exec update token sql err %v", err)
	}
	rowCnt, err := res.RowsAffected()
	if rowCnt != 1 || err != nil {
		log.Fatalf("Affect row not only one, %v", err)
	}
}

//判断是否登录
func (l *Shop_logistics) IsLogin() (string, error) {
	stmt, err := l.Db.Prepare("SELECT count(1) as count FROM shop_logistics WHERE token = ?")
	if err != nil {
		log.Fatalf("Build islogin sql err, %v", err)
	}
	var count int
	err = stmt.QueryRow(l.Token).Scan(&count)
	if err != nil {
		log.Fatalf("Is not login, %v", err)
	}
	if count != 1 {
		return "Please login", err
	}
	return "login suc", err
}

func (d *Shop_v_dispatch) GetDispatchList(l *Shop_logistics) []map[string]interface{} {
	stmt, err := d.Db.Prepare("SELECT dispatch_id, order_sn FROM shop_v_dispatch WHERE logistics_id = ? LIMIT 30")

	if err != nil {
		log.Fatalf("Get dispatch list err, %v", err)
	}

	defer stmt.Close()
	rows, err := stmt.Query(l.Logistics_id)
	if err != nil {
		log.Fatalf("some err, %v", err)
	}

	var dispatchList []map[string]interface{}
	for rows.Next() {
		var d Shop_v_dispatch
		err = rows.Scan(&d.Dispatch_id, &d.Order_sn)
		if err != nil {
			log.Fatalf("for err, %v", err)
		}
		args := map[string]interface{}{
			"dispatch_id": d.Dispatch_id,
			"order_sn":    d.Order_sn,
		}
		dispatchList = append(dispatchList, args)
	}
	return dispatchList
}
