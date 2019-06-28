package repository

import (
	. "MPPL-Modul-4-master/models/purchase"
	"MPPL-Modul-4-master/purchase"
	"github.com/jinzhu/gorm"
)

type returnProductRepository struct {
	Conn *gorm.DB
}

func NewReturnProductRepository(Conn *gorm.DB) purchase.RepositoryReturnProduct {
	return &returnProductRepository{Conn}
}

func (pr *returnProductRepository) Fetch() (res []*ReturnProduct, err error) {
	var returnProduct []*ReturnProduct
	err = pr.Conn.Find(&returnProduct).Error

	if err != nil {
		return nil, err
	}

	return returnProduct, nil
}

func (pr *returnProductRepository) GetById(id uint) (*ReturnProduct, error) {
	var returnProduct_ ReturnProduct
	err := pr.Conn.Find(&returnProduct_, id).Error

	if err != nil {
		return nil, err
	}

	return &returnProduct_, nil
}

func (pr *returnProductRepository) Update(p *ReturnProduct) error {
	var returnProduct_ ReturnProduct
	pr.Conn.Find(returnProduct_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *returnProductRepository) Store(p *ReturnProduct) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *returnProductRepository) Delete(id uint) error {
	var returnProduct_ ReturnProduct
	pr.Conn.Find(&returnProduct_)
	err := pr.Conn.Delete(&returnProduct_).Error

	if err != nil {
		return err
	}

	return nil
}