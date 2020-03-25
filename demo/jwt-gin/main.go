//### 注意：本项目没有使用Go Modules，所以请先关闭它或设置成auto，不然会出现一些问题。
//运行指令：`go env -w GO111MODULE=off`或者`go env -w GO111MODULE=auto`

package main

import (
	"colasoft.rd.go.backend/grpcservice"
	"colasoft.rd.go.backend/pkg/setting"
	"colasoft.rd.go.backend/routers"
	"fmt"
	"net/http"
)

func main() {

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go grpcservice.StartGrpcService()
	s.ListenAndServe()
}
