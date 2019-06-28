package http

import (
	"MPPL-Modul-4-master/purchase"
	. "MPPL-Modul-4-master/purchase/delivery/utils"
	. "MPPL-Modul-4-master/models/purchase"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type RefundMoneyHandler struct {
	RefundMoneyUsecase purchase.UseCase
}

func NewRefundMoneyHandler(e *echo.Echo, refundmoneyusecase purchase.UseCase)  {
	handler := &RefundMoneyHandler{RefundMoneyUsecase: refundmoneyusecase}

	e.GET("/refund_money", handler.FetchRefundMoney)
	e.GET("/refund_money/:id", handler.GetById)
	e.POST("/refund_money", handler.Store)
	e.PUT("/redund_money/:id", handler.Update)
	e.DELETE("refund_money/:id", handler.Delete)
}

func (ph *RefundMoneyHandler) FetchRefundMoney(c echo.Context) error {
	listEl, err := ph.RefundMoneyUsecase.FetchRefundMoney()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *RefundMoneyHandler) GetById(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.RefundMoneyUsecase.GetByIdRefundMoney(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *RefundMoneyHandler) Update(c echo.Context) error {
	var refundmoney_ RefundMoney

	err := c.Bind(&refundmoney_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForRefundMoneyValid(&refundmoney_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.RefundMoneyUsecase.UpdateRefundMoney(&refundmoney_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, refundmoney_)
}

func (ph *RefundMoneyHandler) Store(c echo.Context) error {
	var refundmoney_ RefundMoney

	err := c.Bind(&refundmoney_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForRefundMoneyValid(&refundmoney_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.RefundMoneyUsecase.StoreRefundMoney(&refundmoney_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, refundmoney_)
}

func (ph *RefundMoneyHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.RefundMoneyUsecase.DeleteRefundMoney(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForRefundMoneyValid(p *RefundMoney) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)

	if err != nil {
		return false, err
	}

	return true, nil
}