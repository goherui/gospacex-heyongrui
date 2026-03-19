package main

import (
	"context"
	"flag"
	"fmt"
	"gospacex/goods/goods-service/basic/initializer"
	_ "gospacex/goods/goods-service/basic/initializer"
	payments "gospacex/goods/goods-service/handler/service/payment"
	__ "gospacex/proto/payment"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8081, "The server port")
)

func main() {
	flag.Parse()
	err := initializer.ConsulInit()
	if err != nil {
		fmt.Println("consul初始化失败", err)
	}
	fmt.Println("consul初始化成功")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//product.RegisterStreamGreeterServer(s, &products.Server{})
	__.RegisterStreamGreeterServer(s, &payments.Saver{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := initializer.ConsulShutdown(); err != nil {
		log.Printf("Consul注销失败: %v", err)
	}
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.GracefulStop()
	log.Println("服务已关闭")
}
