package main

import (
	"blog-service/global"
	"blog-service/internal/routers"
	setting "blog-service/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}
}

func main() {
	fmt.Println(global.ServerSetting)
	fmt.Println(global.AppSetting)
	fmt.Println(global.DatabaseSetting)
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil { //如果有错误的话
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil

}
