package models

type FQDN struct {
	ID   uint   `gorm:"primaryKey"`
	FQDN string `gorm:"unique"`
	IPs  []IP   `gorm:"foreignKey:FQDNID"`
}

type IP struct {
	ID      uint   `gorm:"primaryKey"`
	FQDNID  uint   `gorm:"index"`
	Address string `gorm:"index"`
}
