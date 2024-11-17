package models

import "time"

type Donation struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       uint      `json:"user_id"`
	CampaignID   uint      `json:"campaign_id"`
	JenisDonasi  string    `json:"jenis_donasi"`
	Jumlah       int       `json:"jumlah"`
	TanggalDonasi time.Time `json:"tanggal_donasi"`
}
