package main

import (
	"MPPL-Modul-4-master/models/purchase"
	"MPPL-Modul-4-master/purchase/delivery/http"
	"MPPL-Modul-4-master/purchase/repository"
	"MPPL-Modul-4-master/purchase/usecase"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	//http.Handle("/", http.FileServer(http.Dir("/")))
//	http.HandleFunc("/upload", controller.UploadFile())
	fmt.Println("MEME ", viper.GetString("server.address"))
	log.Println("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading")
	//log.Fatal(http.ListenAndServe(":8080", nil))

	db, err := gorm.Open("mysql", "root:@/modul_purchase?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	db.AutoMigrate(&purchase.PurchaseProduct{}, &purchase.ProductExchange{}, &purchase.RefundMoney{}, &purchase.ReturnProduct{}, &purchase.Feedback{}, &purchase.Notifikasi{}, &purchase.EvidenceTransfer{}, &purchase.Transaction{})

	fr := repository.NewFeedbackRepository(db)
	nr := repository.NewNotifikasiRepository(db)
	per := repository.NewProductExchangeRepository(db)
	ppr := repository.NewPurchaseProductRepository(db)
	rmr := repository.NewRefundMoneyRepository(db)
	rpr := repository.NewReturnProductRepository(db)
	etr := repository.NewEvidenceTransferRepository(db)
	tr := repository.NewTransactionRepository(db)

	pu := usecase.NewPurchaseUsecase(ppr, per, rmr, rpr, fr, nr, etr, tr)
	//	ep := repository.NewExchangeProduct(dbConn)

	e := echo.New()
	http.NewPurchaseProductHandler(e, pu)
	http.NewProductExchangeHandler(e, pu)
	http.NewRefundMoneyHandler(e, pu)
	http.NewReturnProductHandler(e, pu)
	http.NewFeedbackHandler(e, pu)
	http.NewNotifikasiHandler(e, pu)
	http.NewEvidenceTransferHandler(e, pu)
	http.NewTransactionHandler(e, pu)

	_ = e.Start(viper.GetString("server.address"))
	fmt.Println("MEME ", viper.GetString("server.address"))

}
