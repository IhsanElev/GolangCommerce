package utility

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		publicid := uuid.NewString()
		tokenString, err := GenerateToken(publicid, "user", "secret")
		require.Nil(t, err)
		require.NotEmpty(t, tokenString)
		log.Println(tokenString)
	})
}

func TestValidateToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		publicid := uuid.NewString()
		role := "user"
		tokenString, err := GenerateToken(publicid, role, "secret")
		require.Nil(t, err)
		require.NotEmpty(t, tokenString)

		jwtId, jwtRole, err := ValidateToken(tokenString, "secret")
		require.Nil(t, err)
		require.NotEmpty(t, jwtId)
		require.NotEmpty(t, jwtRole)
		require.Equal(t, publicid, jwtId)
		require.Equal(t, role, jwtRole)
		log.Println(tokenString)
	})
}
