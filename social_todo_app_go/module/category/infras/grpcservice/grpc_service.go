package grpcservice

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"google.golang.org/grpc"
	"log"
	"net"
	"social_todo_app_go/module/category/query"
	"social_todo_app_go/proto/category"
)

type service struct {
	port int
	sctx sctx.ServiceContext
}

func NewCategoryGRPCService(port int, sctx sctx.ServiceContext) *service {
	return &service{port: port, sctx: sctx}
}

type categoryServer struct {
	category.UnimplementedCategoryServer
	sctx sctx.ServiceContext
}

func newCategoryServer(sctx sctx.ServiceContext) *categoryServer {
	return &categoryServer{sctx: sctx}
}

func (s categoryServer) GetCategoryById(ctx context.Context, request *category.GetCategoryByIdRequest) (*category.GetCategoryByIdResponse, error) {
	var categories []query.CategoryDTO

	ids := make([]uuid.UUID, len(request.Ids))
	for i := range request.Ids {
		ids[i] = uuid.MustParse(request.Ids[i])
	}

	categories, err := query.NewCategoryByIdQuery(s.sctx).Execute(ctx, ids)
	if err != nil {
		return nil, err
	}

	result := make([]*category.CategoryDTO, len(categories))

	for i := range categories {
		result[i] = &category.CategoryDTO{
			Id:   categories[i].Id,
			Name: categories[i].Name,
		}
	}
	
	return &category.GetCategoryByIdResponse{Data: result}, nil
}

func (categoryServer) mustEmbedUnimplementedCategoryServer() {
	//TODO implement me
	panic("implement me")
}

func (s *service) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	category.RegisterCategoryServer(grpcServer, newCategoryServer(s.sctx))
	log.Println(fmt.Sprintf("gRPC server is running on port %d", s.port))
	log.Fatal(grpcServer.Serve(lis))
	return nil
}
