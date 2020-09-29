package main

import (
	"fmt"
	"net/http"

	"go-gin-demo/pkg/setting"
	"go-gin-demo/routers"
)

func main() {

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        routers.InitRouters(),
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
