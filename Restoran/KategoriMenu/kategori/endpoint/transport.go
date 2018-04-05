package endpoint

import (
	"context"

	pb "git.bluebird.id/Restoran/KategoriMenu/kategori/grpc"
	scv "git.bluebird.id/Restoran/KategoriMenu/kategori/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcKategoriMenuServer struct {
	addKategoriMenu                grpctransport.Handler
	readKategoriMenu               grpctransport.Handler
	updateKategoriMenu             grpctransport.Handler
	readKategoriMenuByKategoriMenu grpctransport.Handler
}

func NewGRPCKategoriMenuServer(endpoints KategoriMenuEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.KategoriMenuServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcKategoriMenuServer{
		addKategoriMenu: grpctransport.NewServer(endpoints.AddKategoriMenuEndpoint,
			decodeAddKategoriMenuRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddKategoriMenu", logger)))...),
		readKategoriMenu: grpctransport.NewServer(endpoints.ReadKategoriMenuEndpoint,
			decodeReadKategoriMenuRequest,
			encodeReadKategoriMenuResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKategoriMenu", logger)))...),
		updateKategoriMenu: grpctransport.NewServer(endpoints.UpdateKategoriMenuEndpoint,
			decodeUpdateKategoriMenuRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateKategoriMenu", logger)))...),
		readKategoriMenuByKategoriMenu: grpctransport.NewServer(endpoints.ReadKategoriMenuByKategoriMenuEndpoint,
			decodeReadKategoriMenuByKategoriMenuRequest,
			encodeReadKategoriMenuByKategoriMenuResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKategoriMenuByKategoriMenu", logger)))...),
	}
}

func decodeAddKategoriMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddKategoriMenuReq)
	return scv.KategoriMenu{KategoriMenu: req.GetKategorimenu(), Deskripsi: req.GetDeskripsi(), CreatedBy: req.GetCreatedby()}, nil
}

func decodeReadKategoriMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadKategoriMenuByKategoriMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKategoriMenuByKategoriMenuReq)
	return scv.KategoriMenu{KategoriMenu: req.Kategorimenu}, nil
}

func decodeUpdateKategoriMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateKategoriMenurReq)
	return scv.KategoriMenu{IDKategori: req.Idkategorimenu, KategoriMenu: req.Kategorimenu, Deskripsi: req.Deskripsi, UpdateBy: req.UpdateBy}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadKategoriMenuByKategoriMenuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.KategoriMenu)
	return &pb.ReadKategoriMenuByKategoriMenuResp{Idkategorimenu: resp.IDKategori, Kategorimenu: resp.KategoriMenu, Deskripsi: resp.Deskripsi, Status: resp.Status, Createdby: resp.CreatedBy}, nil
}
func encodeReadKategoriMenuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.KategoriMenus)

	rsp := &pb.ReadKategoriMenuResp{}

	for _, v := range resp {
		itm := &pb.ReadKategoriMenuByKategoriMenuResp{
			Idkategorimenu: v.IDKategori,
			Kategorimenu:   v.KategoriMenu,
			Deskripsi:      v.Deskripsi,
			Status:         v.Status,
			Createdby:      v.CreatedBy,
		}
		rsp.AllKategoriMenu = append(rsp.AllKategoriMenu, itm)
	}
	return rsp, nil
}

func (s *grpcKategoriMenuServer) AddKategoriMenu(ctx oldcontext.Context, menu *pb.AddKategoriMenuReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addKategoriMenu.ServeGRPC(ctx, menu)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcKategoriMenuServer) ReadKategoriMenu(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadKategoriMenuResp, error) {
	_, resp, err := s.readKategoriMenu.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKategoriMenuResp), nil
}

func (s *grpcKategoriMenuServer) UpdateKategoriMenu(ctx oldcontext.Context, cus *pb.UpdateKategoriMenurReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateKategoriMenu.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcKategoriMenuServer) ReadKategoriMenuByKategoriMenu(ctx oldcontext.Context, kategorimenu *pb.ReadKategoriMenuByKategoriMenuReq) (*pb.ReadKategoriMenuByKategoriMenuResp, error) {
	_, resp, err := s.readKategoriMenuByKategoriMenu.ServeGRPC(ctx, kategorimenu)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKategoriMenuByKategoriMenuResp), nil
}
