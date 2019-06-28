package usecase

import (
	. "MPPL-Modul-4-master/purchase"
	. "MPPL-Modul-4-master/models/purchase"
	"crypto/rand"
	"errors"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type purchaseProductUseCase struct {
	repoPurchaseProduct RepositoryPurchaseProduct
	repoProductExchange RepositoryProductExchange
	repoRefundMoney RepositoryRefundMoney
	repoReturnProduct RepositoryReturnProduct
	repoFeedback RepositoryFeedback
	repoNotifikasi RepositoryNotifikasi
	repoEvidenceTransfer RepositoryEvidenceTransfer
	repoTransaction RepositoryTransaction
}

func (pu *purchaseProductUseCase) GetByIdTransaction(id uint) (*Transaction, error) {
	return pu.repoTransaction.GetById(id)
}

func (pu *purchaseProductUseCase) FetchTransaction() (res []*Transaction, err error) {
	res, err = pu.repoTransaction.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) SendNotification(email string) (string error) {
	panic("implement me")
}

func (pu *purchaseProductUseCase) UploadEvidenceTransfer(fileBytes []byte, fileType string) (string, error) {
	const uploadPath = "./files"

	t := time.Now()
	fileName := "evidence_transfer_"+ t.Format("01-02-2006 15:04:05")
	fileName = strings.Replace(fileName, " ", "_", 1)
	fileName = strings.Replace(fileName, ":", "-", 2)
	//fileName := randToken(12)
	fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		return "", errors.New("CANT_READ_FILE_TYPE")
	}
	newPath := filepath.Join(uploadPath, fileName + fileEndings[0])
	fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		return "", errors.New("CANT_WRITE_FILE")
	}
	defer newFile.Close() //id empotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return "", errors.New("CANT_WRITE_FILE")
	}

	return fileName, nil
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (pu *purchaseProductUseCase) FetchProductExchange() (res []*ProductExchange, err error) {
	res, err = pu.repoProductExchange.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) GetByIdProductExchange(id uint) (*ProductExchange, error) {
	res, err := pu.repoProductExchange.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) UpdateProductExchange(p *ProductExchange) error {
	return pu.repoProductExchange.Update(p)
}

func (pu *purchaseProductUseCase) StoreProductExchange(p *ProductExchange) error {
	return pu.repoProductExchange.Store(p)
}

func (pu *purchaseProductUseCase) DeleteProductExchange(id uint) error {
	return pu.repoProductExchange.Delete(id)
}


func (pu *purchaseProductUseCase) FetchFeedback() (res []*Feedback, err error){
	res, err = pu.repoFeedback.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) GetByIdFeedback(id uint) (*Feedback, error) {
	res, err := pu.repoFeedback.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) UpdateFeedback(p *Feedback) error{
	return pu.repoFeedback.Update(p)
}

func (pu *purchaseProductUseCase) StoreFeedback(p *Feedback) error{
	return pu.repoFeedback.Store(p)
}

func (pu *purchaseProductUseCase) DeleteFeedback(id uint) error{
	return pu.repoFeedback.Delete(id)
}

func (pu *purchaseProductUseCase) FetchPurchaseProduct() (res []*PurchaseProduct, err error){
	res, err = pu.repoPurchaseProduct.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) GetByIdPurchaseProduct(id uint) (*PurchaseProduct, error) {
	res, err := pu.repoPurchaseProduct.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) UpdatePurchaseProduct(p *PurchaseProduct) error{
	return pu.repoPurchaseProduct.Update(p)
}

func (pu *purchaseProductUseCase) ConfirmStatusPayment(id uint) error {
	res, err := pu.repoPurchaseProduct.GetById(id)

	if err != nil {
		return err
	}

	return pu.repoPurchaseProduct.ConfirmStatusPayment(res)
}
//
//func (pu *purchaseProductUseCase) UpdateStatusPurchaseProduct(id uint) error {
//	return pu.repoPurchaseProduct.ConfirmStatusPayment(id)
//}

func (pu *purchaseProductUseCase) StorePurchaseProduct(p *PurchaseProduct) error{
	return pu.repoPurchaseProduct.Store(p)
}

func (pu *purchaseProductUseCase) DeletePurchaseProduct(id uint) error{
	return pu.repoPurchaseProduct.Delete(id)
}

func (pu *purchaseProductUseCase) FetchRefundMoney() (res []*RefundMoney, err error){
	res, err = pu.repoRefundMoney.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) GetByIdRefundMoney(id uint) (*RefundMoney, error) {
	res, err := pu.repoRefundMoney.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) UpdateRefundMoney(p *RefundMoney) error{
	return pu.repoRefundMoney.Update(p)
}

func (pu *purchaseProductUseCase) StoreRefundMoney(p *RefundMoney) error{
	return pu.repoRefundMoney.Store(p)
}

func (pu *purchaseProductUseCase) DeleteRefundMoney(id uint) error{
	return pu.repoRefundMoney.Delete(id)
}

func (pu *purchaseProductUseCase) FetchReturnProduct() (res []*ReturnProduct, err error){
	res, err = pu.repoReturnProduct.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) GetByIdReturnProduct(id uint) (*ReturnProduct, error){
	res, err := pu.repoReturnProduct.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) UpdateReturnProduct(p *ReturnProduct) error{
	return pu.repoReturnProduct.Update(p)
}

func (pu *purchaseProductUseCase) StoreReturnProduct(p *ReturnProduct) error{
	return pu.repoReturnProduct.Store(p)
}

func (pu *purchaseProductUseCase) DeleteReturnProduct(id uint) error{
	return pu.repoReturnProduct.Delete(id)
}

func (pu *purchaseProductUseCase) FetchNotifikasi() (res []*Notifikasi, err error){
	res, err = pu.repoNotifikasi.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) GetByIdNotifikasi(id uint) (*Notifikasi, error){
	res, err := pu.repoNotifikasi.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) UpdateNotifikasi(p *Notifikasi) error{
	return pu.repoNotifikasi.Update(p)
}

func (pu *purchaseProductUseCase) StoreNotifikasi(p *Notifikasi) error{
	return pu.repoNotifikasi.Store(p)
}

func (pu *purchaseProductUseCase) DeleteNotikasi(id uint) error{
	return pu.repoNotifikasi.Delete(id)
}

//func (pu *purchaseProductUseCase) SendNotication (email string) (string, error){
//var _ gin.H
//}

func (pu *purchaseProductUseCase) FetchEvidenceTransfer() (res []*EvidenceTransfer, err error) {
	res, err = pu.repoEvidenceTransfer.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) GetByIdEvidenceTransfer(id uint) (*EvidenceTransfer, error) {
	res, err := pu.repoEvidenceTransfer.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *purchaseProductUseCase) UpdateEvidenceTransfer(p *EvidenceTransfer) error {
	return pu.repoEvidenceTransfer.Update(p)
}

func (pu *purchaseProductUseCase) StoreEvidenceTransfer(p *EvidenceTransfer) error {
	return pu.repoEvidenceTransfer.Store(p)
}

func (pu *purchaseProductUseCase) DeleteEvidenceTransfer(id uint) error {
	return pu.repoEvidenceTransfer.Delete(id)
}

func NewPurchaseUsecase(
	repoPurchaseProduct RepositoryPurchaseProduct,
	repoProductExchange RepositoryProductExchange,
	repoRefundMoney RepositoryRefundMoney,
	repoReturnProduct RepositoryReturnProduct,
	repoFeedback RepositoryFeedback,
	repoNotifikasi RepositoryNotifikasi,
	repoEvidenceTransfer RepositoryEvidenceTransfer,
	repoTransaction RepositoryTransaction) UseCase {
		return &purchaseProductUseCase{
			repoPurchaseProduct,
			repoProductExchange,
			repoRefundMoney,
			repoReturnProduct,
			repoFeedback,
			repoNotifikasi,
			repoEvidenceTransfer,
			repoTransaction,
		}
}
