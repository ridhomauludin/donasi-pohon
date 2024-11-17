// routes/routes.go
package routes

import (
	"donasiPohon/controllers"
	"donasiPohon/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
    // Rute untuk registrasi dan login user
    e.POST("/user/register", controllers.RegisterUser)
    e.POST("/user/login", controllers.LoginUser)
    
    // Rute untuk registrasi dan login komunitas
    e.POST("/komunitas/register", controllers.RegisterKomunitas)
    e.POST("/komunitas/login", controllers.LoginKomunitas)

    // Grup rute campaign yang membutuhkan autentikasi
    campaignGroup := e.Group("/campaigns")
    campaignGroup.Use(middlewares.JWTMiddleware2())
    campaignGroup.GET("", controllers.GetCampaigns)
    campaignGroup.POST("", controllers.CreateCampaign)
    campaignGroup.DELETE("/:id", controllers.DeleteCampaign)
    campaignGroup.PUT("/:id", controllers.EditCampaign)

    // Grup rute donation yang membutuhkan autentikasi
    donationGroup := e.Group("/donations")
    donationGroup.Use(middlewares.JWTMiddleware2())
    donationGroup.POST("", controllers.CreateDonation)
}
