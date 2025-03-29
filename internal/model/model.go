package model

import "gorm.io/gorm"

type Filer struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	PhoneNumber string `gorm:"type:varchar(20);not null;unique" json:"phone_number"`
	FbLink      string `gorm:"type:varchar(255)" json:"fb_link"`
	CreatedAt   int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   int64  `gorm:"autoUpdateTime" json:"updated_at"`
}

type Injury struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	Age       string `gorm:"type:varchar(3);not null" json:"age"`
	Address   string `gorm:"type:text;not null" json:"address"`
	Phone     string `gorm:"type:varchar(20);not null" json:"phone"`
	Shelter   string `gorm:"type:varchar(255)" json:"shelter"`
	FilerID   uint   `gorm:"not null;index" json:"filer_id"`
	Filer     Filer  `gorm:"foreignKey:FilerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"filer"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updated_at"`
}

type Death struct {
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string `gorm:"type:varchar(255);not null" json:"name"`
	Age           string `gorm:"type:varchar(3);not null" json:"age"`
	Address       string `gorm:"type:text;not null" json:"address"`
	DeadBodyPlace string `gorm:"type:varchar(255);not null" json:"dead_body_place"`
	FilerID       uint   `gorm:"not null;index" json:"filer_id"`
	Filer         Filer  `gorm:"foreignKey:FilerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"filer"`
	CreatedAt     int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     int64  `gorm:"autoUpdateTime" json:"updated_at"`
}

type Damage struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Type       string `gorm:"type:varchar(255);not null" json:"type"` // Home, Religious Place
	Address    string `gorm:"type:text;not null" json:"address"`
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
	DamageRate int    `gorm:"not null" json:"damage_rate"`
	FilerID    uint   `gorm:"not null;index" json:"filer_id"`
	Filer      Filer  `gorm:"foreignKey:FilerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"filer"`
	CreatedAt  int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  int64  `gorm:"autoUpdateTime" json:"updated_at"`
}

type Donator struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
	Phone      string `gorm:"type:varchar(20);not null" json:"phone"`
	Amount     int    `gorm:"not null" json:"amount"`
	ScreenShot string `gorm:"type:varchar(255)" json:"screenshot"`
	CreatedAt  int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  int64  `gorm:"autoUpdateTime" json:"updated_at"`
}

type DonationFormByFundRaiser struct {
	ID                  uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                string `gorm:"type:varchar(255);not null" json:"name"`
	Phone               string `gorm:"type:varchar(20);not null" json:"phone"`
	FbLink              string `gorm:"type:varchar(255)" json:"fb_link"`
	PaymentMethod       string `gorm:"type:varchar(255);not null" json:"payment_method"`
	ProofOfDonationDone int    `gorm:"not null" json:"proof_of_donation_done"`
	CreatedAt           int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           int64  `gorm:"autoUpdateTime" json:"updated_at"`
}
