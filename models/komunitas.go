package models

type Komunitas struct {
	ID        uint   `gorm:"primaryKey"`
	Nama      string `json:"nama"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"`
	Deskripsi string `json:"deskripsi"`
}
