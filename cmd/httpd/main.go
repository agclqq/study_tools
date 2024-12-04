package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/agclqq/prow-framework/logger"
	_ "github.com/agclqq/prow-framework/validator"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/agclqq/study_tools/app/http/router"
	"github.com/agclqq/study_tools/boot"
	"github.com/agclqq/study_tools/config"
)

func main() {
	wg := &sync.WaitGroup{}

	//wg.Add(1)
	//go pprofServer(wg)

	wg.Add(1)
	go httpServer(wg)

	wg.Wait()
}

func pprofServer(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("start pprofServer at: %s\n", "6060")
	server := &http.Server{
		Addr:    ":6060",
		Handler: nil,
	}
	go func() {
		fmt.Printf("start pprofServer at: %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				fmt.Println(err)
				return
			}
			_ = fmt.Errorf("start pprofServer is error: %s\n", err)
		}
	}()
}

func httpServer(wg *sync.WaitGroup) {
	defer wg.Done()
	if config.GetApp("appEnv") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	//注册事件
	boot.StartEvent()

	eng := gin.New()
	eng.RedirectTrailingSlash = false
	eng.Use(logger.WithConfig(logger.AccessLogConfig(eng, config.GetApp("logFile"), cast.ToInt(config.GetApp("logRetain")))), gin.Recovery())
	router.Register(eng)
	eng.StaticFS("/resource", gin.Dir("./resource", false))
	eng.LoadHTMLGlob("resource/views/**/*")
	server := &http.Server{
		Addr:        ":" + config.GetApp("httpServerPort"),
		Handler:     eng,
		IdleTimeout: 75 * time.Second,
	}
	go func() {
		//if err := server.ListenAndServeTLS("resource/cert.pem", "resource/cert.key"); err != nil {
		//	fmt.Printf("start https server is error: %s\n", err)
		//}
		//(&provider.Event{}).Run()
		fmt.Printf("start http server at: %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				fmt.Println(err)
				return
			}
			fmt.Errorf("start http server is error: %s\n", err)
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	sign := <-ch
	fmt.Println("got a sign:", sign)
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	// 看看实际退出所耗费的时间
	fmt.Println("http server is exited,cost:", time.Since(now).Milliseconds(), "ms")
}
