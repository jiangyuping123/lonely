package conf

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type HttpServerCfg struct {
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DataBaseCfg struct {
	DbType    string
	DbName    string
	UserName  string
	Password  string
	IpAddress string
	Port      int
}

type AppCfg struct {
	RunMode string
}

var httpServerCfg HttpServerCfg
var dbCfg DataBaseCfg
var appCfg AppCfg

func GetDbCfg() DataBaseCfg {
	return dbCfg
}

func GetAppCfg() AppCfg {
	return appCfg
}

func GetHttpServerCfg() HttpServerCfg {
	return httpServerCfg
}

func LoadCfg() {

	var err error
	var Cfg *ini.File

	Cfg, err = ini.Load("conf/config/cfg.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/config/cfg.ini': %v", err)
	}

	////////////////////////// app ////////////////////////////////////
	sec, err := Cfg.GetSection("App")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	appCfg.RunMode = sec.Key("RUN_MODE").MustString("debug")

	///////////////////////// httpserver /////////////////////////////
	sec, err = Cfg.GetSection("HttpServer")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	httpServerCfg.HTTPPort = sec.Key("HTTP_PORT").MustInt(8888)
	httpServerCfg.ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	httpServerCfg.WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

	/////////////////////////// database ///////////////////////////
	sec, err = Cfg.GetSection("DataBase")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbCfg.DbType = sec.Key("DBTYPE").MustString("mysql")
	dbCfg.DbName = sec.Key("DBNAME").MustString("")
	dbCfg.UserName = sec.Key("USERNAME").MustString("")
	dbCfg.Password = sec.Key("PASSWORD").MustString("")
	dbCfg.IpAddress = sec.Key("IPADDRESS").MustString("")
	dbCfg.Port = sec.Key("PORT").MustInt(3306)
}
