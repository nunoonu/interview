package handlers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log/slog"
)

const BEARER_SCHEMA = "Bearer "

func VerifyJWT() gin.HandlerFunc {

	secretKey := "secret"
	return func(ctx *gin.Context) {
		jwtToken := ctx.GetHeader("Authorization")
		if jwtToken == "" {
			slog.Error("No token found")
			ctx.AbortWithStatus(401)
		}

		tokenStr := jwtToken[len(BEARER_SCHEMA):]
		t, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
				slog.Error("Invalid token")
				return nil, errors.New("Invalid token")
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			slog.Error("Verify token fail", slog.String("Err", err.Error()))
			ctx.AbortWithStatus(401)
		}
		claims, ok := t.Claims.(jwt.MapClaims)
		if !ok {
			slog.Error("Error parsing claims")
			ctx.AbortWithStatus(401)
		}

		rct := mapRouteContext(claims)
		if rct.Role != "admin" {
			slog.Error("Role is not admin")
			ctx.AbortWithStatus(401)
		}

		ctx.Set("rct", rct)

		ctx.Next()
	}

}

func mapRouteContext(claims jwt.MapClaims) *RouteContext {
	return &RouteContext{
		Role:   claims["role"].(string),
		UserID: claims["userId"].(string),
	}
}

type RouteContext struct {
	Role   string `json:"role"`
	UserID string `json:"userId"`
}
