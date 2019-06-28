package purchase

import (
	. "MPPL-Modul-4-master/models/purchase"
)

type UseCase interface {
	FetchProductExchange() (res []*ProductExchange, err error)
	GetByIdProductExchange(id uint) (*ProductExchange, error)
	UpdateProductExchange(p *ProductExchange) error
	StoreProductExchange(p *ProductExchange) error
	DeleteProductExchange(id uint) error

	FetchFeedback() (res []*Feedback, err error)
	GetByIdFeedback(id uint) (*Feedback, error)
	UpdateFeedback(p *Feedback) error
	StoreFeedback(p *Feedback) error
	DeleteFeedback(id uint) error

	FetchPurchaseProduct() (res []*PurchaseProduct, err error)
	GetByIdPurchaseProduct(id uint) (*PurchaseProduct, error)
	UpdatePurchaseProduct(p *PurchaseProduct) error
	StorePurchaseProduct(p *PurchaseProduct) error
	DeletePurchaseProduct(id uint) error
	ConfirmStatusPayment(id uint) error

	FetchRefundMoney() (res []*RefundMoney, err error)
	GetByIdRefundMoney(id uint) (*RefundMoney, error)
	UpdateRefundMoney(p *RefundMoney) error
	StoreRefundMoney(p *RefundMoney) error
	DeleteRefundMoney(id uint) error

	FetchReturnProduct() (res []*ReturnProduct, err error)
	GetByIdReturnProduct(id uint) (*ReturnProduct, error)
	UpdateReturnProduct(p *ReturnProduct) error
	StoreReturnProduct(p *ReturnProduct) error
	DeleteReturnProduct(id uint) error

	FetchNotifikasi() (res []*Notifikasi, err error)
	GetByIdNotifikasi(id uint) (*Notifikasi, error)
	UpdateNotifikasi(p *Notifikasi) error
	StoreNotifikasi(p *Notifikasi) error
	DeleteNotikasi(id uint) error

	FetchEvidenceTransfer() (res []*EvidenceTransfer, err error)
	GetByIdEvidenceTransfer(id uint) (*EvidenceTransfer, error)
	UpdateEvidenceTransfer(p *EvidenceTransfer) error
	StoreEvidenceTransfer(p *EvidenceTransfer) error
	DeleteEvidenceTransfer(id uint) error
	UploadEvidenceTransfer(fileBytes []byte, fileType string) (string, error)

	SendNotification(email string) (string error)

	GetByIdTransaction(id uint) (*Transaction,	 error)
	FetchTransaction() (res []*Transaction, err error)
}