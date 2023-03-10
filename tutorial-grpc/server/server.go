package server

import (
	"context"
	"sync"
	v1 "tutorialgrpc/gen/product/v1"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	mu       sync.Mutex
	products map[string]v1.Product
	v1.UnimplementedProductServiceServer
}

func New() *Server {
	return &Server{
		products: make(map[string]v1.Product),
	}
}

func (s *Server) AddProduct(ctx context.Context, req *v1.AddProductRequest) (*v1.AddProductResponse, error) {
	p := v1.Product{
		Id:    uuid.New().String(),
		Name:  req.Name,
		Price: req.Price,
	}

	s.mu.Lock()
	s.products[p.Id] = p
	s.mu.Unlock()

	return &v1.AddProductResponse{
		ProductId: p.Id,
	}, nil
}

func (s *Server) DeleteProduct(ctx context.Context, req *v1.DeleteProductRequest) (*v1.DeleteProductResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p, ok := s.products[req.ProductId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "product with ID: %q does not exist", req.ProductId)
	}

	delete(s.products, req.ProductId)

	return &v1.DeleteProductResponse{
		Product: &p,
	}, nil
}

func (s *Server) ListProducts(ctx context.Context, req *v1.ListProductsRequest) (*v1.ListProductsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ps := make([]*v1.Product, 0, len(s.products))
	for id := range s.products {
		p := s.products[id]
		ps = append(ps, &p)
	}

	return &v1.ListProductsResponse{
		Products: ps,
	}, nil

}
