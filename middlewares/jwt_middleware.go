// middlewares/jwt_middleware.go
package middlewares

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
    secret := os.Getenv("JWT_SECRET")
    return middleware.JWTWithConfig(middleware.JWTConfig{
        SigningKey: []byte(secret),
    })
}

func JWTMiddleware2() echo.MiddlewareFunc {
    secret := os.Getenv("JWT_SECRET2")
    return middleware.JWTWithConfig(middleware.JWTConfig{
        SigningKey: []byte(secret),
    })
}

// Middleware untuk membatasi akses hanya untuk Komunitas

// Middleware to restrict access to komunitas only
func RestrictToKomunitas(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(jwt.MapClaims)
        
        // Cek apakah token memiliki komunitas_id
        komunitasID, ok := claims["komunitas_id"]
        if !ok || komunitasID == nil {
            return c.JSON(http.StatusForbidden, map[string]string{
                "message": "Access restricted to komunitas only",
            })
        }

        // Lanjut ke handler berikutnya jika komunitas_id ada
        return next(c)
    }
}

// Middleware untuk membatasi akses hanya untuk User
func RestrictToUser(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        userType := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_type"].(string)
        if userType != "user" {
            return c.JSON(http.StatusForbidden, map[string]string{"error": "Akses hanya untuk User"})
        }
        return next(c)
    }
}
