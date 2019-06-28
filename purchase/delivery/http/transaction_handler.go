package http

import (
	"MPPL-Modul-4-master/purchase"
	. "MPPL-Modul-4-master/purchase/delivery/utils"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type TransactionHandler struct {
	TransactionUseCase purchase.UseCase
}

func NewTransactionHandler(e *echo.Echo, transactionusecase purchase.UseCase){
	handler := &TransactionHandler{TransactionUseCase:transactionusecase}

	e.GET("/transactions/:id", handler.GetById)
	e.GET("/transactions", handler.Fetch)

}

func (ph *TransactionHandler) Fetch(c echo.Context) error {
	listEl, err := ph.TransactionUseCase.FetchTransaction()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *TransactionHandler) GetById(c echo.Context) error {

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.TransactionUseCase.GetByIdTransaction(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}
