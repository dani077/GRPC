package endpoint

import (
	"context"
	"fmt"
	"time"

	pb "git.bluebird.id/Restoran/BahanMenu/bahanmenu/grpc"
	svc "git.bluebird.id/Restoran/BahanMenu/bahanmenu/server"

	util "git.bluebird.id/Restoran/BahanMenu/util/grpc"
	disc "git.bluebird.id/Restoran/BahanMenu/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.BahanMenuService"
)

func NewGRPCBahanMenuClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.BahanMenuService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addBahanMenuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddBahanMenuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addBahanMenuEp = retry
	}

	var readBahanMenuByNamaBahanMenuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadBahanMenuByNamaBahanMenuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readBahanMenuByNamaBahanMenuEp = retry
	}

	var readBahanMenuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadBahanMenuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readBahanMenuEp = retry
	}

	var updateBahanMenuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateBahanMenu, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateBahanMenuEp = retry
	}

	return BahanMenuEndpoint{AddBahanMenuEndpoint: addBahanMenuEp, ReadBahanMenuByNamaBahanMenuEndpoint: readBahanMenuByNamaBahanMenuEp, ReadBahanMenuEndpoint: readBahanMenuEp, UpdateBahanMenuEndpoint: updateBahanMenuEp}, nil
}
func encodeAddBahanMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	fmt.Println("tes")
	req := request.(svc.BahanMenu)
	return &pb.AddBahanMenuReq{
		Namabahan: req.NamaBahan, Quantity: req.Quantity, Createdby: req.CreatedBy,
	}, nil
}
func encodeReadBahanMenuByNamaBahanMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.BahanMenu)
	return &pb.ReadBahanMenuByNamaBahanMenuReq{Namabahan: req.NamaBahan}, nil
}

func encodeReadBahanMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateBahanMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.BahanMenu)
	return &pb.UpdateBahanMenurReq{
		Idbahan:   req.IDBahan,
		Namabahan: req.NamaBahan,
		Quantity:  req.Quantity,
		Status:    req.Status,
		UpdateBy:  req.UpdateBy,
	}, nil
}
func decodeBahanMenuResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadBahanMenuByNamaBahanMenuRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadBahanMenuByNamaBahanMenuResp)
	return svc.BahanMenu{
		IDBahan:   resp.Idbahan,
		NamaBahan: resp.Namabahan,
		Quantity:  resp.Quantity,
		Status:    resp.Status,
		CreatedBy: resp.Createdby,
	}, nil
}

func decodeReadBahanMenuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadBahanMenuResp)
	var rsp svc.BahanMenus

	for _, v := range resp.AllBahanMenu {
		itm := svc.BahanMenu{
			IDBahan:   v.Idbahan,
			NamaBahan: v.Namabahan,
			Quantity:  v.Quantity,
			Status:    v.Status,
			CreatedBy: v.Createdby,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddBahanMenuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	fmt.Println("masuk makeclient")
	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddBahanMenu",
		encodeAddBahanMenuRequest,
		decodeBahanMenuResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddBahanMenu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddBahanMenu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
func makeClientReadBahanMenuByNamaBahanMenuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadBahanMenuByNamaBahanMenu",
		encodeReadBahanMenuByNamaBahanMenuRequest,
		decodeReadBahanMenuByNamaBahanMenuRespones,
		pb.ReadBahanMenuByNamaBahanMenuResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadBahanMenuByNamaBahanMenu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadBahanMenuByNamaBahanMenu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadBahanMenuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadBahanMenu",
		encodeReadBahanMenuRequest,
		decodeReadBahanMenuResponse,
		pb.ReadBahanMenuResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadBahanMenu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadBahanMenu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateBahanMenu(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateBahanMenu",
		encodeUpdateBahanMenuRequest,
		decodeBahanMenuResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateBahanMenu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateBahanMenu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
