package purchase

import (
	"github.com/jinzhu/gorm"
)

type PurchaseProduct struct {
	gorm.Model

	IDTransaction int `json:"id_transaction"`
	PaymentStatus string
	DeliveryStatus string
	TransactionStatus string
	ResiNumber string
	EvidenceTransfer string
}
