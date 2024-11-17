package controllers

import (
	"donasiPohon/config"
	"donasiPohon/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateDonation(c echo.Context) error {
    donation := new(models.Donation)
    if err := c.Bind(donation); err != nil {
        return err
    }

    var campaign models.Campaign
    if err := config.DB.First(&campaign, donation.CampaignID).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Campaign not found")
    }

    campaign.JumlahDonasi += donation.Jumlah

    if err := config.DB.Save(&campaign).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to update campaign")
    }

    if err := config.DB.Create(&donation).Error; err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    return c.JSON(http.StatusOK, donation)
}

