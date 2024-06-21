package transaction

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSetSubTotal(t *testing.T) {
	var trx = Transaction{
		ProductPrice: 10_000,
		Amount:       10,
	}

	expected := uint(100_000)
	trx.SetSubTotal()
	log.Printf("%+v\n", trx)
	require.Equal(t, expected, trx.SubTotal)
}

func TestSetGrandTotal(t *testing.T) {
	t.Run("without set subtotal", func(t *testing.T) {
		var trx = Transaction{
			ProductPrice: 10_000,
			Amount:       10,
		}

		expected := uint(100_000)
		trx.SetGrandTotal()

		require.Equal(t, expected, trx.GrandTotal)
	})
	t.Run("without platform fee", func(t *testing.T) {
		var trx = Transaction{
			ProductPrice: 10_000,
			Amount:       10,
		}

		expected := uint(100_000)
		trx.SetSubTotal()
		trx.SetGrandTotal()

		require.Equal(t, expected, trx.GrandTotal)
	})
}

func TestSetProductJson(t *testing.T) {
	var product = Product{
		Id:    1,
		SKU:   uuid.NewString(),
		Name:  "Product 1",
		Price: 10_000,
	}
	var trx = Transaction{}
	err := trx.SetProductJSON(product)

	require.Nil(t, err)
	require.NotNil(t, trx.ProductJSON)
	productFromTrx, err := trx.GetProduct()
	require.Nil(t, err)
	require.NotEmpty(t, productFromTrx)
	require.Equal(t, productFromTrx, product)

}

func TestGetStatus(t *testing.T) {
	type tabletest struct {
		title    string
		expected string
		trx      Transaction
	}
	var table = []tabletest{{
		title:    "created",
		expected: TRX_CREATED,
		trx:      Transaction{Status: TransactionStatus_Created},
	},
		{
			title:    "on progress",
			expected: TRX_ON_PROGRESS,
			trx:      Transaction{Status: TransactionStatus_Progress},
		},
		{
			title:    "in delivery",
			expected: TRX_IN_DELIVERY,
			trx:      Transaction{Status: TransactionStatus_InDelivery},
		},
		{
			title:    "completed",
			expected: TRX_COMPLETED,
			trx:      Transaction{Status: TransactionStatus_Completed},
		},
		{
			title:    "unknown status",
			expected: TRX_UNKNOWN,
			trx:      Transaction{Status: 0},
		},
	}
	for _, tc := range table {
		t.Run(tc.title, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.trx.GetStatus())
		})
	}

}
