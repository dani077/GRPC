package endpoint

import (
	"context"

	svc "git.bluebird.id/Restoran/KategoriMenu/kategori/server"
	kit "github.com/go-kit/kit/endpoint"
)

type KategoriMenuEndpoint struct {
	AddKategoriMenuEndpoint                kit.Endpoint
	ReadKategoriMenuEndpoint               kit.Endpoint
	UpdateKategoriMenuEndpoint             kit.Endpoint
	ReadKategoriMenuByKategoriMenuEndpoint kit.Endpoint
}

func NewKategoriMenuEndpoint(service svc.KategoriMenuService) KategoriMenuEndpoint {
	addKategoriMenuEp := makeAddKategoriMenuEndpoint(service)
	readKategoriMenuEp := makeReadKategoriMenuEndpoint(service)
	updateKategoriMenuEp := makeUpdateKategoriMenuEndpoint(service)
	readKategoriMenuByKategoriMenuEp := makeReadKategoriMenuByKategoriMenuEndpoint(service)
	return KategoriMenuEndpoint{AddKategoriMenuEndpoint: addKategoriMenuEp,
		ReadKategoriMenuEndpoint:               readKategoriMenuEp,
		UpdateKategoriMenuEndpoint:             updateKategoriMenuEp,
		ReadKategoriMenuByKategoriMenuEndpoint: readKategoriMenuByKategoriMenuEp,
	}
}

func makeAddKategoriMenuEndpoint(service svc.KategoriMenuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.KategoriMenu)
		err := service.AddKategoriMenuService(ctx, req)
		return nil, err
	}
}

func makeReadKategoriMenuEndpoint(service svc.KategoriMenuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadKategoriMenuService(ctx)
		return result, err
	}
}

func makeUpdateKategoriMenuEndpoint(service svc.KategoriMenuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.KategoriMenu)
		err := service.UpdateKategoriMenuService(ctx, req)
		return nil, err
	}
}
func makeReadKategoriMenuByKategoriMenuEndpoint(service svc.KategoriMenuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.KategoriMenu)
		result, err := service.ReadKategoriMenuByKategoriMenuService(ctx, req.KategoriMenu)
		return result, err
	}
}
