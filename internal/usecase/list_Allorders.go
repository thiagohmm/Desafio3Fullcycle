package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListAllOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func ListAllOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,

) *ListAllOrdersUseCase {
	return &ListAllOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListAllOrdersUseCase) ExecuteList() ([]*OrderOutputListByDTO, error) {

	dto, err := l.OrderRepository.GetAll()

	if err != nil {
		return []*OrderOutputListByDTO{}, err
	}
	var orders []*OrderOutputListByDTO
	for _, d := range dto {
		orders = append(orders, &OrderOutputListByDTO{
			ID:         d.ID,
			Price:      d.Price,
			Tax:        d.Tax,
			FinalPrice: d.FinalPrice,
		})
	}
	return orders, nil

}
