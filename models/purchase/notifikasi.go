package purchase

import "github.com/jinzhu/gorm"

type Notifikasi struct {
	gorm.Model

	title string
	body string
}