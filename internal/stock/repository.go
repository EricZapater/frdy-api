package stock

import (
	"database/sql"
	"fmt"
)

type StockRepository interface {
	GetStockByItemID(itemID string) (*Stock, error)
	UpdateStockQuantity(itemID string, quantity int) error
	GetAllStocks() ([]Stock, error)
}

type stockRepository struct {
	db *sql.DB // Assuming DB is a struct that handles database operations
}

func NewStockRepository(db *sql.DB) StockRepository {
	return &stockRepository{db: db}
}

func (r *stockRepository) GetStockByItemID(itemID string) (*Stock, error) {
	var stock Stock
	err := r.db.QueryRow(`
		SELECT s.id, s.item_id, i.code as item_code,i.description as item_description, s.quantity
		FROM stocks s
			INNER JOIN items i ON s.item_id = i.id
		WHERE i.item_id = $1`, itemID).Scan(&stock.ID, &stock.ItemID, &stock.ItemCode, &stock.ItemDescription, &stock.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No stock found for this item
		}
		return nil, fmt.Errorf("error fetching stock: %w", err)
	}
	return &stock, nil
}

func (r *stockRepository) UpdateStockQuantity(itemID string, quantity int) error {
	_, err := r.db.Exec(`
		INSERT INTO stocks (item_id, quantity)
		VALUES ($1, $2)
			ON CONFLICT (item_id) DO UPDATE SET
    		quantity = stocks.quantity + EXCLUDED.quantity;`, itemID, quantity)
	if err != nil {
		return fmt.Errorf("error updating stock quantity: %w", err)
	}
	return nil
}

func (r *stockRepository) GetAllStocks() ([]Stock, error) {
	var stocks []Stock
	rows, err := r.db.Query(`
		SELECT s.id, s.item_id, i.code as item_code, i.description as item_description, s.quantity
		FROM stocks s
			INNER JOIN items i ON s.item_id = i.id`)
	if err != nil {
		return nil, fmt.Errorf("error fetching all stocks: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var stock Stock
		if err := rows.Scan(&stock.ID, &stock.ItemID, &stock.ItemCode, &stock.ItemDescription, &stock.Quantity); err != nil {
			return nil, fmt.Errorf("error scanning stock row: %w", err)
		}
		stocks = append(stocks, stock)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over stock rows: %w", err)
	}

	return stocks, nil
}