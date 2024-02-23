package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CheckerHealth(checkerHealthRoutes fiber.Router) {
	checkerHealthRoutes.Get("/", func(c *fiber.Ctx) error {

		/*
			- Variabel biasa BISA DIAMBIL nilai pointernya, caranya dengan MENAMBAHKAN tanda
			  ampersand (&) tepat sebelum nama variabel. Metode ini disebut dengan referencing.
			- Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara
				MENAMBAHKAN tanda asterisk (*) tepat sebelum nama variabel.
				Metode ini disebut dengan dereferencing.

		*/

		// start pointer
		var numberA int = 4
		var thisNumber *int = &numberA // ambil nilai pointer numberA dan masukkan ke variable thisNumber

		fmt.Println("numberA (value)   :", numberA)  // 4
		fmt.Println("numberA (address) :", &numberA) // 0xc20800a220 ambil nilai pointer

		fmt.Println("thisNumber (value)   :", *thisNumber) // 4 ambil nilai asli
		fmt.Println("thisNumber (address) :", thisNumber)  // 0xc20800a220

		numberA = 7

		fmt.Println("change numberA: ", numberA)
		fmt.Println("impact to thisNumber: ", *thisNumber)

		// welcome content
		resp := map[string]interface{}{
			"status":  "OK",
			"message": "Welcome to the Todo-List-API",
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	})
}
