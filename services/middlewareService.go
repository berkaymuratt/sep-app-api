package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//go:generate mockgen -destination=../mocks/service/mockMiddlewareService.go -package=services github.com/berkaymuratt/sep-app-api/services MiddlewareServiceI
type MiddlewareServiceI interface {
	Middleware(c *fiber.Ctx) error
	CORSMiddleware() func(*fiber.Ctx) error
}

type MiddlewareService struct {
	MiddlewareServiceI
	jwtService JwtService
}

func NewMiddlewareService(
	jwtService JwtService,
) MiddlewareService {
	return MiddlewareService{
		jwtService: jwtService,
	}
}

func (middlewareService MiddlewareService) Middleware(c *fiber.Ctx) error {
	jwtToken := c.Get("jwt")
	_, err := middlewareService.jwtService.CheckJwt(jwtToken)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthanticated",
		})
	}

	return c.Next()
}

func (middlewareService MiddlewareService) CORSMiddleware() func(*fiber.Ctx) error {

	corsConfig := cors.Config{
		AllowOrigins:     "http://localhost:7357",
		AllowMethods:     "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE",
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, jwt",
		AllowCredentials: true,
		MaxAge:           3600,
	}

	corsMiddleware := cors.New(corsConfig)
	return corsMiddleware
}
