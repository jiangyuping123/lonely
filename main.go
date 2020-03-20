package main

import (
	"AdapterServer/conf"
	"AdapterServer/db"
	"AdapterServer/grpcservice"
	"AdapterServer/routers"
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
	s.ListenAndServe()
}
