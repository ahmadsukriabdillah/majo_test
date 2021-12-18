package midleware

import (
	"errors"
	"fmt"
	"majo_test/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenMaker utils.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(utils.AuthorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			utils.NewResponseError(ctx, errors.New("authorization header is not provided"))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			utils.NewResponseError(ctx, errors.New("invalid authorization header format"))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != utils.AuthorizationTypeBearer {
			utils.NewResponseError(ctx, errors.New(fmt.Sprintf("unsupported authorization type %s", authorizationType)))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			utils.NewResponseError(ctx, errors.New("invalid token"))
			return
		}

		ctx.Set(utils.AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
