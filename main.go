package main

import (
	"github.com/zhangxuesong/josephblog/pkg/config"
	"github.com/zhangxuesong/josephblog/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.InitRouter()

	server := &http.Server{
		Addr:         config.Config.Service.Port,
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		MaxHeaderBytes: 2 << 20,
	}
	server.ListenAndServe()
}
