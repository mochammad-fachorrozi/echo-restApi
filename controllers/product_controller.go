package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"echo-restApi/models"
	"echo-restApi/utils"

	"github.com/labstack/echo/v4"
)

// GetProducts mengembalikan semua produk
func GetProducts(c echo.Context) error {
	db := utils.DB

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mendapatkan data produk"})
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal membaca data produk"})
		}
		products = append(products, product)
	}

	return c.JSON(http.StatusOK, products)
}

// GetProductByID mengembalikan produk berdasarkan ID
func GetProductByID(c echo.Context) error {
	db := utils.DB

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID produk tidak valid"})
	}

	row := db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	var product models.Product
	err = row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Produk tidak ditemukan"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mendapatkan data produk"})
	}

	return c.JSON(http.StatusOK, product)
}

// CreateProduct menambahkan produk baru
func CreateProduct(c echo.Context) error {
	db := utils.DB

	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Data tidak valid"})
	}

	// Lakukan validasi data di sini jika diperlukan

	// Query SQL untuk menambahkan data ke tabel
	_, err := db.Exec("INSERT INTO products (name, price, quantity) VALUES (?, ?, ?)", product.Name, product.Price, product.Quantity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menambahkan produk"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Produk berhasil ditambahkan"})
}

// UpdateProduct mengupdate produk berdasarkan ID
func UpdateProduct(c echo.Context) error {
	db := utils.DB

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID produk tidak valid"})
	}

	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Data tidak valid"})
	}

	// Lakukan validasi data di sini jika diperlukan

	// Query SQL untuk mengupdate data produk berdasarkan ID
	result, err := db.Exec("UPDATE products SET name=?, price=?, quantity=? WHERE id=?", product.Name, product.Price, product.Quantity, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengupdate produk"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Produk tidak ditemukan"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Produk berhasil diupdate"})
}

// DeleteProduct menghapus produk berdasarkan ID
func DeleteProduct(c echo.Context) error {
	db := utils.DB

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID produk tidak valid"})
	}

	// Query SQL untuk menghapus data produk berdasarkan ID
	result, err := db.Exec("DELETE FROM products WHERE id=?", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus produk"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Produk tidak ditemukan"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Produk berhasil dihapus"})
}
