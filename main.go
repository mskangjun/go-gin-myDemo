package main

import (
	"fmt"
	"net/http"

	"myDemo/pkg/setting"
	"myDemo/routers"
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

	_ = s.ListenAndServe()
}
