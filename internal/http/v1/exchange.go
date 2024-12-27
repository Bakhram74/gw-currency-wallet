package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
	"github.com/Bakhram74/gw-currency-wallet/internal/service/entity"
	"github.com/Bakhram74/gw-currency-wallet/pkg/client/redis"
	"github.com/Bakhram74/gw-currency-wallet/pkg/httpserver"
	"github.com/Bakhram74/proto-exchange/pb"
	"github.com/gin-gonic/gin"
)

// @Summary Get rates
// @Security ApiKeyAuth
// @Description Getting exchange rates from rub_rates table
// @Tags exchange
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]float32
// @Failure      400,404,500  {func}  httpserver.ErrorResponse
// @Router /exchange/rates [get]
func (r *Router) rates(ctx *gin.Context) {

	resp, err := r.grpcClient.GetExchangeRates(ctx, &pb.Empty{})
	if err != nil {
		httpserver.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp.Rates)
}

// @Summary Exchange currency
// @Security ApiKeyAuth
// @Description Exchange currency
// @Tags exchange
// @Accept  json
// @Produce  json
// @Param input body entity.ExchangeReq true "FromCurrency, ToCurrency, Amount"
// @Success 200 {object}  entity.ExchangeResponse
// @Failure      400,404,500  {func}  httpserver.ErrorResponse
// @Router /exchange/rates [get]
func (r *Router) exchange(ctx *gin.Context) {
	var reqBody entity.ExchangeReq
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		httpserver.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := getUserId(ctx)
	if err != nil {
		httpserver.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	from := strings.ToUpper(reqBody.From)
	to := strings.ToUpper(reqBody.To)

	var rate float32

	req := &pb.CurrencyRequest{
		FromCurrency: from,
		ToCurrency:   to,
	}

	rate, err = redis.GetRate(ctx, userID, from, to)
	if err != nil {
		rateResp, err := r.grpcClient.GetExchangeRateForCurrency(ctx, req)
		if err != nil {
			httpserver.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		rate = rateResp.Rate
	}

	exchanged, err := r.service.Exchange.ExchangeCurrency(ctx, userID, from, to, rate, reqBody.Amount)
	if err != nil {
		if errors.Is(err, repository.ErrInsufficientBalance) {
			httpserver.ErrorResponse(ctx, http.StatusBadRequest, ErrInsufficientAmount.Error())
			return
		}
		httpserver.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	resp := entity.ExchangeResponse{
		Message:    "Exchange successful",
		Amount:     exchanged.ExchangedAmount,
		NewBalance: map[string]float32{from: exchanged.FromBalance, to: exchanged.ToBalance},
	}
	ctx.JSON(http.StatusOK, resp)
}
