package purchase

import (
	. "MPPL-Modul-4-master/models/purchase"
)

type RepositoryProductExchange interface {
	Fetch() (res []*ProductExchange, err error)
	GetById(id uint) (*ProductExchange, error)
	Update(p *ProductExchange) error
	Store(p *ProductExchange) error
	Delete(id uint) error
}

type RepositoryFeedback interface {
	Fetch() (res []*Feedback, err error)
	GetById(id uint) (*Feedback, error)
	Update(p *Feedback) error
	Store(p *Feedback) error
	Delete(id uint) error
}

type RepositoryPurchaseProduct interface {
	Fetch() (res []*PurchaseProduct, err error)
	GetById(id uint) (*PurchaseProduct, error)
	Update(p *PurchaseProduct) error
	Store(p *PurchaseProduct) error
	Delete(id uint) error
	ConfirmStatusPayment(pp *PurchaseProduct) error
}

type RepositoryRefundMoney interface {
	Fetch() (res []*RefundMoney, err error)
	GetById(id uint) (*RefundMoney, error)
	Update(p *RefundMoney) error
	Store(p *RefundMoney) error
	Delete(id uint) error
}

type RepositoryReturnProduct interface {
	Fetch() (res []*ReturnProduct, err error)
	GetById(id uint) (*ReturnProduct, error)
	Update(p *ReturnProduct) error
	Store(p *ReturnProduct) error
	Delete(id uint) error
}

type RepositoryNotifikasi interface {
	Fetch() (res []*Notifikasi, err error)
	GetById(id uint) (*Notifikasi, error)
	Update(p *Notifikasi) error
	Store(p *Notifikasi) error
	Delete(id uint) error

	SendNotification(email string) (string, error)
}

type RepositoryEvidenceTransfer interface {
	Fetch() (res []*EvidenceTransfer, err error)
	GetById(id uint) (*EvidenceTransfer, error)
	Update(p *EvidenceTransfer) error
	Store(p *EvidenceTransfer) error
	Delete(id uint) error
}

type RepositoryTransaction interface {
	GetById(id uint) (*Transaction, error)
	Fetch() (res []*Transaction, err error)
}