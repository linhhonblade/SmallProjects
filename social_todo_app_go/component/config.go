package component

import (
	"flag"
	sctx "github.com/linhhonblade/service-context"
)

type config struct {
	id                    string
	urlRPCCategory        string
	portGRPCCategory      int
	urlGRPCCategoryServer string
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
