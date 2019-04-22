package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tozastation/clasick-core/interface/handler/di"
	"github.com/tozastation/clasick-core/interface/handler/gorm"
	rpcuser "github.com/tozastation/clasick-core/interface/rpc/user"
	"log"
	"net"
	"time"
)

func main() {
	/*
	 * 初期化処理
	 * 1: ポート作成
	 * 2: RPCを登録
	 * 3: サービススタート
	 */
	listenPort, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Wait Cloud Proxy Process")
	time.Sleep(2 * time.Second)
	fmt.Println("Dependency Injection")
	diHandler := di.NewDIHandler(gorm.NewGormHandler())
	grpcServer := diHandler.CreateServer()
	rpcuser.RegisterUserRpcServer(grpcServer, diHandler.InitializeUser())
	err = grpcServer.Serve(listenPort)
	if err != nil {
		panic(err)
	}
}
