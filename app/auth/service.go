package auth

import (
	"GolangEcommerceDDD/infra/response"
	"GolangEcommerceDDD/internal/config"
	"context"
	"fmt"
)

type Repository interface {
	CreateAuth(ctx context.Context, model AuthEntity) (err error)
	GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error)
}
type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {
	authEntity := NewFromRegisterRequest(req)
	if err = authEntity.Validate(); err != nil {
		return
	}
	if err = authEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
		return
	}
	model, err := s.repo.GetAuthByEmail(ctx, authEntity.Email)
	fmt.Println(model)
	if err != nil {
		if err != response.ErrNotFound {
			return
		}
	}
	if model.isExist() {
		return response.ErrEmailAlreadyUsed
	}
	return s.repo.CreateAuth(ctx, authEntity)
}

func (s service) login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
	authEntity := NewFromLoginRequest(req)
	if err = authEntity.Validate(); err != nil {
		return
	}

	model, err := s.repo.GetAuthByEmail(ctx, authEntity.Email)
	if err != nil {
		if err != response.ErrNotFound {
			return
		}
	}
	if err = authEntity.VerifyFromPlainPassword(model.Password); err != nil {
		err = response.ErrPasswordNotMatch
		return
	}
	token, err = model.GenerateToken(config.Cfg.App.Encryption.JWTSecret)
	return
}
