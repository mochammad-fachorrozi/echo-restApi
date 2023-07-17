package routes

import (
	"echo-restApi/controllers"

	"github.com/labstack/echo/v4"
)

// ProductRoutes mengatur routing untuk tema product
func ProductRoutes(e *echo.Echo) {
	// Route untuk mendapatkan semua product
	e.GET("/produk", controllers.GetProducts)

	// Route untuk mendapatkan product berdasarkan ID
	e.GET("/produk/:id", controllers.GetProductByID)

	// Route untuk menambahkan product baru
	e.POST("/produk", controllers.CreateProduct)

	// Route untuk mengupdate product berdasarkan ID
	e.PUT("/produk/:id", controllers.UpdateProduct)

	// Route untuk menghapus product berdasarkan ID
	e.DELETE("/produk/:id", controllers.DeleteProduct)
}
