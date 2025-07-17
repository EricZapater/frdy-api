package stock

import "errors"

type StockService interface {
	GetStockByItemID(itemID string) (Stock, error)
	UpdateStockQuantity(itemID string, quantity int) error
	GetAllStocks() ([]Stock, error)
}

type stockService struct {
	repo StockRepository
}

func NewStockService(repo StockRepository) StockService {
	return &stockService{repo: repo}
}

func (s *stockService) GetStockByItemID(itemID string) (Stock, error) {
	stock, err := s.repo.GetStockByItemID(itemID)
	if err != nil {
		return Stock{}, err
	}
	if stock == nil {
		return Stock{}, errors.New("stock not found")
	}
	return *stock, nil
}
func (s *stockService) UpdateStockQuantity(itemID string, quantity int) error {
	if itemID == "" || quantity < 0 {
		return errors.New("invalid item ID or quantity")
	}
	return s.repo.UpdateStockQuantity(itemID, quantity)
}
func (s *stockService) GetAllStocks() ([]Stock, error) {
	stocks, err := s.repo.GetAllStocks()
	if err != nil {
		return nil, err
	}
	return stocks, nil
}