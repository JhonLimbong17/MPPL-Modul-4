package repository

import (
	. "MPPL-Modul-4-master/models/purchase"
	"MPPL-Modul-4-master/purchase"
	"github.com/jinzhu/gorm"
)

type productExchangeRepository struct {
	Conn *gorm.DB
}

func NewProductExchangeRepository(Conn *gorm.DB) purchase.RepositoryProductExchange {
	return &productExchangeRepository{Conn}
}

func (pr *productExchangeRepository) Fetch() (res []*ProductExchange, err error) {
	var productExchange []*ProductExchange
	err = pr.Conn.Find(&productExchange).Error

	if err != nil {
		return nil, err
	}

	return productExchange, nil
}

func (pr *productExchangeRepository) GetById(id uint) (*ProductExchange, error) {
	var productExchange_ ProductExchange
	err := pr.Conn.Find(&productExchange_, id).Error

	if err != nil {
		return nil, err
	}

	return &productExchange_, nil
}

func (pr *productExchangeRepository) Update(p *ProductExchange) error {
	var productExchange_ ProductExchange
	pr.Conn.Find(productExchange_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *productExchangeRepository) Store(p *ProductExchange) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *productExchangeRepository) Delete(id uint) error {
	var productExchange_ ProductExchange
	pr.Conn.Find(&productExchange_)
	err := pr.Conn.Delete(&productExchange_).Error

	if err != nil {
		return err
	}

	return nil
}