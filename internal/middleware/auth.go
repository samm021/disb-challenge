package middleware

import (
	"disbursement-service/domain"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("SECRET_KEY")

type UserClaim struct {
	jwt.RegisteredClaims
	UserId string
	Pin    int16
}

func Authenticate(userService domain.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var userClaim UserClaim

		tokenString := strings.ReplaceAll(ctx.Get("Authorization"), "Bearer ", "")
		if tokenString == "" {
			return ctx.SendStatus(401)
		}

		token, err := jwt.ParseWithClaims(tokenString, &userClaim, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			return ctx.SendStatus(401)
		}

		user, err := userService.FindByUserId(ctx.Context(), userClaim.UserId)
		if err != nil {
			return ctx.SendStatus(401)
		}

		if user.Pin != userClaim.Pin {
			return ctx.SendStatus(401)
		}

		// sets user for next operation in the request
		ctx.Locals("x-user", user)
		return ctx.Next()
	}
}
