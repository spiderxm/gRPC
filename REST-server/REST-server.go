package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"grpc-sample/helpers"
	"log"
)

type EmailInfo struct {
	Email   string `json:"email"`
	Subject string `json:"subject" jsonschema:"required"`
	Body    string `json:"body"`
}

func main() {
	e := godotenv.Load(".env")
	if e != nil {
		log.Fatalln("Could not start REST Server not able to load dependencies")
	}
	app := fiber.New()

	app.Post("/api/sendEmail", func(c *fiber.Ctx) error {
		emailData := new(EmailInfo)
		if err := c.BodyParser(emailData); err != nil {
			c.Status(422)
			return c.JSON(map[string]string{
				"Message": "Please provide valid input",
			})
		}
		resp, status := helpers.SendEmailDefault(emailData.Email, emailData.Body, emailData.Subject)
		c.Status(status)
		return c.JSON(resp)
	})

	log.Fatal(app.Listen(":3000"))
}
