package rpc

import (
	"sync"

	"gomall/app/checkout/conf"
	"gomall/rpc_gen/kitex_gen/cart/cartservice"

	//"gomall/rpc_gen/kitex_gen/order/orderservice"
	"gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"gomall/rpc_gen/kitex_gen/product/productcatalogservice"

	"gomall/rpc_gen/kitex_gen/order/orderservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	err           error
	registryAddr  string
	serviceName   string
	commonSuite   client.Option
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

// func initProductClient() {
// 	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
// 	checkoututils.MustHandleError(err)
// }

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}
}
func initPaymentClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}
}
func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}
}