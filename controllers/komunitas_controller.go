// controllers/komunitas_controller.go
package controllers

import (
	"donasiPohon/config"
	"donasiPohon/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Fungsi registrasi komunitas
func RegisterKomunitas(c echo.Context) error {
    komunitas := new(models.Komunitas)
    if err := c.Bind(komunitas); err != nil {
        return err
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(komunitas.Password), 8)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Gagal hashing password")
    }
    komunitas.Password = string(hashedPassword)

    if err := config.DB.Create(&komunitas).Error; err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    return c.JSON(http.StatusOK, map[string]string{
        "message": "Registrasi komunitas berhasil",
    })
}

// Fungsi login komunitas
func LoginKomunitas(c echo.Context) error {
    komunitas := new(models.Komunitas)
    if err := c.Bind(komunitas); err != nil {
        return err
    }

    var dbKomunitas models.Komunitas
    config.DB.Where("email = ?", komunitas.Email).First(&dbKomunitas)

    if dbKomunitas.ID == 0 {
        return c.JSON(http.StatusUnauthorized, "Komunitas tidak ditemukan")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(dbKomunitas.Password), []byte(komunitas.Password)); err != nil {
        return c.JSON(http.StatusUnauthorized, "Password salah")
    }

    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["komunitas_id"] = dbKomunitas.ID
    claims["user_type"] = "komunitas"
    claims["exp"] = time.Now().AddDate(0, 1, 0).Unix() // Token berlaku 1 bulan dari sekarang

    t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, map[string]string{
        "token": t,
    })
}
