package main

import (
	"echo-restApi/routes"
	"echo-restApi/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Inisialisasi database
	utils.InitDB()

	// Routing untuk produk
	routes.ProductRoutes(e)

	// Jalankan server di port 8080
	e.Start(":8080")
}
