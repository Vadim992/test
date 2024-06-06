package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {

	//router := fiber.New()
	//
	//router.Get("/name/*", helloWorld)
	//router.Static("/", ".")
	//
	//router.Listen(":3000")

	//a := time.Duration(600000)

	fmt.Println(time.Duration(time.Minute * 6000000).Seconds())
}

func helloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString(ctx.Params("*"))
}
