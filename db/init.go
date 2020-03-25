package db

import (
	"AdapterServer/conf"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DbHandle *sql.DB

func InitDbConn() {

	var dbError error

	dbCfg := conf.GetDbCfg()

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		dbCfg.UserName, dbCfg.Password, dbCfg.IpAddress, dbCfg.Port, dbCfg.DbName)

	// 打开连接失败
	DbHandle, dbError = sql.Open("mysql", dbDSN)
	//defer DbHandle.Close();

	if dbError != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + dbError.Error())
	}

	// 最大连接数
	DbHandle.SetMaxOpenConns(100)

	// 闲置连接数
	DbHandle.SetMaxIdleConns(20)

	// 最大连接周期
	DbHandle.SetConnMaxLifetime(100 * time.Second)

	if dbError = DbHandle.Ping(); nil != dbError {
		panic("数据库链接失败: " + dbError.Error())
	}
}
