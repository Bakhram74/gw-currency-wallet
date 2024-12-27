package v1

import (
	"errors"
	"net/http"

	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
	"github.com/Bakhram74/gw-currency-wallet/internal/service/entity"
	"github.com/Bakhram74/gw-currency-wallet/pkg/httpserver"
	"github.com/gin-gonic/gin"
)

// @Summary register
// @Description Create user with his wallet.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body entity.RegisterReq true "name password email"
// @Success 200 {string} string "User registered successfully"
// @Failure      400,404,500  {func}  httpserver.ErrorResponse
// @Router /register [post]
func (r *Router) register(ctx *gin.Context) {
	var reqBody entity.RegisterReq

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		httpserver.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := r.service.Auth.Register(ctx, reqBody.Name, reqBody.Password, reqBody.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserExists) || errors.Is(err, repository.ErrEmailFormat) {
			httpserver.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		httpserver.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, map[string]string{
		"message": "User registered successfully",
	})
}



// @Summary login
// @Description Login user.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body entity.LoginReq true "name password"
// @Success 200 {string} string "token"
// @Failure      400,404,500  {func}  httpserver.ErrorResponse
// @Router /login [post]
func (r *Router) login(ctx *gin.Context) {
	var reqBody entity.LoginReq

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		httpserver.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	token, err := r.service.Auth.Login(ctx, reqBody.Name, reqBody.Password)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			httpserver.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}
		httpserver.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
