package v1

import (
	"net/http"

	"github.com/Bakhram74/gw-currency-wallet/internal/service/entity"
	"github.com/Bakhram74/gw-currency-wallet/pkg/httpserver"
	"github.com/gin-gonic/gin"
)

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
