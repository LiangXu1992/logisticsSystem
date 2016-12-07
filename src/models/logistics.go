package models

import (
    "databases"
    //"common"
    //"time"
    //"strconv"
    "fmt"

    "common"
    "log"
    "strconv"
    "time"
)

//连接数据库
var conn databases.Conn
func init() {
    //初始化mysql
    conn.InitDB()
}

//物流员结构体
type Shop_logistics struct {
    Logistics_id int `gorm:"primary_key"`
    Logistics_name string
    Token string
    Client_type string
    Password string
    Login_time int
    App_token string
}

//物流单结构体
type Shop_v_dispatch struct {
    Dispatch_id int
    Order_sn string
}

func (l *Shop_logistics)GetLogisticsInfo() {
    //查找物流员的信息
    err := conn.DB.QueryRow("SELECT logistics_id, token FROM shop_logistics where logistics_name = ? AND logistics_pwd = ?", l.Logistics_name, common.STMd5(l.Password)).Scan(&l.Logistics_id, &l.Token)
    if err != nil {
        log.Fatalf("Query logistics error %v", err)
    }

    //更新token，apptoken
    l.UpdateToken()
}

func (l *Shop_logistics)UpdateToken() {
    //更新物流员登录token
    randomInt := strconv.FormatInt(time.Now().Unix(), 10) + common.RandNum()                   //randomInt由时间戳和一个随机数组成的一个int
    randomStr := l.Logistics_name + randomInt                                                  //物流员名字加上随机int产生的一个str
    stmp, err := conn.DB.Prepare("UPDATE shop_logistics set token = ?,login_time = ? WHERE logistics_id = ?")
    if err != nil {
        log.Fatalf("Build update token prepare token err %v", err)
    }
    res, err := stmp.Exec(common.STMd5(randomStr), int(time.Now().Unix()), l.Logistics_id)
    if err != nil {
        log.Fatalf("Exec update token sql err %v", err)
    }
    rowCnt, err := res.RowsAffected()
    lastId, err := res.LastInsertId()
    if(rowCnt != 1 || err != nil) {
        log.Fatalf("Affect row not only one, %v", err)
    }

    //更新物流员推送token
    fmt.Println(lastId)
}

//判断是否登录
func (l *Shop_logistics)IsLogin() {
    //conn.DB.Where("token = ?", l.Token).First(l)
}

func (d *Shop_v_dispatch) GetDispatchList(l *Shop_logistics) {
    //rows, err := conn.DB.Raw("SELECT * FROM shop_v_dispatch WHERE logistics_id = ?", l.Logistics_id).Rows()
    //if (err != nil) {
    //    return
    //}
    //var dispatch_id int
    //for rows.Next() {
    //    rows.Scan(&dispatch_id, d.Order_sn)
    //    fmt.Println(dispatch_id)
    //}
    //conn.DB.Find(d)
    fmt.Println(d)


}
