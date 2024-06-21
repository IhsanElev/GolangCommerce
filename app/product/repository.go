package product

import (
	"GolangEcommerceDDD/infra/response"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateProduct(ctx context.Context, model Product) (err error) {
	query := `
	INSERT INTO product (
		sku,name, stock, price,created_at,updated_at
		) VALUES (
		:sku, :name, :stock, :price,:created_at, :updated_at
		)
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, model)
	return

}
func (r repository) GetAllProductWithPaginationCursor(ctx context.Context, model ProductPagination) (products []Product, err error) {
	query := `
	SELECT id, sku, name, stock, price, created_at, updated_at from product where id > $1 ORDER BY id asc LIMIT $2
	`
	r.db.SelectContext(ctx, &products, query, model.Cursor, model.Size)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return
	}
	return
}

func (r repository) GetProductBySKU(ctx context.Context, sku string) (product Product, err error) {
	query := `
	SELECT 
		id, sku, name
		,stock, price
		,created_at
		,updated_at
	FROM product
	WHERE sku=$1
	`

	err = r.db.GetContext(ctx, &product, query, sku)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}
	return
}
