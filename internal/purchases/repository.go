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
	ReceivePurchaseHeader(id string) (PurchaseHeader, error)
	GetNextNumber()(string, error)

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
		INSERT INTO purchase_headers (id, code, supplier_name, created_at)
		VALUES ($1, $2, $3, $4)`,
		header.ID, header.Code,header.SupplierName, header.CreatedAt,
	)
	if err != nil {
		return PurchaseHeader{}, fmt.Errorf("error inserting purchase header: %w", err)
	}
	return header, nil
}

func (r *purchaseRepository) UpdatePurchaseHeader(header PurchaseHeader) (PurchaseHeader, error) {
	_, err := r.db.Exec(`
		UPDATE purchase_headers
		SET code = $1, supplier_name = $2
		WHERE id = $3`,
		header.Code, header.SupplierName, header.ID,
	)
	if err != nil {
		return PurchaseHeader{}, fmt.Errorf("error updating purchase header: %w", err)
	}
	return header, nil
}

func (r *purchaseRepository) FindPurchaseByID(id string) (PurchaseHeader, error) {
	row := r.db.QueryRow(`
		SELECT id, code, supplier_name,  created_at, received
		FROM purchase_headers
		WHERE id = $1`, id)

	var header PurchaseHeader
	if err := row.Scan(&header.ID, &header.Code, &header.SupplierName, &header.CreatedAt, &header.Received); err != nil {
		if err == sql.ErrNoRows {
			return PurchaseHeader{}, fmt.Errorf("purchase header not found: %w", err)
		}
		return PurchaseHeader{}, fmt.Errorf("error scanning purchase header: %w", err)
	}
	return header, nil
}

func (r *purchaseRepository) FindAllPurchases() ([]PurchaseHeader, error) {
	rows, err := r.db.Query(`
		SELECT id, code, supplier_name, created_at, received
		FROM purchase_headers`)
	if err != nil {
		return nil, fmt.Errorf("error querying all purchase headers: %w", err)
	}
	defer rows.Close()

	var headers []PurchaseHeader
	for rows.Next() {
		var header PurchaseHeader
		if err := rows.Scan(&header.ID, &header.Code, &header.SupplierName, &header.CreatedAt, &header.Received); err != nil {
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
		INSERT INTO purchase_details (id, item_id, purchase_header_id, quantity, cost, amount)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		detail.ID, detail.ItemID, detail.PurchaseHeaderID, detail.Quantity, detail.Cost, detail.Amount,
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
		WHERE pd.purchase_header_id  = $1
		`, headerID)
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

func (r *purchaseRepository) ReceivePurchaseHeader(id string) (PurchaseHeader, error) {
	_, err := r.db.Exec(`
		UPDATE purchase_headers
		SET received = TRUE
		WHERE id = $1`, id)
	if err != nil {
		return PurchaseHeader{}, fmt.Errorf("error receiving purchase header: %w", err)
	}

	header, err := r.FindPurchaseByID(id)
	if err != nil {
		return PurchaseHeader{}, fmt.Errorf("error fetching updated purchase header: %w", err)
	}
	return header, nil
}

func(r *purchaseRepository) GetNextNumber()(string, error){
	var nextCounter string
	err := r.db.QueryRow(`
	SELECT   
		REPEAT(
			'0',
			10 - LENGTH(CAST(COALESCE(MAX(CAST(code AS integer)), 0) + 1 AS varchar))
		) || CAST(COALESCE(MAX(CAST(code AS integer)), 0) + 1 AS varchar) AS next_counter
		FROM purchase_headers
	`).Scan(&nextCounter)
	if err != nil {
		if err == sql.ErrNoRows {
			return "0000000001", nil
		}
		return "", err
	}
	return nextCounter, nil
}