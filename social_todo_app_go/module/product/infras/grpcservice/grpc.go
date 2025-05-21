package grpcservice

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"google.golang.org/grpc"
	"log"
	"net"
	"social_todo_app_go/module/product/query"
	"social_todo_app_go/proto/product"
)

type service struct {
	port int
	sctx sctx.ServiceContext
}

func NewProductGRPCService(port int, sctx sctx.ServiceContext) *service {
	return &service{port: port, sctx: sctx}
}

type productServer struct {
	product.UnimplementedProductServer
	sctx sctx.ServiceContext
}

func newProductServer(sctx sctx.ServiceContext) *productServer {
	return &productServer{sctx: sctx}
}

func (s *productServer) GetProductById(ctx context.Context, request *product.GetProductByIdRequest) (*product.GetProductByIdResponse, error) {
	ids := make([]uuid.UUID, len(request.Ids))
	for i := range request.Ids {
		ids[i] = uuid.MustParse(request.Ids[i])
	}
	products, err := query.NewGetProductByIdQuery(s.sctx).Execute(ctx, ids)
	if err != nil {
		return nil, err
	}
	result := make([]*product.ProductDTO, len(products))
	for i := range products {
		result[i] = &product.ProductDTO{
			Id:   products[i].Id.String(),
			Name: products[i].Name,
		}
	}
	return &product.GetProductByIdResponse{
		Data: result,
	}, nil
}

func (s *productServer) mustEmbedUnimplementedProductServer() {
	//TODO implement me
	panic("implement me")
}

func (s *service) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create new gRPC server
	grpcServer := grpc.NewServer()
	product.RegisterProductServer(grpcServer, newProductServer(s.sctx))
	log.Println("grpc server is running on port:", s.port)
	log.Fatal(grpcServer.Serve(lis))
	return nil
}
