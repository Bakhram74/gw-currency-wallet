package v1

import (
	"errors"
	"net/http"

	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
	"github.com/Bakhram74/gw-currency-wallet/internal/service/entity"
	"github.com/gin-gonic/gin"
)

func (r *Router) register(ctx *gin.Context) {
	var reqBody entity.RegisterReq

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := r.service.Auth.Register(ctx, reqBody.Name, reqBody.Password, reqBody.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserExists) || errors.Is(err, repository.ErrEmailFormat) {
			errorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, map[string]string{
		"message": "User registered successfully",
	})
}

func (r *Router) login(ctx *gin.Context) {
	var reqBody entity.LoginReq

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	token, err := r.service.Auth.Login(ctx, reqBody.Name, reqBody.Password)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			errorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
