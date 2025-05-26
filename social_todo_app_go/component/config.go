package component

import (
	"flag"
	sctx "github.com/linhhonblade/service-context"
)

type config struct {
	id                        string
	urlRPCCategory            string
	portGRPCCategory          int
	urlGRPCCategoryServer     string
	defaultDraftOrderDuration float64
	portGRPCProduct           int
	urlGRPCProductServer      string
	appDomain                 string
	appEnv                    string
}

func NewConfig(id string) *config {
	return &config{id: id}
}

func (c *config) ID() string {
	return c.id
}

func (c *config) InitFlags() {
	flag.StringVar(
		&c.urlRPCCategory,
		"rpc-category-url",
		"http://localhost:3000/v1/category/rpc",
		"URL of category RPC",
	)
	flag.IntVar(
		&c.portGRPCCategory,
		"grpc-category-port",
		8000,
		"Port of category gRPC",
	)
	flag.StringVar(
		&c.urlGRPCCategoryServer,
		"grpc-category-url",
		":8000",
		"URL of category gRPC",
	)
	flag.Float64Var(
		&c.defaultDraftOrderDuration,
		"default-draft-order-duration",
		1,
		"Default draft order duration (in hour)",
	)
	flag.IntVar(
		&c.portGRPCProduct,
		"grpc-product-port",
		8001,
		"Port of product gRPC")
	flag.StringVar(
		&c.urlGRPCProductServer,
		"grpc-product-url",
		":8001",
		"URL of product gRPC")
	flag.StringVar(
		&c.appDomain,
		"app-domain",
		"localhost",
		"Domain of the application, used for set cookie")
	flag.StringVar(
		&c.appEnv,
		"env",
		"dev",
		"Environment of the application, used for set cookie")
}

func (c *config) Activate(context sctx.ServiceContext) error {
	return nil
}

func (c *config) Stop() error {
	return nil
}

func (c config) GetUrlRPCCategory() string {
	return c.urlRPCCategory
}

func (c config) GetPortGRPCCategory() int {
	return c.portGRPCCategory
}

func (c config) GetUrlGRPCCategoryServer() string {
	return c.urlGRPCCategoryServer
}

func (c config) GetPortGRPCProduct() int {
	return c.portGRPCProduct
}

func (c config) GetUrlGRPCProductServer() string {
	return c.urlGRPCProductServer
}

func (c config) GetAppDomain() string {
	return c.appDomain
}

func (c config) GetAppEnv() string {
	return c.appEnv
}
