package product

import (
	"GolangEcommerceDDD/external/database"
	"GolangEcommerceDDD/infra/response"
	"GolangEcommerceDDD/internal/config"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

// init initializes the product service
var svc service

func init() {
	// load configuration from file
	filename := "../../cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}
	// connect to postgres database
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	// create repository
	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateProduct_Success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "Product 1",
		Stock: 10,
		Price: 10000,
	}
	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}

func TestCreateProduct_Failed(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "",
		Stock: 10,
		Price: 100,
	}
	err := svc.CreateProduct(context.Background(), req)
	require.NotNil(t, err)
	require.Equal(t, response.ErrProductRequired, err)

}

func TestListProduct_success(t *testing.T) {
	pagination := ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	}
	products, err := svc.ListProducts(context.Background(), pagination)
	require.Nil(t, err)
	require.NotNil(t, products)
	log.Printf("%+v", products)
}

func TestDetailProduct_success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "baju baru ihsan",
		Stock: 1,
		Price: 1_0000,
	}
	ctx := context.Background()

	err := svc.CreateProduct(ctx, req)
	require.Nil(t, err)
	products, err := svc.ListProducts(ctx, ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	})
	require.Nil(t, err)
	require.NotNil(t, products)
	require.Greater(t, len(products), 0)
	log.Printf("%+v", products)
}
