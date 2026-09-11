package main

import (
	"khrai-chui-khrai/internal/calculator"
	"khrai-chui-khrai/internal/domain"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://khrai-chui-khrai-fe.vercel.app/"},
	}))

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("api/v1/calculate", calculate)

	log.Fatal(app.Listen(":3000"))
}

func calculate(c fiber.Ctx) error {

	var request domain.CalculateRequest

	err := c.Bind().Body(&request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	// valadate
	if request.TotalAmount != request.GovernmentSupport+request.PaidAmount {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid amount.",
		})
	}

	var purchasedTotal int64
	for _, o := range request.OrderItems {
		purchasedTotal += o.PurchasedAmount
	}

	if purchasedTotal != request.TotalAmount {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Total amount mismatch.",
		})
	}

	result := calculator.CalculateService(request)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": result,
	})
}
