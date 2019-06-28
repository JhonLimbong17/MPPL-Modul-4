package repository

import (
	. "MPPL-Modul-4-master/models/purchase"
	"MPPL-Modul-4-master/purchase"
	"github.com/jinzhu/gorm"
)

type purchaseProductRepository struct {
	Conn *gorm.DB
}

func NewPurchaseProductRepository(Conn *gorm.DB) purchase.RepositoryPurchaseProduct {
	return &purchaseProductRepository{Conn}
}

func (pr *purchaseProductRepository) Fetch() (res []*PurchaseProduct, err error) {
	var purchaseProduct []*PurchaseProduct
	err = pr.Conn.Find(&purchaseProduct).Error

	if err != nil {
		return nil, err
	}

	return purchaseProduct, nil
}

func (pr *purchaseProductRepository) GetById(id uint) (*PurchaseProduct, error) {
	var purchaseProduct_ PurchaseProduct
	err := pr.Conn.Find(&purchaseProduct_, id).Error

	if err != nil {
		return nil, err
	}

	return &purchaseProduct_, nil
}

func (pr *purchaseProductRepository) Update(p *PurchaseProduct) error {
	var purchaseProduct_ PurchaseProduct
	pr.Conn.Find(purchaseProduct_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *purchaseProductRepository) Store(p *PurchaseProduct) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *purchaseProductRepository) Delete(id uint) error {
	var purchaseProduct_ PurchaseProduct
	pr.Conn.Find(&purchaseProduct_)
	err := pr.Conn.Delete(&purchaseProduct_).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *purchaseProductRepository) ConfirmStatusPayment(pp *PurchaseProduct) error {
	pp.TransactionStatus = "confirmed"
	err := pr.Conn.Save(&pp).Error

	if err != nil {
		return err
	}

	return nil
}