package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type OrderInputListByIdDTO struct {
	ID string `json:"id"`
}

type OrderOutputListByDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderByIdUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
}

func ListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	//OrderCreated events.EventInterface,
	//EventDispatcher events.EventDispatcherInterface,
) *ListOrderByIdUseCase {
	return &ListOrderByIdUseCase{
		OrderRepository: OrderRepository,
		//OrderCreated:    OrderCreated,
		//EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrderByIdUseCase) ExecuteListById(input OrderInputListByIdDTO) (OrderOutputListByDTO, error) {
	order := entity.Order{
		ID: input.ID,
	}
	dto, err := l.OrderRepository.ListById(order.ID)
	if err != nil {
		return OrderOutputListByDTO{}, err
	}

	return OrderOutputListByDTO{
		ID:         dto.ID,
		Price:      dto.Price,
		Tax:        dto.Tax,
		FinalPrice: dto.FinalPrice,
	}, nil
}
