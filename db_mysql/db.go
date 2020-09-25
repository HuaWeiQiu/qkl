package db_mysql

import (
	"beeproject/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	fmt.Printf("连接mysql数据库")
	config := beego.AppConfig
	dbDriver := config.String("db_drivername")
	dbUser := config.String("db_user")
	dbPwd := config.String("db_pwd")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	connUrl := dbUser + ":" + dbPwd + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("数据连接错误,请检查配置")
	}
	DB = db
}
func InterUser(user models.User) (int64, error) {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes := hashMd5.Sum(nil)
	user.Password = hex.EncodeToString(bytes)
	result, err := DB.Exec("insert into user(nick,password) values(?,?)", user.Nick, user.Password)
	if err != nil {
		return -1, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, err
}
func QueryUser() {
	DB.QueryRow("select * from ")
}
