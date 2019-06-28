package repository

import (
	"github.com/jinzhu/gorm"
	. "MPPL-Modul-4-master/models/purchase"
	. "MPPL-Modul-4-master/purchase"
)

type transactionRepository struct {
	Conn *gorm.DB
}

func NewTransactionRepository(Conn *gorm.DB) RepositoryTransaction{
	return &transactionRepository{Conn}
}

func (pr *transactionRepository) Fetch() (res []*Transaction, err error) {
	var transaction []*Transaction
	err = pr.Conn.Find(&transaction).Error

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (pr *transactionRepository) GetById(id uint) (*Transaction, error) {
	var transaction_ Transaction
	err := pr.Conn.Find(&transaction_, id).Error

	if err != nil {
		return nil, err
	}

	return &transaction_, nil
}

func (pr *transactionRepository) Update (p *Transaction) error {
	var transaction_ Transaction
	pr.Conn.Find(transaction_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *transactionRepository) Store (p *Transaction) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *transactionRepository) Delete (id uint) error {
	var transaction_ Transaction
	pr.Conn.Find(&transaction_)
	err := pr.Conn.Delete(&transaction_).Error

	if err != nil {
		return err
	}

	return nil
}