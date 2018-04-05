package endpoint

import (
	"context"
	"time"

	pb "git.bluebird.id/Restoran/KategoriMenu/kategori/grpc"
	svc "git.bluebird.id/Restoran/KategoriMenu/kategori/server"

	util "git.bluebird.id/Restoran/KategoriMenu/util/grpc"
	disc "git.bluebird.id/Restoran/KategoriMenu/util/microservice"

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
	grpcName = "grpc.KategoriMenuService"
)

func NewGRPCKategoriMenuClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.KategoriMenuService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addKategoriMenuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddKategoriMenuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addKategoriMenuEp = retry
	}

	var readKategoriMenuByKategoriMenuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKategoriMenuByKategoriMenuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKategoriMenuByKategoriMenuEp = retry
	}

	var readKategoriMenuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKategoriMenuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKategoriMenuEp = retry
	}

	var updateKategoriMenuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateKategoriMenu, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateKategoriMenuEp = retry
	}

	return KategoriMenuEndpoint{AddKategoriMenuEndpoint: addKategoriMenuEp, ReadKategoriMenuByKategoriMenuEndpoint: readKategoriMenuByKategoriMenuEp, ReadKategoriMenuEndpoint: readKategoriMenuEp, UpdateKategoriMenuEndpoint: updateKategoriMenuEp}, nil
}
func encodeAddKategoriMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.KategoriMenu)
	return &pb.AddKategoriMenuReq{
		Kategorimenu: req.KategoriMenu, Deskripsi: req.Deskripsi, Createdby: req.CreatedBy,
	}, nil
}
func encodeReadKategoriMenuByKategoriMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.KategoriMenu)
	return &pb.ReadKategoriMenuByKategoriMenuReq{Kategorimenu: req.KategoriMenu}, nil
}

func encodeReadKategoriMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateKategoriMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.KategoriMenu)
	return &pb.UpdateKategoriMenurReq{
		Idkategorimenu: req.IDKategori,
		Kategorimenu:   req.KategoriMenu,
		Deskripsi:      req.Deskripsi,
		UpdateBy:       req.UpdateBy,
	}, nil
}
func decodeKategoriMenuResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadKategoriMenuByKategoriMenuRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKategoriMenuByKategoriMenuResp)
	return svc.KategoriMenu{
		IDKategori:   resp.Idkategorimenu,
		KategoriMenu: resp.Kategorimenu,
		Deskripsi:    resp.Deskripsi,
		Status:       resp.Status,
		CreatedBy:    resp.Createdby,
	}, nil
}

func decodeReadKategoriMenuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKategoriMenuResp)
	var rsp svc.KategoriMenus

	for _, v := range resp.AllKategoriMenu {
		itm := svc.KategoriMenu{
			IDKategori:   v.Idkategorimenu,
			KategoriMenu: v.Kategorimenu,
			Deskripsi:    v.Deskripsi,
			Status:       v.Status,
			CreatedBy:    v.Createdby,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddKategoriMenuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddKategoriMenu",
		encodeAddKategoriMenuRequest,
		decodeKategoriMenuResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddKategoriMenu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddKategoriMenu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
func makeClientReadKategoriMenuByKategoriMenuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKategoriMenuByKategoriMenu",
		encodeReadKategoriMenuByKategoriMenuRequest,
		decodeReadKategoriMenuByKategoriMenuRespones,
		pb.ReadKategoriMenuByKategoriMenuResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKategoriMenuByKategoriMenu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKategoriMenuByKategoriMenu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKategoriMenuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKategoriMenu",
		encodeReadKategoriMenuRequest,
		decodeReadKategoriMenuResponse,
		pb.ReadKategoriMenuResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKategoriMenu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKategoriMenu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateKategoriMenu(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateKategoriMenu",
		encodeUpdateKategoriMenuRequest,
		decodeKategoriMenuResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateKategoriMenu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateKategoriMenu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
