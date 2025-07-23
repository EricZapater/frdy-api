package sales

import (
	"errors"
	"frdy-api/internal/stock"
	"time"

	"github.com/google/uuid"
)

type SalesService interface {
	CreateSalesHeader(request SalesHeaderRequest) (SalesHeader, error)
	UpdateSalesHeader(id string, request SalesHeaderRequest) (SalesHeader, error)
	FindSalesByHeaderID(id string) (SalesHeader, error)
	FindSalesByHeaderCode(code string) (SalesHeader, error)
	FindSalesByItemCode(itemCode string) ([]SalesHeader, error)
	FindSalesByCustomerName(customerName string) ([]SalesHeader, error)
	FindAllSales() ([]SalesHeader, error)
	DeleteSalesByHeaderID(id string) error
	CreateSalesDetail(request SalesDetailRequest) (SalesDetail, error)
	UpdateSalesDetail(id string, request SalesDetailRequest) (SalesDetail, error)
	FindSalesDetailsByHeaderID(headerID string) ([]SalesDetail, error)
	DeleteSalesDetailByID(id string) error
	SendSalesHeader(id string) (SalesHeader, error)
}

type salesService struct {
	repo SalesRepository
	stock stock.StockService
}

func NewSalesService(repo SalesRepository, stock stock.StockService) SalesService {
	return &salesService{repo: repo, stock:stock}
}

func (s *salesService) CreateSalesHeader(request SalesHeaderRequest) (SalesHeader, error) {
	if  request.CustomerName == "" {
		return SalesHeader{}, errors.New("invalid request")
	}

	counter, err := s.repo.GetNextNumber()
	if err != nil {
		return SalesHeader{}, errors.New("cannot get counter")
	}

	header := SalesHeader{
		ID:           uuid.New(),
		Code:         counter,
		CustomerName: request.CustomerName,
		CustomerPhone: request.CustomerPhone,
		CreatedAt:    time.Now().Format(time.RFC3339),
		Sent:         false, // Default to not sent
	}

	return s.repo.CreateSalesHeader(header)
}

func (s *salesService) UpdateSalesHeader(id string, request SalesHeaderRequest) (SalesHeader, error) {
	if request.Code == "" || request.CustomerName == "" {
		return SalesHeader{}, errors.New("invalid request")
	}

	headerID, err := uuid.Parse(id)
	if err != nil {
		return SalesHeader{}, errors.New("invalid ID format")
	}

	header := SalesHeader{
		ID:           headerID,
		Code:         request.Code,
		CustomerName: request.CustomerName,
		CustomerPhone: request.CustomerPhone,
	}

	return s.repo.UpdateSalesHeader(header)
}

func (s *salesService) FindSalesByHeaderID(id string) (SalesHeader, error) {
	header, err := s.repo.FindSalesByHeaderID(id)
	if err != nil {
		return SalesHeader{}, err
	}
	return header, nil
}

func (s *salesService) FindSalesByHeaderCode(code string) (SalesHeader, error) {
	header, err := s.repo.FindSalesByHeaderCode(code)
	if err != nil {
		return SalesHeader{}, err
	}
	return header, nil
}

func (s *salesService) FindSalesByItemCode(itemCode string) ([]SalesHeader, error) {
	headers, err := s.repo.FindSalesByItemCode(itemCode)
	if err != nil {
		return nil, err
	}
	return headers, nil
}

func (s *salesService) FindSalesByCustomerName(customerName string) ([]SalesHeader, error) {
	headers, err := s.repo.FindSalesByCustomerName(customerName)
	if err != nil {
		return nil, err
	}
	return headers, nil
}

func (s *salesService) FindAllSales() ([]SalesHeader, error) {
	headers, err := s.repo.FindAllSales()
	if err != nil {
		return nil, err
	}
	return headers, nil
}

func (s *salesService) DeleteSalesByHeaderID(id string) error {
	if id == "" {
		return errors.New("invalid ID")
	}

	return s.repo.DeleteSalesByHeaderID(id)
}

func (s *salesService) CreateSalesDetail(request SalesDetailRequest) (SalesDetail, error) {
	if request.SalesHeaderID == "" || request.ItemID == "" || request.Quantity <= 0 || request.Price <= 0 {
		return SalesDetail{}, errors.New("invalid request")
	}

	detail := SalesDetail{
		ID:              uuid.New(),
		SalesHeaderID:   request.SalesHeaderID,
		ItemID:          request.ItemID,
		Quantity:        request.Quantity,
		Price:           request.Price,
		Amount:          float64(request.Quantity) * request.Price, // Calculate amount
	}

	return s.repo.CreateSalesDetail(detail)
}

func (s *salesService) UpdateSalesDetail(id string, request SalesDetailRequest) (SalesDetail, error) {
	if request.SalesHeaderID == "" || request.ItemID == "" || request.Quantity <= 0 || request.Price <= 0 {
		return SalesDetail{}, errors.New("invalid request")
	}

	detailID, err := uuid.Parse(id)
	if err != nil {
		return SalesDetail{}, errors.New("invalid ID format")
	}

	detail := SalesDetail{
		ID:              detailID,
		SalesHeaderID:   request.SalesHeaderID,
		ItemID:          request.ItemID,
		Quantity:        request.Quantity,
		Price:           request.Price,
		Amount:          float64(request.Quantity) * request.Price, // Calculate amount
	}

	return s.repo.UpdateSalesDetail(detail)
}

func (s *salesService) FindSalesDetailsByHeaderID(headerID string) ([]SalesDetail, error) {
	if headerID == "" {
		return nil, errors.New("invalid header ID")
	}

	details, err := s.repo.FindSalesDetailsByHeaderID(headerID)
	if err != nil {
		return nil, err
	}
	return details, nil
}
func (s *salesService) DeleteSalesDetailByID(id string) error {
	if id == "" {
		return errors.New("invalid ID")
	}

	return s.repo.DeleteSalesDetailByID(id)
}

func (s *salesService) SendSalesHeader(id string) (SalesHeader, error) {
	if id == "" {
		return SalesHeader{}, errors.New("invalid ID")
	}

	header, err := s.repo.SendSalesHeader(id)
	if err != nil {
		return SalesHeader{}, err
	}
	details, err := s.repo.FindSalesDetailsByHeaderID(id)
	if err != nil {
		return SalesHeader{}, err
	}
	for _, detail := range details {
		if err := s.stock.UpdateStockQuantity(detail.ItemID, -detail.Quantity); err != nil {
			return SalesHeader{}, err
		}
	}

	return header, nil
}