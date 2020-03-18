package main

import (
	"AdapterServer/conf"
	"AdapterServer/grpcservice"
	"AdapterServer/routers"
	"fmt"
	"net/http"
)

func main() {

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.HTTPPort),
		Handler:        router,
		ReadTimeout:    conf.ReadTimeout,
		WriteTimeout:   conf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go grpcservice.StartGrpcService()
	s.ListenAndServe()
}
