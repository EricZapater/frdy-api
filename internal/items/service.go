package items

import (
	"errors"

	"github.com/google/uuid"
)

type ItemService interface {
	Create(item ItemRequest) (Item, error)
	Update(id string, item ItemRequest) (Item, error)
	Delete(id string) error
	FindByID(id string) (Item, error)
	FindByCode(code string) (Item, error)
	FindAll() ([]Item, error)
}

type itemService struct {
	repo ItemRepository
}

func NewItemService(repo ItemRepository) ItemService {
	return &itemService{repo: repo}
}

func (s *itemService) Create(item ItemRequest) (Item, error) {
	if item.Code == "" || item.Description == "" || item.Cost <= 0 || item.Price <= 0 {
		return Item{}, errors.New("invalid request")
	}

	reference := Item{
		ID:          uuid.New(),
		Code:        item.Code,
		Description: item.Description,
		Cost:        item.Cost,
		Price:       item.Price,
		IsActive:    true, // Default to active
	}

	return s.repo.Create(reference)
}

func (s *itemService) Update(id string, item ItemRequest) (Item, error) {
	if item.Code == "" || item.Description == "" || item.Cost <= 0 || item.Price <= 0 {
		return Item{}, errors.New("invalid request")
	}

	referenceID, err := uuid.Parse(id)
	if err != nil {
		return Item{}, errors.New("invalid ID format")
	}

	reference := Item{
		ID:          referenceID,
		Code:        item.Code,
		Description: item.Description,
		Cost:        item.Cost,
		Price:       item.Price,
		IsActive:    true, // Default to active
	}

	return s.repo.Update(reference)
}

func (s *itemService) Delete(id string) error {
	referenceID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	return s.repo.Delete(referenceID)
}

func (s *itemService) FindByID(id string) (Item, error) {
	referenceID, err := uuid.Parse(id)
	if err != nil {
		return Item{}, errors.New("invalid ID format")
	}

	return s.repo.FindByID(referenceID)
}

func (s *itemService) FindByCode(code string) (Item, error) {
	if code == "" {
		return Item{}, errors.New("code cannot be empty")
	}

	return s.repo.FindByCode(code)
}

func (s *itemService) FindAll() ([]Item, error) {
	return s.repo.FindAll()
}