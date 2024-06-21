package transaction

import (
	"GolangEcommerceDDD/external/database"
	"GolangEcommerceDDD/internal/config"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateTransaction(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := CreateTransactionRequestPayload{
			ProductSKU:   "3de4ce55-7d7a-4f8c-8be8-5e346d94ac25",
			Amount:       2,
			UserPublicId: "f4481548-7b23-4714-9760-798568a5506e",
		}
		err := svc.CreateTransaction(context.Background(), req)
		require.Nil(t, err)
	})
}
