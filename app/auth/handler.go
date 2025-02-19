package auth

import (
	infrafiber "GolangEcommerceDDD/infra/fiber"
	"GolangEcommerceDDD/infra/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func newHandler(svc service) *handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) Register(ctx *fiber.Ctx) error {
	var req = RegisterRequestPayload{}
	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
			infrafiber.WithHttpCode(http.StatusBadRequest),
			infrafiber.WithMessage("Register Fail"),
		).Send(ctx)
	}
	if err := h.svc.register(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[(err.Error())]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
			infrafiber.WithHttpCode(http.StatusBadRequest),
			infrafiber.WithMessage("Register Success"),
		).Send(ctx)
	}
	return infrafiber.NewResponse(infrafiber.WithHttpCode(http.StatusCreated)).Send(ctx)

}

func (h *handler) Login(ctx *fiber.Ctx) error {
	var req = LoginRequestPayload{}
	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
			infrafiber.WithHttpCode(http.StatusBadRequest),
			infrafiber.WithMessage("Login Fail"),
		).Send(ctx)
	}
	token, err := h.svc.login(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[(err.Error())]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}
	return infrafiber.NewResponse(infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithPayload(map[string]interface{}{
			"access_token": token,
		}),
		infrafiber.WithMessage("Login Success"),
	).Send(ctx)

}
