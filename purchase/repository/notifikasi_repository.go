package repository

import (
	. "MPPL-Modul-4-master/models/purchase"
	"MPPL-Modul-4-master/purchase"
	"github.com/jinzhu/gorm"
)

type notifikasiRepository struct {
	Conn *gorm.DB
}

func (pr *notifikasiRepository) SendNotification(email string) (string, error) {
	panic("implement me")
}

func NewNotifikasiRepository(Conn *gorm.DB) purchase.RepositoryNotifikasi {
	return &notifikasiRepository{Conn}
}

func (pr *notifikasiRepository) Delete(id uint) error {
	var notifikasi_ Notifikasi
	pr.Conn.Find(&notifikasi_)
	err := pr.Conn.Delete(&notifikasi_).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *notifikasiRepository) Fetch() (res []*Notifikasi, err error) {
	panic("implement me")
}

func (pr *notifikasiRepository) GetById(id uint) (*Notifikasi, error) {
	var notifikasi_ Notifikasi
	err := pr.Conn.Find(&notifikasi_, id).Error

	if err != nil {
		return nil, err
	}

	return &notifikasi_, nil
}

func (pr *notifikasiRepository) Store(p *Notifikasi) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *notifikasiRepository) Update(p *Notifikasi) error {
	var notifikasi_ Notifikasi
	pr.Conn.Find(notifikasi_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}