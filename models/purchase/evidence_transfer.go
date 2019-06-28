package purchase

import "github.com/jinzhu/gorm"

type EvidenceTransfer struct {
	gorm.Model

	//IDEvidenceTransfer uint

	IDPurchaseProduct uint `json:"id_purchase_product""`
	IDTransaction uint `json:"id_transaction"`

	imageURL string
}
