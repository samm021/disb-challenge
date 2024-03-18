package middleware

import (
	"disbursement-service/domain"
	"disbursement-service/dto"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("SECRET_KEY")

type UserClaim struct {
	jwt.RegisteredClaims
	UserId string `json:"user_id"`
	Pin    int16  `json:"pin"`
}

func Authenticate(userService domain.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var userClaim UserClaim

		tokenString := strings.ReplaceAll(ctx.Get("Authorization"), "Bearer ", "")
		if tokenString == "" {
			return ctx.SendStatus(401)
		}

		jwt.ParseWithClaims(tokenString, &userClaim, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		// TODO: Proper validation of token
		// if err != nil || !token.Valid {
		// 	return ctx.SendStatus(401)
		// }

		user := userService.FindByUserId(ctx.Context(), userClaim.UserId)
		if user.UserId != userClaim.UserId {
			return ctx.SendStatus(401)
		}

		if user.Pin != userClaim.Pin {
			return ctx.SendStatus(401)
		}

		// sets user for next operation in the request
		userData := dto.UserData{UserId: user.UserId}
		ctx.Locals("x-user", userData)
		return ctx.Next()
	}
}
