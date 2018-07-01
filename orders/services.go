package orders

import (
	"github.com/dgraph-io/dgo"
	"github.com/kataras/iris"
)

//Service defines the service logics of orders
type Service struct {
	repo *Repository
}

//NewOrderService Create new JobService instance
func NewOrderService(dg *dgo.Dgraph) *Service {
	service := &Service{}
	service.repo = &Repository{dg}
	return service
}

//Create will create an order for given json
func (s *Service) Create(ctx iris.Context, order Order) *Order {
	order.CreatedBy = 2
	assigned, err := s.repo.CreateOrder(order)
	if err != nil {

	}

	newOrder, err := s.repo.GetOrder(assigned.Uids["blank-0"])
	if err != nil {

	}
	return newOrder
}
