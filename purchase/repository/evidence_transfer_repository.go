package repository

import (
	. "MPPL-Modul-4-master/models/purchase"
	"MPPL-Modul-4-master/purchase"
	"github.com/jinzhu/gorm"
)

type evidenceTransferRepository struct {
	Conn *gorm.DB
}

func NewEvidenceTransferRepository(Conn *gorm.DB) purchase.RepositoryEvidenceTransfer {
	return &evidenceTransferRepository{Conn}
}

func (pr *evidenceTransferRepository) Fetch() (res []*EvidenceTransfer, err error) {
	var evidenceTransfer []*EvidenceTransfer
	err = pr.Conn.Find(&evidenceTransfer).Error

	if err != nil {
		return nil, err
	}

	return evidenceTransfer, nil
}

func (pr *evidenceTransferRepository) GetById(id uint) (*EvidenceTransfer, error) {
	var evidenceTransfer_ EvidenceTransfer
	err := pr.Conn.Find(&evidenceTransfer_, id).Error

	if err != nil {
		return nil, err
	}

	return &evidenceTransfer_, nil
}

func (pr *evidenceTransferRepository) Update(p *EvidenceTransfer) error {
	var evidenceTransfer_ EvidenceTransfer
	pr.Conn.Find(evidenceTransfer_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *evidenceTransferRepository) Store(p *EvidenceTransfer) error {
	err := pr.Conn.Create(&p).Error
	if err != nil {
		return err
	}

	return nil
}

func (pr *evidenceTransferRepository) Delete(id uint) error {
	var evidenceTransfer_ EvidenceTransfer
	pr.Conn.Find(&evidenceTransfer_)
	err := pr.Conn.Delete(&evidenceTransfer_).Error

	if err != nil {
		return err
	}

	return nil
}

func NewEvidenceRepository(Conn *gorm.DB) purchase.RepositoryEvidenceTransfer {
	return &evidenceTransferRepository{Conn}
}