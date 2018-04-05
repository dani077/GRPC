package main

import (
	"context"
	"time"

	cli "git.bluebird.id/Restoran/BahanMenu/bahanmenu/endpoint"
	svc "git.bluebird.id/Restoran/BahanMenu/bahanmenu/server"
	opt "git.bluebird.id/Restoran/BahanMenu/util/grpc"
	util "git.bluebird.id/Restoran/BahanMenu/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCBahanMenuClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}
	//Add Menu
	//client.AddBahanMenuService(context.Background(), svc.BahanMenu{NamaBahan: "Susu Botol", Quantity: 9, CreatedBy: "Joloollo"})
	//namaBahan, _ := client.ReadBahanMenuByNamaBahanMenuService(context.Background(), "Gula")
	//fmt.Println("Menu based on Namabahan:", namaBahan)
	//menn, _ := client.ReadBahanMenuService(context.Background())
	//fmt.Println("all Menu:", menn)
	client.UpdateBahanMenuService(context.Background(), svc.BahanMenu{NamaBahan: "Gula", Quantity: 10, UpdateBy: "Joko", Status: 1, IDBahan: 1})

}
