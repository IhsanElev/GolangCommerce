package product

import (
	infrafiber "GolangEcommerceDDD/infra/fiber"
	"GolangEcommerceDDD/infra/response"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) CreateProduct(ctx *fiber.Ctx) error {
	req := CreateProductRequestPayload{}
	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(infrafiber.WithMessage("Invalid Payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}
	if err := h.svc.CreateProduct(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(infrafiber.WithMessage("Invalid Payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}
	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("Create Product Success"),
	).Send(ctx)
}
func (h handler) GetListProduct(ctx *fiber.Ctx) error {
	var req ListProductRequestPayload

	if err := ctx.QueryParser(&req); err != nil {
		// Handle parsing error
		return infrafiber.NewResponse(
			infrafiber.WithMessage("Invalid Payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}
	products, err := h.svc.ListProducts(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(infrafiber.WithMessage("Invalid Payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}
	productListResponse := NewProductListResponseFromEntity(products)
	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("Get List Product Success"),
		infrafiber.WithPayload(productListResponse),
		infrafiber.WithQuery(req.GenerateDefaultValue()),
	).Send(ctx)
}

func (h handler) GetProductDetail(ctx *fiber.Ctx) error {
	sku := ctx.Params("sku", "")
	log.Println(sku)
	if sku == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	product, err := h.svc.ProductDetail(ctx.UserContext(), sku)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	productDetail := ProductDetailResponse{
		Id:        product.ID,
		Name:      product.Name,
		SKU:       product.SKU,
		Stock:     product.Stock,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get product detail success"),
		infrafiber.WithPayload(productDetail),
	).Send(ctx)

}
