package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	grpcRouter "github.com/agclqq/study_tools/app/grpc/router"
	"github.com/agclqq/study_tools/config"
)

func main() {
	wg := &sync.WaitGroup{}

	//wg.Add(1)
	//go pprofServer(wg)

	wg.Add(1)
	go grpcServer(wg)

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
			fmt.Printf("start pprofServer is error: %s\n", err)
		}
	}()
}

func grpcServer(wg *sync.WaitGroup) {
	defer wg.Done()

	lis, err := net.Listen("tcp", ":"+config.GetApp("grpc_server_port"))
	if err != nil {
		_ = fmt.Errorf("failed to listen: %v", err)
	}

	kp := keepalive.ServerParameters{
		Time:    20 * time.Second,
		Timeout: 5 * time.Second,
	}
	kep := keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}

	s := grpc.NewServer(
		grpc.KeepaliveParams(kp),
		grpc.KeepaliveEnforcementPolicy(kep),
	)

	grpcRouter.Register(s)

	go func() {
		fmt.Printf("start grpc server at: %s\n", lis.Addr().String())
		if err = s.Serve(lis); err != nil {
			fmt.Printf("start grpc server is error: %v", err)
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	sign := <-ch
	fmt.Println("got a sign:", sign)
	now := time.Now()
	s.GracefulStop()
	// 看看实际退出所耗费的时间
	fmt.Println("grpc server is exited,cost:", time.Since(now).Milliseconds(), "ms")
}
