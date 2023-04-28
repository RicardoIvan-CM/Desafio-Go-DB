package challenges

import "github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"

type Service interface {
	GetTotalsByCustomerCondition() (domain.TotalsByCustomerCondition, error)
	GetTopSoldProducts() ([]domain.TopSoldProduct, error)
	GetTopActiveCustomersSpent() ([]domain.TopActiveCustomerSpent, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetTotalsByCustomerCondition() (domain.TotalsByCustomerCondition, error) {
	result, err := s.r.GetTotalsByCustomerCondition()
	if err != nil {
		return domain.TotalsByCustomerCondition{}, err
	}
	return result, nil
}

func (s *service) GetTopSoldProducts() ([]domain.TopSoldProduct, error) {
	result, err := s.r.GetTopSoldProducts()
	if err != nil {
		return []domain.TopSoldProduct{}, err
	}
	return result, nil
}

func (s *service) GetTopActiveCustomersSpent() ([]domain.TopActiveCustomerSpent, error) {
	result, err := s.r.GetTopActiveCustomersSpent()
	if err != nil {
		return []domain.TopActiveCustomerSpent{}, err
	}
	return result, nil
}
