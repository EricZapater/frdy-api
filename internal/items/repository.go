package items

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type ItemRepository interface {
	Create(reference Item) (Item, error)
	Update(reference Item) (Item, error)
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (Item, error)
	FindByCode(code string) (Item, error)
	FindAll() ([]Item, error)
}

type itemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) Create(item Item) (Item, error) {
	_, err := r.db.Exec(`
		INSERT INTO items (id, code, description, cost, price, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		item.ID, item.Code, item.Description, item.Cost, item.Price, item.IsActive,
	)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func (r *itemRepository) Update(item Item) (Item, error) {
	_, err := r.db.Exec(`
		UPDATE items
		SET code = $1, description = $2, cost = $3, price = $4, is_active = $5
		WHERE id = $6`,
		item.Code, item.Description, item.Cost, item.Price, item.IsActive, item.ID,
	)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func (r *itemRepository) Delete(id uuid.UUID) error {
	_, err := r.db.Exec(`
		DELETE FROM items
		WHERE id = $1`, id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *itemRepository) FindByID(id uuid.UUID) (Item, error) {
	var item Item
	err := r.db.QueryRow(`
		SELECT id, code, description, cost, price, is_active
		FROM items
		WHERE id = $1`, id,
	).Scan(&item.ID, &item.Code, &item.Description, &item.Cost, &item.Price, &item.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return Item{}, fmt.Errorf("reference not found: %w", err)
		}
		return Item{}, err
	}
	return item, nil
}

func (r *itemRepository) FindByCode(code string) (Item, error) {
	var item Item
	err := r.db.QueryRow(`
		SELECT id, code, description, cost, price, is_active
		FROM items
		WHERE code = $1`, code,
	).Scan(&item.ID, &item.Code, &item.Description, &item.Cost, &item.Price, &item.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return Item{}, fmt.Errorf("reference not found: %w", err)
		}
		return Item{}, err
	}
	return item, nil
}

func (r *itemRepository) FindAll() ([]Item, error) {
	rows, err := r.db.Query(`
		SELECT id, code, description, cost, price, is_active
		FROM items`)
		
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Code, &item.Description, &item.Cost, &item.Price, &item.IsActive); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}