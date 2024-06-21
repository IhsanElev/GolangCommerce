package auth

import (
	"GolangEcommerceDDD/external/database"
	"context"
	"fmt"
	"log"
	"testing"

	"GolangEcommerceDDD/internal/config"

	"github.com/google/uuid"
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

func TestRegister_Success(t *testing.T) {
	req := RegisterRequestPayload{
		Email:    fmt.Sprintf("%v@gmail.com", uuid.NewString()),
		Password: "mysecretpassword",
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)
}

func TestLogin_Success(t *testing.T) {
	email := fmt.Sprintf("%v@gmail.com", uuid.NewString())
	password := "mysecretpassword"
	req := RegisterRequestPayload{
		Email:    email,
		Password: password,
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)

	token, err := svc.login(context.Background(), LoginRequestPayload{
		Email:    email,
		Password: password,
	})
	require.Nil(t, err)
	require.NotEmpty(t, token)
	log.Println(token)
}
