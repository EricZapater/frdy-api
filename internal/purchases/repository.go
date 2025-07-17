package purchases

import (
	"database/sql"
	"fmt"
)

type PurchaseRepository interface {
	CreatePurchaseHeader(header PurchaseHeader) (PurchaseHeader, error)
	UpdatePurchaseHeader(header PurchaseHeader) (PurchaseHeader, error)
	FindPurchaseByID(id string) (PurchaseHeader, error)
	FindAllPurchases() ([]PurchaseHeader, error)
	DeletePurchaseByID(id string) error

	CreatePurchaseDetail(detail PurchaseDetail) (PurchaseDetail, error)
	UpdatePurchaseDetail(detail PurchaseDetail) (PurchaseDetail, error)
	FindDetailsByPurchaseID(headerID string) ([]PurchaseDetail, error)
	DeletePurchaseDetailByID(id string) error
}

type purchaseRepository struct {
	db *sql.DB
}

func NewPurchaseRepository(db *sql.DB) PurchaseRepository {
	return &purchaseRepository{db: db}
}

// Headers

func (r *purchaseRepository) CreatePurchaseHeader(header PurchaseHeader) (PurchaseHeader, error) {
	_, err := r.db.Exec(`
		INSERT INTO purchase_headers (id, code, created_at)
		VALUES ($1, $2, $3)`,
		header.ID, header.Code, header.CreatedAt,
	)
	if err != nil {
		return PurchaseHeader{}, fmt.Errorf("error inserting purchase header: %w", err)
	}
	return header, nil
}

func (r *purchaseRepository) UpdatePurchaseHeader(header PurchaseHeader) (PurchaseHeader, error) {
	_, err := r.db.Exec(`
		UPDATE purchase_headers
		SET code = $1
		WHERE id = $2`,
		header.Code, header.ID,
	)
	if err != nil {
		return PurchaseHeader{}, fmt.Errorf("error updating purchase header: %w", err)
	}
	return header, nil
}

func (r *purchaseRepository) FindPurchaseByID(id string) (PurchaseHeader, error) {
	row := r.db.QueryRow(`
		SELECT id, code, created_at
		FROM purchase_headers
		WHERE id = $1`, id)

	var header PurchaseHeader
	if err := row.Scan(&header.ID, &header.Code, &header.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return PurchaseHeader{}, fmt.Errorf("purchase header not found: %w", err)
		}
		return PurchaseHeader{}, fmt.Errorf("error scanning purchase header: %w", err)
	}
	return header, nil
}

func (r *purchaseRepository) FindAllPurchases() ([]PurchaseHeader, error) {
	rows, err := r.db.Query(`
		SELECT id, code, created_at
		FROM purchase_headers`)
	if err != nil {
		return nil, fmt.Errorf("error querying all purchase headers: %w", err)
	}
	defer rows.Close()

	var headers []PurchaseHeader
	for rows.Next() {
		var header PurchaseHeader
		if err := rows.Scan(&header.ID, &header.Code, &header.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning purchase header: %w", err)
		}
		headers = append(headers, header)
	}
	return headers, nil
}

func (r *purchaseRepository) DeletePurchaseByID(id string) error {
	_, err := r.db.Exec(`
		DELETE FROM purchase_headers
		WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting purchase header: %w", err)
	}
	return nil
}

// Details

func (r *purchaseRepository) CreatePurchaseDetail(detail PurchaseDetail) (PurchaseDetail, error) {
	_, err := r.db.Exec(`
		INSERT INTO purchase_details (id, item_id, quantity, cost, amount)
		VALUES ($1, $2, $3, $4, $5)`,
		detail.ID, detail.ItemID, detail.Quantity, detail.Cost, detail.Amount,
	)
	if err != nil {
		return PurchaseDetail{}, fmt.Errorf("error inserting purchase detail: %w", err)
	}
	return detail, nil
}

func (r *purchaseRepository) UpdatePurchaseDetail(detail PurchaseDetail) (PurchaseDetail, error) {
	_, err := r.db.Exec(`
		UPDATE purchase_details
		SET item_id = $1, quantity = $2, cost = $3, amount = $4
		WHERE id = $5`,
		detail.ItemID, detail.Quantity, detail.Cost, detail.Amount, detail.ID,
	)
	if err != nil {
		return PurchaseDetail{}, fmt.Errorf("error updating purchase detail: %w", err)
	}
	return detail, nil
}

func (r *purchaseRepository) FindDetailsByPurchaseID(headerID string) ([]PurchaseDetail, error) {
	rows, err := r.db.Query(`
		SELECT pd.id, pd.item_id, i.code AS item_code, i.description AS item_description,
		       pd.quantity, pd.cost, pd.amount
		FROM purchase_details pd
		INNER JOIN items i ON pd.item_id = i.id
		WHERE pd.id IN (
			SELECT pd.id
			FROM purchase_details pd
			JOIN purchase_headers ph ON TRUE
			WHERE ph.id = $1
		)`, headerID)
	if err != nil {
		return nil, fmt.Errorf("error querying purchase details: %w", err)
	}
	defer rows.Close()

	var details []PurchaseDetail
	for rows.Next() {
		var detail PurchaseDetail
		if err := rows.Scan(&detail.ID, &detail.ItemID, &detail.ItemCode, &detail.ItemDescription,
			&detail.Quantity, &detail.Cost, &detail.Amount); err != nil {
			return nil, fmt.Errorf("error scanning purchase detail: %w", err)
		}
		details = append(details, detail)
	}
	return details, nil
}

func (r *purchaseRepository) DeletePurchaseDetailByID(id string) error {
	_, err := r.db.Exec(`
		DELETE FROM purchase_details
		WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting purchase detail: %w", err)
	}
	return nil
}
