package rpc

import (
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/zxq97/gotool/config"
	"github.com/zxq97/gotool/constant"
	"github.com/zxq97/gotool/discover"
	"github.com/zxq97/gotool/register"
	"go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

func NewGrpcConn(etcdClient *clientv3.Client, svcName string, hystrixConf *config.HystrixConf) (*grpc.ClientConn, error) {
	ed := discover.NewEtcdDiscover(etcdClient, svcName)
	resolver.Register(ed)
	initBreaker(hystrixConf)
	conn, err := grpc.Dial(
		ed.Scheme()+":///",
		grpc.WithInsecure(),
		grpc.WithResolvers(ed),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				timeout(constant.DefaultTimeout),
				demote(svcName),
		)))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewGrpcServer(etcdClient *clientv3.Client, svcKey, port string) (*grpc.Server, *register.EtcdRegister) {
	er := register.NewEtcdRegister(etcdClient, svcKey, port)
	opt := []grpc.ServerOption{
		grpc.UnaryInterceptor(recovery),
	}
	svc := grpc.NewServer(opt...)
	return svc, er
}
