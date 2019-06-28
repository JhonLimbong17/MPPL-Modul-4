package purchase

import "github.com/jinzhu/gorm"

type RefundMoney struct {
	gorm.Model

	BankAccountRefund string
	ReceiverRefund string
	EmailRefund string
	AreaBank string
	StatusRefund string
	BankDestination string
	IDReturnProduct int `json:"id_return_product"`
}
