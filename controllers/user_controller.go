// controllers/user_controller.go
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

func RegisterUser(c echo.Context) error {
    user := new(models.User)
    if err := c.Bind(user); err != nil {
        return err
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
    user.Password = string(hashedPassword)

    if err := config.DB.Create(&user).Error; err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    return c.JSON(http.StatusOK, user)
}

func LoginUser(c echo.Context) error {
    user := new(models.User)
    if err := c.Bind(user); err != nil {
        return err
    }

    var dbUser models.User
    config.DB.Where("email = ?", user.Email).First(&dbUser)

    if dbUser.ID == 0 {
        return c.JSON(http.StatusUnauthorized, "User tidak ditemukan")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
        return c.JSON(http.StatusUnauthorized, "Password salah")
    }

    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = dbUser.ID
    claims["user_type"] = "user"
    claims["exp"] = time.Now().AddDate(0, 1, 0).Unix() // Token berlaku 1 bulan dari sekarang
    t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET2")))
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, map[string]string{
        "token": t,
    })
}
