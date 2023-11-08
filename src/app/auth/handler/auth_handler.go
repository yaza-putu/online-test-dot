package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yaza-putu/online-test-dot/src/app/auth/repository"
	"github.com/yaza-putu/online-test-dot/src/app/auth/service"
	"github.com/yaza-putu/online-test-dot/src/app/auth/validation"
	"github.com/yaza-putu/online-test-dot/src/http/request"
	"github.com/yaza-putu/online-test-dot/src/http/response"
	"github.com/yaza-putu/online-test-dot/src/logger"
	"net/http"
)

type authHandler struct {
	authService service.AuthInterface
}

func NewAuthHandler() *authHandler {
	return &authHandler{
		authService: service.NewAuthService(repository.NewUserRepository(), service.NewToken()),
	}
}

func (a *authHandler) Create(ctx echo.Context) error {
	// request validation & capture data
	req := validation.TokenValidation{}
	b := ctx.Bind(&req)
	if b != nil {
		return ctx.JSON(http.StatusBadRequest, response.Api(
			response.SetCode(400), response.SetMessage(fmt.Sprint("Bad request : %s", b.Error())),
		))
	}

	// validation form
	res, err := request.Validation(&req)
	logger.New(err, logger.SetType(logger.INFO))

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, res)
	}

	r := a.authService.Login(ctx.Request().Context(), req.Email, req.Password)

	return ctx.JSON(r.Code, r)
}
