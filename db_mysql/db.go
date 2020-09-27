package db_mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"kekele/models"
)
var Db *sql.DB
//连接数据库
func init() {
	fmt.Println("MYSQL数据库已连接")
	config :=beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbName := config.String("db_name")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_pssword")
	dbIp := config.String("db_ip")

	connUrl := dbUser + ":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	fmt.Println(connUrl)
	db, err :=sql.Open(dbDriver,connUrl)
	if err != nil {
		panic("数据库连接错误，请检查配置")
	}

	Db = db

}
//将用户信息保存到数据库表当中
func InsertUser(user models.Sql) (int64,error) {
	//将密码转换成哈希，并储存hash值
	hashMd5 :=md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes := hashMd5.Sum(nil)
	user.Password = hex.EncodeToString(bytes)
	//fmt.Println("要保存的用户名是", user.Nick, "要保存的密码是", user.Password)

	result, err := Db.Exec("insert into user(name,password) value (?,?)", user.Nick, user.Password)
	if err != nil {
		return -1,err
	}
	id,err :=result.RowsAffected()
	if err != nil {
		return -1,err
	}

	return id,nil
}
