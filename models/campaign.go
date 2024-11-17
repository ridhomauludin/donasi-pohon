package models

import "time"

type Campaign struct {
	ID            uint      `gorm:"primaryKey"`
	KomunitasID   uint      `json:"komunitas_id"`
	Judul         string    `json:"judul"`
	Deskripsi     string    `json:"deskripsi"`
	JenisDonasi   string    `json:"jenis_donasi"`
	TargetDonasi  int       `json:"target_donasi"`
	JumlahDonasi  int       `json:"jumlah_donasi"`
	Status        string    `json:"status"`
	TanggalDibuat time.Time `json:"tanggal_dibuat"`
	TanggalSelesai time.Time `json:"tanggal_selesai"`
	Tujuan 		string `json:"tujuan"`
}
