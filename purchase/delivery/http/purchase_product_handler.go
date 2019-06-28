package http

import (
	. "MPPL-Modul-4-master/models/purchase"
	"MPPL-Modul-4-master/purchase"
	. "MPPL-Modul-4-master/purchase/delivery/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type PurchaseProductHandler struct {
	PurchaseProductUseCase purchase.UseCase
}

func NewPurchaseProductHandler(e *echo.Echo, purchaseproductusecase purchase.UseCase)  {
	handler := &PurchaseProductHandler{PurchaseProductUseCase: purchaseproductusecase}

	e.GET("/purchase_product", handler.FetchPurchaseProduct)
	e.GET("/purchase_product/:id", handler.GetById)
	e.POST("/purchase_product", handler.Store)
	e.PUT("/purchase_product/:id", handler.Update)
	e.DELETE("/purchase_product/:id", handler.Delete)
	e.GET("/purchase_product/update_status/:id", handler.UpdateStatusConfirm)

	//router := mux.NewRouter()

	//router.HandleFunc("/purchase_product/update_status/{id}", handler.updateStatus)

	//log.Fatal(http.ListenAndServe(":8080", router))
}

//func (ph *PurchaseProductHandler) updateStatus(c echo.Context) error {
//	id_, err := strconv.Atoi(c.Param("id"))
//	id := uint(id_)
//
//	el, err := ph.PurchaseProductUseCase.GetByIdPurchaseProduct(id)
//
//	confirm:= c.Param("TransactionStatus")
//
//}

func (ph *PurchaseProductHandler) FetchPurchaseProduct(c echo.Context) error {
	listEl, err := ph.PurchaseProductUseCase.FetchPurchaseProduct()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *PurchaseProductHandler) GetById(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.PurchaseProductUseCase.GetByIdPurchaseProduct(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *PurchaseProductHandler) Update(c echo.Context) error {
	var purchaseproduct_ PurchaseProduct

	err := c.Bind(&purchaseproduct_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForPurchaseProductValid(&purchaseproduct_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.PurchaseProductUseCase.UpdatePurchaseProduct(&purchaseproduct_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, purchaseproduct_)
}

func (ph *PurchaseProductHandler) UpdateStatusConfirm(c echo.Context) error {
	//var purchaseproduct_ PurchaseProduct
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	//err := c.Bind(&purchaseproduct_)

	if err != nil {
		return c.JSON(200, "Confirmation failed")
	}

	//if ok, err := isRequestForPurchaseProductValid(&purchaseproduct_); !ok {
	//	return c.JSON(http.StatusBadRequest, err.Error())
	//}

	err = ph.PurchaseProductUseCase.ConfirmStatusPayment(id)

	if err != nil {
		return c.JSON(200, "Confirmation failed")
	}

	return c.JSON(200, "Confirmation success")
}

func (ph *PurchaseProductHandler) Store(c echo.Context) error {
	var purchaseproduct_ PurchaseProduct

	err := c.Bind(&purchaseproduct_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForPurchaseProductValid(&purchaseproduct_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.PurchaseProductUseCase.StorePurchaseProduct(&purchaseproduct_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, purchaseproduct_)
}

func (ph *PurchaseProductHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.PurchaseProductUseCase.DeletePurchaseProduct(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForPurchaseProductValid(p *PurchaseProduct) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}