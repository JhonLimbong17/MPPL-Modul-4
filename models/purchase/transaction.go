package purchase

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Transaction struct {
	gorm.Model

	ProductURL string
	ProductName string
	ProductPrice int
	ProductQuantity int
	CourierName string
	CourierPrice	int
	DateOrder time.Time
	UserName string
	UserTelephone string
	UserAddress string
	BankAccount string

	PayBefore time.Time
}

