package auth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuthEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		a := AuthEntity{
			Email:    "a@a.com",
			Password: "123456",
		}
		err := a.Validate()
		require.Nil(t, err)
	})
}
