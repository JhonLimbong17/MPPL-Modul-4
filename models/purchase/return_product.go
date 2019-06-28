package purchase

import "github.com/jinzhu/gorm"

type ReturnProduct struct {
	gorm.Model

	IDPurchaseProduct int `json:"id_purchase_product"`
	EmailReturnProduct string
	Shipping string
	ResiNumber string
	DeliveryName string
	//reason string
	ImageUrl string
	reason string
}
