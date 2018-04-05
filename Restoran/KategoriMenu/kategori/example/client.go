package main

import (
	"context"
	"time"

	cli "git.bluebird.id/Restoran/KategoriMenu/kategori/endpoint"
	svc "git.bluebird.id/Restoran/KategoriMenu/kategori/server"
	opt "git.bluebird.id/Restoran/KategoriMenu/util/grpc"
	util "git.bluebird.id/Restoran/Menu/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCKategoriMenuClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}
	//client.AddKategoriMenuService(context.Background(), svc.KategoriMenu{KategoriMenu: "Snack", Deskripsi: "Makanan Ringan", CreatedBy: "Jono"})
	//kategorimenu, _ := client.ReadKategoriMenuByKategoriMenuService(context.Background(), "Snack")
	//fmt.Println("Kategori based on KategoriMenu:", kategorimenu)
	//menn, _ := client.ReadKategoriMenuService(context.Background())
	//fmt.Println("all Kategori:", menn)
	client.UpdateKategoriMenuService(context.Background(), svc.KategoriMenu{KategoriMenu: "snake", Deskripsi: "makanan ringan", UpdateBy: "Joko", IDKategori: 1})

}
