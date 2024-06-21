package product

import (
	"GolangEcommerceDDD/app/auth"
	infrafiber "GolangEcommerceDDD/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)
	productRouter := router.Group("products")
	{
		productRouter.Get("", handler.GetListProduct)

		productRouter.Post("",
			infrafiber.CheckAuth(),
			infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}),
			handler.CreateProduct)

		productRouter.Get("/sku/:sku", handler.GetProductDetail)
	}

}
