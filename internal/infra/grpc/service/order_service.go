package service

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase   usecase.CreateOrderUseCase
	ListAllOrdersUseCase usecase.ListAllOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listAllOrdersUseCase usecase.ListAllOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase:   createOrderUseCase,
		ListAllOrdersUseCase: listAllOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrder(ctx context.Context, in *pb.Blank) (*pb.ListOrderRequest, error) {
	dto, err := s.ListAllOrdersUseCase.ExecuteList()
	if err != nil {
		return nil, err
	}
	var orders []*pb.CreateOrderResponse
	for _, d := range dto {
		orders = append(orders, &pb.CreateOrderResponse{

			Id:         d.ID,
			Price:      float32(d.Price),
			Tax:        float32(d.Tax),
			FinalPrice: float32(d.FinalPrice),
		})

	}
	return &pb.ListOrderRequest{Orders: orders}, nil

}
