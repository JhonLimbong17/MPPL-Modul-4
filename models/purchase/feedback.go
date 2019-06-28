package purchase

import "github.com/jinzhu/gorm"

type Feedback struct {
	gorm.Model

	IDPurchaseProduct int `json:"id_purchase_product"`
	Rating int8
	Comment string
	ImageURL string
}