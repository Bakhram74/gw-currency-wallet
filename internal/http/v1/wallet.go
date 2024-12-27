package v1

import (
	"errors"
	"net/http"

	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
	"github.com/Bakhram74/gw-currency-wallet/internal/service/entity"
	"github.com/Bakhram74/gw-currency-wallet/pkg/httpserver"
	"github.com/gin-gonic/gin"
)

// @Summary Get balance
// @Security ApiKeyAuth
// @Description Handler for Getting balance from wallet
// @Tags wallet
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Balance
// @Failure      400,404,500  {func}  httpserver.ErrorResponse
// @Router /balance [get]
func (r *Router) balance(ctx *gin.Context) {
	userID, err := getUserId(ctx)

	if err != nil {
		httpserver.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	w, err := r.service.Balance.GetBalance(ctx, userID)
	if err != nil {
		httpserver.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	resp := entity.Balance{
		USD: w.Usd,
		RUB: w.Rub,
		EUR: w.Eur,
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"balance": resp,
	})
}

// @Summary Deposit
// @Security ApiKeyAuth
// @Description Deposit to users wallet
// @Tags wallet
// @Accept  json
// @Produce  json
// @Param input body entity.Transaction true "amount, currency"
// @Success 200 {object} entity.DepositResponse
// @Failure      400,404,500  {func}  httpserver.ErrorResponse
// @Router /wallet/deposit [post]
func (r *Router) deposit(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		httpserver.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	var reqBody entity.Transaction

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		httpserver.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if reqBody.Amount <= 0 {
		httpserver.ErrorResponse(ctx, http.StatusBadRequest, ErrInvalidCurrencyAmount.Error())
		return
	}

	isValid := entity.IsValidCurrency(string(reqBody.Currency))
	if !isValid {
		httpserver.ErrorResponse(ctx, http.StatusBadRequest, ErrInvalidCurrencyAmount.Error())
		return
	}

	wallet, err := r.service.Balance.DepositBalance(ctx, userID, reqBody)
	if err != nil {
		httpserver.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	balance := entity.Balance{
		USD: wallet.Usd,
		RUB: wallet.Rub,
		EUR: wallet.Eur,
	}

	resp := entity.DepositResponse{
		Message:    "Account topped up successfully",
		NewBalance: balance,
	}

	ctx.JSON(http.StatusOK, resp)
}


// @Summary Withdraw
// @Security ApiKeyAuth
// @Description Withdraw from users wallet
// @Tags wallet
// @Accept  json
// @Produce  json
// @Param input body entity.Transaction true "amount, currency"
// @Success 200 {object} entity.DepositResponse
// @Failure      400,404,500  {func}  httpserver.ErrorResponse
// @Router /wallet/withdraw [post]
func (r *Router) withdraw(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		httpserver.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	var reqBody entity.Transaction

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		httpserver.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if reqBody.Amount <= 0 {
		httpserver.ErrorResponse(ctx, http.StatusBadRequest, ErrInvalidCurrencyAmount.Error())
		return
	}

	isValid := entity.IsValidCurrency(string(reqBody.Currency))
	if !isValid {
		httpserver.ErrorResponse(ctx, http.StatusBadRequest, ErrInvalidCurrencyAmount.Error())
		return
	}

	wallet, err := r.service.Balance.WithdrawBalance(ctx, userID, reqBody)
	if err != nil {
		if errors.Is(err, repository.ErrInsufficientBalance) {
			httpserver.ErrorResponse(ctx, http.StatusBadRequest, ErrInsufficientAmount.Error())
			return
		}
		httpserver.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	balance := entity.Balance{
		USD: wallet.Usd,
		RUB: wallet.Rub,
		EUR: wallet.Eur,
	}

	resp := entity.DepositResponse{
		Message:    "Withdrawal successful",
		NewBalance: balance,
	}

	ctx.JSON(http.StatusOK, resp)
}
