package endpoint

import (
	"context"
	"fmt"

	pb "git.bluebird.id/Restoran/BahanMenu/bahanmenu/grpc"
	scv "git.bluebird.id/Restoran/BahanMenu/bahanmenu/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcBahanMenuServer struct {
	addBahanMenu                 grpctransport.Handler
	readBahanMenu                grpctransport.Handler
	updateBahanMenu              grpctransport.Handler
	readBahanMenuByNamaBahanMenu grpctransport.Handler
}

func NewGRPCBahanMenuServer(endpoints BahanMenuEndpoint, tracer stdopentracing.Tracer,

	logger log.Logger) pb.BahanMenuServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcBahanMenuServer{
		addBahanMenu: grpctransport.NewServer(endpoints.AddBahanMenuEndpoint,
			decodeAddBahanMenuRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddBahanMenu", logger)))...),
		readBahanMenu: grpctransport.NewServer(endpoints.ReadBahanMenuEndpoint,
			decodeReadBahanMenuRequest,
			encodeReadBahanMenuResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadBahanMenu", logger)))...),
		updateBahanMenu: grpctransport.NewServer(endpoints.UpdateBahanMenuEndpoint,
			decodeUpdateBahanMenuRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateBahanMenu", logger)))...),
		readBahanMenuByNamaBahanMenu: grpctransport.NewServer(endpoints.ReadBahanMenuByNamaBahanMenuEndpoint,
			decodeReadBahanMenuByNamaBahanMenuRequest,
			encodeReadBahanMenuByNamaBahanMenuResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadBahanMenuByNamaBahanMenu", logger)))...),
	}
}

func decodeAddBahanMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	fmt.Println("tes masuk ga")
	req := request.(*pb.AddBahanMenuReq)
	return scv.BahanMenu{NamaBahan: req.GetNamabahan(), Quantity: req.GetQuantity(), CreatedBy: req.GetCreatedby()}, nil
}

func decodeReadBahanMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadBahanMenuByNamaBahanMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadBahanMenuByNamaBahanMenuReq)
	return scv.BahanMenu{NamaBahan: req.Namabahan}, nil
}

func decodeUpdateBahanMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateBahanMenurReq)
	return scv.BahanMenu{IDBahan: req.Idbahan, NamaBahan: req.Namabahan, Quantity: req.Quantity, Status: req.Status, UpdateBy: req.UpdateBy}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadBahanMenuByNamaBahanMenuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.BahanMenu)
	return &pb.ReadBahanMenuByNamaBahanMenuResp{Idbahan: resp.IDBahan, Namabahan: resp.NamaBahan, Status: resp.Status, Quantity: resp.Quantity, Createdby: resp.CreatedBy}, nil
}
func encodeReadBahanMenuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.BahanMenus)

	rsp := &pb.ReadBahanMenuResp{}

	for _, v := range resp {
		itm := &pb.ReadBahanMenuByNamaBahanMenuResp{
			Idbahan:   v.IDBahan,
			Namabahan: v.NamaBahan,
			Quantity:  v.Quantity,
			Status:    v.Status,
			Createdby: v.CreatedBy,
		}
		rsp.AllBahanMenu = append(rsp.AllBahanMenu, itm)
	}
	return rsp, nil
}

func (s *grpcBahanMenuServer) AddBahanMenu(ctx oldcontext.Context, bahanmenu *pb.AddBahanMenuReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addBahanMenu.ServeGRPC(ctx, bahanmenu)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcBahanMenuServer) ReadBahanMenu(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadBahanMenuResp, error) {
	_, resp, err := s.readBahanMenu.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadBahanMenuResp), nil
}

func (s *grpcBahanMenuServer) UpdateBahanMenu(ctx oldcontext.Context, cus *pb.UpdateBahanMenurReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateBahanMenu.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcBahanMenuServer) ReadBahanMenuByNamaBahanMenu(ctx oldcontext.Context, namabahan *pb.ReadBahanMenuByNamaBahanMenuReq) (*pb.ReadBahanMenuByNamaBahanMenuResp, error) {
	_, resp, err := s.readBahanMenuByNamaBahanMenu.ServeGRPC(ctx, namabahan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadBahanMenuByNamaBahanMenuResp), nil
}
