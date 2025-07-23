package purchases

import (
	"errors"
	"frdy-api/internal/stock"
	"time"

	"github.com/google/uuid"
)

type PurchaseService interface {
	CreatePurchaseHeader(request PurchaseHeaderRequest) (PurchaseHeader, error)
	UpdatePurchaseHeader(id string, request PurchaseHeaderRequest) (PurchaseHeader, error)
	FindPurchaseByID(id string) (PurchaseHeader, error)
	FindAllPurchases() ([]PurchaseHeader, error)
	DeletePurchaseByID(id string) error
	ReceivePurchaseHeader(id string) (PurchaseHeader, error)

	CreatePurchaseDetail(request PurchaseDetailRequest) (PurchaseDetail, error)
	UpdatePurchaseDetail(id string, request PurchaseDetailRequest) (PurchaseDetail, error)
	FindDetailsByPurchaseID(headerID string) ([]PurchaseDetail, error)
	DeletePurchaseDetailByID(id string) error
}

type purchaseService struct {
	repo PurchaseRepository
	stock stock.StockService
}

func NewPurchaseService(repo PurchaseRepository, stock stock.StockService) PurchaseService {
	return &purchaseService{repo: repo, stock: stock}
}

// Header methods

func (s *purchaseService) CreatePurchaseHeader(request PurchaseHeaderRequest) (PurchaseHeader, error) {
	next_counter, err := s.repo.GetNextNumber()
	if err != nil {
		return PurchaseHeader{}, errors.New("invalid counter")
	}

	header := PurchaseHeader{
		ID:        uuid.New().String(),
		Code:      next_counter,
		SupplierName: request.SupplierName,
		CreatedAt: time.Now(),
	}

	return s.repo.CreatePurchaseHeader(header)
}

func (s *purchaseService) UpdatePurchaseHeader(id string, request PurchaseHeaderRequest) (PurchaseHeader, error) {
	if request.Code == "" {
		return PurchaseHeader{}, errors.New("code is required")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return PurchaseHeader{}, errors.New("invalid ID format")
	}

	header := PurchaseHeader{
		ID:        id,
		Code:      request.Code,
		CreatedAt: request.CreatedAt, // podries ignorar si no cal actualitzar-lo
	}

	return s.repo.UpdatePurchaseHeader(header)
}

func (s *purchaseService) FindPurchaseByID(id string) (PurchaseHeader, error) {
	if id == "" {
		return PurchaseHeader{}, errors.New("id is required")
	}
	return s.repo.FindPurchaseByID(id)
}

func (s *purchaseService) FindAllPurchases() ([]PurchaseHeader, error) {
	return s.repo.FindAllPurchases()
}

func (s *purchaseService) DeletePurchaseByID(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return s.repo.DeletePurchaseByID(id)
}

// Detail methods

func (s *purchaseService) CreatePurchaseDetail(request PurchaseDetailRequest) (PurchaseDetail, error) {
	if request.ItemID == "" || request.Quantity <= 0 || request.Cost <= 0 {
		return PurchaseDetail{}, errors.New("invalid purchase detail request")
	}

	detail := PurchaseDetail{
		ID:              uuid.New().String(),
		PurchaseHeaderID: request.PurchaseHeaderID,
		ItemID:          request.ItemID,
		Quantity:        request.Quantity,
		Cost:            request.Cost,
		Amount:          float64(request.Quantity) * request.Cost,
	}

	return s.repo.CreatePurchaseDetail(detail)
}

func (s *purchaseService) UpdatePurchaseDetail(id string, request PurchaseDetailRequest) (PurchaseDetail, error) {
	if id == "" || request.ItemID == "" || request.Quantity <= 0 || request.Cost <= 0 {
		return PurchaseDetail{}, errors.New("invalid purchase detail request")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return PurchaseDetail{}, errors.New("invalid ID format")
	}

	detail := PurchaseDetail{
		ID:       id,
		ItemID:   request.ItemID,
		Quantity: request.Quantity,
		Cost:     request.Cost,
		Amount:   float64(request.Quantity) * request.Cost,
	}

	return s.repo.UpdatePurchaseDetail(detail)
}

func (s *purchaseService) FindDetailsByPurchaseID(headerID string) ([]PurchaseDetail, error) {
	if headerID == "" {
		return nil, errors.New("header ID is required")
	}
	return s.repo.FindDetailsByPurchaseID(headerID)
}

func (s *purchaseService) DeletePurchaseDetailByID(id string) error {
	if id == "" {
		return errors.New("detail ID is required")
	}
	return s.repo.DeletePurchaseDetailByID(id)
}

func (s *purchaseService) ReceivePurchaseHeader(id string) (PurchaseHeader, error) {
	if id == "" {
		return PurchaseHeader{}, errors.New("id is required")
	}

	header, err := s.repo.ReceivePurchaseHeader(id)
	if err != nil {
		return PurchaseHeader{}, err
	}
	details, err := s.repo.FindDetailsByPurchaseID(header.ID)
	if err != nil {
		return PurchaseHeader{}, err
	}
	for _, detail := range details {
		if err := s.stock.UpdateStockQuantity(detail.ItemID, detail.Quantity); err != nil {
			return PurchaseHeader{}, err
		}
	}

	return header, nil
}