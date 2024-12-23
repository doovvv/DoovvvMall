package rpc

import (
	"sync"

	"gomall/app/frontend/conf"
	frontendUtils "gomall/app/frontend/utils"
	cartservice "gomall/rpc_gen/kitex_gen/cart/cartservice"
	"gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	productcatalogservice "gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	usersrvice "gomall/rpc_gen/kitex_gen/user/userservice"

	"gomall/rpc_gen/kitex_gen/order/orderservice"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient     usersrvice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}
func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MUstHandleError(err)
	opts = append(opts, client.WithResolver(r))

	UserClient, err = usersrvice.NewClient("user", opts...)
	frontendUtils.MUstHandleError(err)
}
func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MUstHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MUstHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MUstHandleError(err)
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MUstHandleError(err)
}
func initCheckoutClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MUstHandleError(err)
	opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	frontendUtils.MUstHandleError(err)
}
func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MUstHandleError(err)
	opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order", opts...)
	frontendUtils.MUstHandleError(err)
}
