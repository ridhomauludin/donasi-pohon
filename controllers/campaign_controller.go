// controllers/campaign_controller.go
package controllers

import (
	"donasiPohon/config"
	"donasiPohon/models"
	"donasiPohon/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)


func GetCampaigns(c echo.Context) error {
    var campaigns []models.Campaign
    config.DB.Find(&campaigns)
    return c.JSON(http.StatusOK, campaigns)
}


func CreateCampaign(c echo.Context) error {
    campaign := new(models.Campaign)
    if err := c.Bind(campaign); err != nil {
        return err
    }

    // Panggil API Gemini untuk menghasilkan tujuan kampanye

    prompt := "dari judul " + campaign.Judul +" dan deskripsi " + campaign.Deskripsi + "buatkan saya tujuannya"

    goal, err := utils.GenerateContent(prompt)
    if err != nil {
        log.Printf("gagal saat generate content: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }

    log.Printf("prompt: %v", prompt)
    log.Printf("result prompt: %v", goal)

    campaign.Tujuan = goal // Set tujuan kampanye yang dihasilkan

    if err := config.DB.Create(&campaign).Error; err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    return c.JSON(http.StatusOK, campaign)
}

func DeleteCampaign(c echo.Context) error {
    id := c.Param("id")
    var campaign models.Campaign


    if err := config.DB.Where("id = ?", id).First(&campaign).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Campaign tidak ditemukan")
    }


    if err := config.DB.Delete(&campaign).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, "Gagal menghapus campaign")
    }


    return c.JSON(http.StatusOK, "Campaign berhasil dihapus")
}

func EditCampaign(c echo.Context) error {
    id := c.Param("id") // Mendapatkan ID campaign dari parameter URL
    var campaign models.Campaign

    // Cari campaign berdasarkan ID
    if err := config.DB.Where("id = ?", id).First(&campaign).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{
            "message": "Campaign tidak ditemukan",
        })
    }

    // Bind data dari body request
    updatedData := new(models.Campaign)
    if err := c.Bind(updatedData); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "message": "Gagal memproses data",
        })
    }

    // Update hanya field yang diterima
    if updatedData.Judul != "" {
        campaign.Judul = updatedData.Judul
    }
    if updatedData.Deskripsi != "" {
        campaign.Deskripsi = updatedData.Deskripsi
    }
    if updatedData.TargetDonasi != 0 {
        campaign.TargetDonasi = updatedData.TargetDonasi
    }
    if updatedData.Status != "" {
        campaign.Status = updatedData.Status
    }
    if !updatedData.TanggalSelesai.IsZero() {
        campaign.TanggalSelesai = updatedData.TanggalSelesai
    }

    // Simpan perubahan ke database
    if err := config.DB.Save(&campaign).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "message": "Gagal mengupdate campaign",
        })
    }

    return c.JSON(http.StatusOK, campaign)
}

