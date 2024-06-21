package product

import (
	"GolangEcommerceDDD/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Product := Product{
			Name:  "Baju Baru",
			Stock: 1,
			Price: 10000,
		}
		err := Product.Validate()
		require.Nil(t, err)

	})
	t.Run("Product Invalid", func(t *testing.T) {
		Product := Product{
			Name:  "Baj",
			Stock: 1,
			Price: 10000,
		}
		err := Product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductInvalid, err)

	})
	t.Run("Product Required", func(t *testing.T) {
		Product := Product{
			Name:  "",
			Stock: 1,
			Price: 10000,
		}
		err := Product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})
}

func TestStock(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Product := Product{
			Name:  "Baju Baru",
			Stock: 1,
			Price: 10000,
		}
		err := Product.Validate()
		require.Nil(t, err)

	})
	t.Run("Stock Invalid", func(t *testing.T) {
		Product := Product{
			Name:  "Baju baru",
			Stock: 0,
			Price: 10000,
		}
		err := Product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrStockInvalid, err)

	})
	t.Run("Price Invalid", func(t *testing.T) {
		Product := Product{
			Name:  "Baju Baru",
			Stock: 1,
			Price: 0,
		}
		err := Product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})
}
