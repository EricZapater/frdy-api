package sales

import (
	"database/sql"
	"fmt"
)

type SalesRepository interface {
	CreateSalesHeader(header SalesHeader) (SalesHeader, error)
	UpdateSalesHeader(header SalesHeader) (SalesHeader, error)
	FindSalesByHeaderID(id string) (SalesHeader, error)
	FindSalesByHeaderCode(code string) (SalesHeader, error)
	FindSalesByItemCode(itemCode string) ([]SalesHeader, error)
	FindSalesByCustomerName(customerName string) ([]SalesHeader, error)
	FindAllSales() ([]SalesHeader, error)
	DeleteSalesByHeaderID(id string) error
	CreateSalesDetail(detail SalesDetail) (SalesDetail, error)
	UpdateSalesDetail(detail SalesDetail) (SalesDetail, error)
	FindSalesDetailsByHeaderID(headerID string) ([]SalesDetail, error)
	DeleteSalesDetailByID(id string) error
	SendSalesHeader(id string) (SalesHeader, error)
	GetNextNumber()(string, error)
}

type salesRepository struct {
	db *sql.DB
}

func NewSalesRepository(db *sql.DB) SalesRepository {
	return &salesRepository{db: db}
}
func (r *salesRepository) CreateSalesHeader(header SalesHeader) (SalesHeader, error) {
	_, err := r.db.Exec(`
		INSERT INTO sales_headers (id, code, customer_name, customer_phone, created_at, sent)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		header.ID, header.Code, header.CustomerName, header.CustomerPhone, header.CreatedAt, header.Sent,
	)
	if err != nil {
		return SalesHeader{}, fmt.Errorf("error inserting sales header: %w", err)
	}
	return header, nil
}

func (r *salesRepository) UpdateSalesHeader(header SalesHeader) (SalesHeader, error) {
	_, err := r.db.Exec(`
		UPDATE sales_headers
		SET code = $1, customer_name = $2, customer_phone= $3, sent = $4
		WHERE id = $5`,
		header.Code, header.CustomerName,header.CustomerPhone, header.Sent, header.ID,
	)
	if err != nil {
		return SalesHeader{}, fmt.Errorf("error updating sales header: %w", err)
	}
	return header, nil
}
func (r *salesRepository) FindSalesByHeaderID(id string) (SalesHeader, error) {
	row := r.db.QueryRow(`
		SELECT id, code, customer_name,COALESCE(customer_phone,'') as customer_phone, created_at, sent
		FROM sales_headers
		WHERE id = $1`, id)

	var header SalesHeader
	if err := row.Scan(&header.ID, &header.Code, &header.CustomerName,&header.CustomerPhone, &header.CreatedAt, &header.Sent); err != nil {
		if err == sql.ErrNoRows {
			return SalesHeader{}, fmt.Errorf("sales header not found: %w", err)
		}
		return SalesHeader{}, fmt.Errorf("error scanning sales header: %w", err)
	}
	return header, nil
}
func (r *salesRepository) FindSalesByHeaderCode(code string) (SalesHeader, error) {
	row := r.db.QueryRow(`
		SELECT id, code, customer_name, COALESCE(customer_phone,'') as customer_phone, created_at, sent
		FROM sales_headers
		WHERE code = $1`, code)

	var header SalesHeader
	if err := row.Scan(&header.ID, &header.Code, &header.CustomerName, &header.CustomerPhone, &header.CreatedAt, &header.Sent); err != nil {
		if err == sql.ErrNoRows {
			return SalesHeader{}, fmt.Errorf("sales header not found: %w", err)
		}
		return SalesHeader{}, fmt.Errorf("error scanning sales header: %w", err)
	}
	return header, nil
}
func (r *salesRepository) FindSalesByItemCode(itemCode string) ([]SalesHeader, error) {
	rows, err := r.db.Query(`
		SELECT sh.id, sh.code, sh.customer_name, COALESCE(sh.customer_phone,'') as customer_phone, sh.created_at, sh.sent
		FROM sales_headers sh
		JOIN sales_details sd ON sh.id = sd.sales_header_id
		WHERE sd.item_code = $1`, itemCode)
	if err != nil {
		return nil, fmt.Errorf("error querying sales by item code: %w", err)
	}
	defer rows.Close()

	var headers []SalesHeader
	for rows.Next() {
		var header SalesHeader
		if err := rows.Scan(&header.ID, &header.Code, &header.CustomerName, &header.CustomerPhone, &header.CreatedAt, &header.Sent); err != nil {
			return nil, fmt.Errorf("error scanning sales header: %w", err)
		}
		headers = append(headers, header)
	}
	return headers, nil
}
func (r *salesRepository) FindSalesByCustomerName(customerName string) ([]SalesHeader, error) {
	rows, err := r.db.Query(`
		SELECT id, code, customer_name, COALESCE(customer_phone,'') as customer_phone, created_at, sent
		FROM sales_headers
		WHERE customer_name ILIKE $1`, "%"+customerName+"%")
	if err != nil {
		return nil, fmt.Errorf("error querying sales by customer name: %w", err)
	}
	defer rows.Close()

	var headers []SalesHeader
	for rows.Next() {
		var header SalesHeader
		if err := rows.Scan(&header.ID, &header.Code, &header.CustomerName, &header.CustomerPhone, &header.CreatedAt, &header.Sent); err != nil {
			return nil, fmt.Errorf("error scanning sales header: %w", err)
		}
		headers = append(headers, header)
	}
	return headers, nil
}
func (r *salesRepository) FindAllSales() ([]SalesHeader, error) {
	rows, err := r.db.Query(`
		SELECT id, code, customer_name, COALESCE(customer_phone,'') as customer_phone,  created_at, sent
		FROM sales_headers`)
	if err != nil {
		return nil, fmt.Errorf("error querying all sales: %w", err)
	}
	defer rows.Close()

	var headers []SalesHeader
	for rows.Next() {
		var header SalesHeader
		if err := rows.Scan(&header.ID, &header.Code, &header.CustomerName, &header.CustomerPhone, &header.CreatedAt, &header.Sent); err != nil {
			return nil, fmt.Errorf("error scanning sales header: %w", err)
		}
		headers = append(headers, header)
	}
	return headers, nil
}

func (r *salesRepository) DeleteSalesByHeaderID(id string) error {
	_, err := r.db.Exec(`
		DELETE FROM sales_headers
		WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting sales header: %w", err)
	}
	return nil
}
func (r *salesRepository) CreateSalesDetail(detail SalesDetail) (SalesDetail, error) {
	_, err := r.db.Exec(`
		INSERT INTO sales_details (id, sales_header_id, item_id, quantity, price, amount)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		detail.ID, detail.SalesHeaderID, detail.ItemID,
		detail.Quantity, detail.Price, detail.Amount,
	)
	if err != nil {
		return SalesDetail{}, fmt.Errorf("error inserting sales detail: %w", err)
	}
	return detail, nil
}

func (r *salesRepository) UpdateSalesDetail(detail SalesDetail) (SalesDetail, error) {
	_, err := r.db.Exec(`
		UPDATE sales_details
		SET item_id = $1, quantity = $2, price = $3, amount = $4
		WHERE id = $5`,
		detail.ItemID, detail.Quantity, detail.Price, detail.Amount, detail.ID,
	)
	if err != nil {
		return SalesDetail{}, fmt.Errorf("error updating sales detail: %w", err)
	}
	return detail, nil
}

func (r *salesRepository) FindSalesDetailsByHeaderID(headerID string) ([]SalesDetail, error) {
	rows, err := r.db.Query(`
		SELECT sd.id, sd.sales_header_id, sd.item_id, i.code as item_code, i.description as item_description, sd.quantity, sd.price, sd.amount
		FROM sales_details sd
		INNER JOIN items i ON sd.item_id = i.id
		WHERE sales_header_id = $1`, headerID)
	if err != nil {
		return nil, fmt.Errorf("error querying sales details by header ID: %w", err)
	}
	defer rows.Close()

	var details []SalesDetail
	for rows.Next() {
		var detail SalesDetail
		if err := rows.Scan(&detail.ID, &detail.SalesHeaderID, &detail.ItemID,
			&detail.ItemCode, &detail.ItemDescription, &detail.Quantity,
			&detail.Price, &detail.Amount); err != nil {
			return nil, fmt.Errorf("error scanning sales detail: %w", err)
		}
		details = append(details, detail)
	}
	return details, nil
}

func (r *salesRepository) DeleteSalesDetailByID(id string) error {
	_, err := r.db.Exec(`
		DELETE FROM sales_details
		WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting sales detail: %w", err)
	}
	return nil
}

func (r *salesRepository) SendSalesHeader(id string) (SalesHeader, error) {
	_, err := r.db.Exec(`
		UPDATE sales_headers
		SET sent = true
		WHERE id = $1`, id)
	if err != nil {
		return SalesHeader{}, fmt.Errorf("error sending sales header: %w", err)
	}

	header, err := r.FindSalesByHeaderID(id)
	if err != nil {
		return SalesHeader{}, fmt.Errorf("error finding sales header after sending: %w", err)
	}
	return header, nil
}

func (r *salesRepository) GetNextNumber()(string, error){
	var nextCounter string
	err := r.db.QueryRow(`
	SELECT   
		REPEAT(
			'0',
			10 - LENGTH(CAST(COALESCE(MAX(CAST(code AS integer)), 0) + 1 AS varchar))
		) || CAST(COALESCE(MAX(CAST(code AS integer)), 0) + 1 AS varchar) AS next_counter
		FROM sales_headers
	`).Scan(&nextCounter)
	if err != nil {
		if err == sql.ErrNoRows {
			return "0000000001", nil
		}
		return "", err
	}
	return nextCounter, nil
}