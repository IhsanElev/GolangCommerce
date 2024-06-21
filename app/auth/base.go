package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)
	_ = handler

	authRouter := router.Group("auth")
	{
		authRouter.Post("Register", handler.Register)
		authRouter.Post("Login", handler.Login)
	}
}
