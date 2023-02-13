package main

import (
	"context"
	"github.com/bobgo0912/b0b-common/pkg/config"
	"github.com/bobgo0912/b0b-common/pkg/etcd"
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/bobgo0912/b0b-common/pkg/nats"
	"github.com/bobgo0912/b0b-common/pkg/redis"
	"github.com/bobgo0912/b0b-common/pkg/server"
	on "github.com/bobgo0912/bob-order/interal/nats"
	"github.com/bobgo0912/bob-order/interal/rpc"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, ca := context.WithCancel(context.Background())
	log.InitLog()
	newConfig := config.NewConfig(config.Json)
	newConfig.Category = "../config"
	newConfig.InitConfig()
	mainServer := server.NewMainServer()
	etcdClient := etcd.NewClientFromCnf()
	grpcServer := server.NewGrpcServer(config.Cfg.Host, config.Cfg.RpcPort)
	rpc.RegService(grpcServer)
	mainServer.AddServer(grpcServer)
	err := mainServer.Start(ctx)
	if err != nil {
		log.Panic(err)
	}
	client, err := redis.NewClient()
	if err != nil {
		log.Panic(err)
	}
	natsClient, err := nats.NewJetClient()
	if err != nil {
		log.Panic(err)
	}
	handler := on.NewOrderHandler(natsClient)
	err = handler.Start(ctx, client)
	if err != nil {
		log.Panic(err)
	}
	mainServer.Discover(ctx, etcdClient)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	ca()
	time.Sleep(3 * time.Second)
}
