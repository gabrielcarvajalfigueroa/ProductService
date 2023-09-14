package main

import (

	"fmt"

	db "github.com/gabrielcarvajalfigueroa/ProductService/config"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)



func main(){

	fmt.Println("Server starting...")
	godotenv.Load()
	fmt.Println("Loaded env variables...")
	
	//Connect to database

	db.Connect()

	app := fiber.New()

	app.Listen(":3000")


}