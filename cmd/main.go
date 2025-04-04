package main

import (
	//"fmt"
	//"net/http"
    
	"goapp/database"

	"github.com/gofiber/fiber/v2"
)

/*
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Docker")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server started on port 8080")
    http.ListenAndServe(":8080", nil)
}
*/

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!!, This is Go-Fiber")
	})

    database.ConnectDB()

	app.Listen(":8080")
}