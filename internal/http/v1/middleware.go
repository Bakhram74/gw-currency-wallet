package v1

import (
	"errors"
	"fmt"
	"net/http"

	"strings"

	"github.com/Bakhram74/gw-currency-wallet/pkg/httpserver"
	"github.com/Bakhram74/gw-currency-wallet/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker *jwt.JWTMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			httpserver.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())

			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			httpserver.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())

			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			httpserver.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())

			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			httpserver.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
func getUserId(ctx *gin.Context) (string, error) {
	payload, ok := ctx.Get(authorizationPayloadKey)
	if !ok {
		return "", errors.New("user id not found")
	}
	tokenPayload, ok := payload.(*jwt.Payload)
	if !ok {
		return "", errors.New("user id is of invalid type")
	}
	return tokenPayload.ID, nil
}
