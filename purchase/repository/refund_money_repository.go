package repository

import (
	. "MPPL-Modul-4-master/models/purchase"
	"MPPL-Modul-4-master/purchase"
	"github.com/jinzhu/gorm"
)

type refundMoneyRepository struct {
	Conn *gorm.DB
}

func NewRefundMoneyRepository(Conn *gorm.DB) purchase.RepositoryRefundMoney {
	return &refundMoneyRepository{Conn}
}

func (pr *refundMoneyRepository) Fetch() (res []*RefundMoney, err error) {
	var refundMoney []*RefundMoney
	err = pr.Conn.Find(&refundMoney).Error

	if err != nil {
		return nil, err
	}

	return refundMoney, nil
}

func (pr *refundMoneyRepository) GetById(id uint) (*RefundMoney, error) {
	var refundMoney_ RefundMoney
	err := pr.Conn.Find(&refundMoney_, id).Error

	if err != nil {
		return nil, err
	}

	return &refundMoney_, nil
}

func (pr *refundMoneyRepository) Update(p *RefundMoney) error {
	var refundMoney_ RefundMoney
	pr.Conn.Find(refundMoney_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *refundMoneyRepository) Store(p *RefundMoney) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *refundMoneyRepository) Delete(id uint) error {
	var refundMoney_ RefundMoney
	pr.Conn.Find(&refundMoney_)
	err := pr.Conn.Delete(&refundMoney_).Error

	if err != nil {
		return err
	}

	return nil
}