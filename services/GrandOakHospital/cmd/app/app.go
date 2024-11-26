package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/alivecheck", func(c *fiber.Ctx) error {
		fmt.Println("Health Check request received!")
		return c.SendString("ok")
	})

	app.Get("/doctors/available", func(c *fiber.Ctx) error {
		doctors := []Doctor{
			{
				ID:        1,
				FullName:  "Dr. Michael Green",
				Gender:    "Male",
				Phone:     "555-4321",
				Specialty: "Rheumatology",
				Schedule: []Schedule{
					{Day: "Monday", AvailableRange: "08:00 AM - 04:00 PM"},
					{Day: "Wednesday", AvailableRange: "10:00 AM - 02:00 PM"},
					{Day: "Friday", AvailableRange: "01:00 PM - 05:00 PM"},
				},
			},
			{
				ID:        2,
				FullName:  "Dr. Emily Carter",
				Gender:    "Female",
				Phone:     "555-6789",
				Specialty: "Nephrology",
				Schedule: []Schedule{
					{Day: "Tuesday", AvailableRange: "09:00 AM - 03:00 PM"},
					{Day: "Thursday", AvailableRange: "08:00 AM - 12:00 PM"},
				},
			},
			{
				ID:        3,
				FullName:  "Dr. Sarah Lee",
				Gender:    "Female",
				Phone:     "555-9870",
				Specialty: "Geriatrics",
				Schedule: []Schedule{
					{Day: "Monday", AvailableRange: "10:00 AM - 01:00 PM"},
					{Day: "Wednesday", AvailableRange: "02:00 PM - 06:00 PM"},
					{Day: "Friday", AvailableRange: "09:00 AM - 01:00 PM"},
				},
			},
			{
				ID:        4,
				FullName:  "Dr. James Brown",
				Gender:    "Male",
				Phone:     "555-1123",
				Specialty: "Podiatry",
				Schedule: []Schedule{
					{Day: "Monday", AvailableRange: "08:00 AM - 02:00 PM"},
					{Day: "Thursday", AvailableRange: "12:00 PM - 04:00 PM"},
					{Day: "Saturday", AvailableRange: "10:00 AM - 02:00 PM"},
				},
			},
		}

		fmt.Println("Available doctors requested!")
		return c.JSON(doctors)
	})

	app.Listen(":6969")
}
