package endpoint

import (
	"context"
	"fmt"

	svc "git.bluebird.id/Restoran/BahanMenu/bahanmenu/server"
	kit "github.com/go-kit/kit/endpoint"
)

type BahanMenuEndpoint struct {
	AddBahanMenuEndpoint                 kit.Endpoint
	ReadBahanMenuEndpoint                kit.Endpoint
	UpdateBahanMenuEndpoint              kit.Endpoint
	ReadBahanMenuByNamaBahanMenuEndpoint kit.Endpoint
}

func NewBahanMenuEndpoint(service svc.BahanMenuService) BahanMenuEndpoint {
	addBahanMenuEp := makeAddBahanMenuEndpoint(service)
	readBahanMenuEp := makeReadBahanMenuEndpoint(service)
	updateBahanMenuEp := makeUpdateBahanMenuEndpoint(service)
	readBahanMenuByNamaBahanMenuEp := makeReadBahanMenuByNamaBahanMenuEndpoint(service)
	return BahanMenuEndpoint{AddBahanMenuEndpoint: addBahanMenuEp,
		ReadBahanMenuEndpoint:                readBahanMenuEp,
		UpdateBahanMenuEndpoint:              updateBahanMenuEp,
		ReadBahanMenuByNamaBahanMenuEndpoint: readBahanMenuByNamaBahanMenuEp,
	}
}

func makeAddBahanMenuEndpoint(service svc.BahanMenuService) kit.Endpoint {
	fmt.Println("masuk sini")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.BahanMenu)
		err := service.AddBahanMenuService(ctx, req)
		return nil, err
	}
}

func makeReadBahanMenuEndpoint(service svc.BahanMenuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadBahanMenuService(ctx)
		return result, err
	}
}

func makeUpdateBahanMenuEndpoint(service svc.BahanMenuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.BahanMenu)
		err := service.UpdateBahanMenuService(ctx, req)
		return nil, err
	}
}
func makeReadBahanMenuByNamaBahanMenuEndpoint(service svc.BahanMenuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.BahanMenu)
		result, err := service.ReadBahanMenuByNamaBahanMenuService(ctx, req.NamaBahan)
		return result, err
	}
}
