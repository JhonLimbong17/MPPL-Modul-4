package purchase

import "github.com/jinzhu/gorm"

type ProductExchange struct {
	gorm.Model

	idReturnProduct         int `json:"id_return_product"`
	ShippingAddress         string
	ReceiverProductExchange string
	ContactProductExchange  string
	StatusProductExchange   string
}
