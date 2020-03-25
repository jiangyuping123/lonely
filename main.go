package main

import (
	"AdapterServer/conf"
	"AdapterServer/db"
	"AdapterServer/grpcservice"
	"AdapterServer/routers"
	"AdapterServer/syncdata"
	"fmt"
	"net/http"
)

func main() {

	conf.LoadCfg()
	db.InitDbConn()

	httpServerCfg := conf.GetHttpServerCfg()
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", httpServerCfg.HTTPPort),
		Handler:        router,
		ReadTimeout:    httpServerCfg.ReadTimeout,
		WriteTimeout:   httpServerCfg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go grpcservice.StartGrpcService()
	go syncdata.SyncData()
	s.ListenAndServe()
}
