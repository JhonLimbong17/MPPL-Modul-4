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

type ReturnProductHandler struct {
	ReturnProductUsecase purchase.UseCase
}

func NewReturnProductHandler(e *echo.Echo, returnproductusecase purchase.UseCase)  {
	handler := &ReturnProductHandler{ReturnProductUsecase: returnproductusecase}

	e.GET("/return_product", handler.FetchReturnProduct)
	e.GET("/return_product/:id", handler.GetById)
	e.POST("/return_product", handler.Store)
	e.PUT("/return_product/:id", handler.Update)
	e.DELETE("/return_product/:id", handler.Delete)
	e.PUT("/return_product/confirm-return-product", handler.confirm)
}

func (ph *ReturnProductHandler) confirm(c echo.Context) error {
	//confirm := c.Param("")
	return nil
}

func (ph *ReturnProductHandler) FetchReturnProduct(c echo.Context) error {
	listEl, err := ph.ReturnProductUsecase.FetchReturnProduct()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *ReturnProductHandler) GetById(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.ReturnProductUsecase.GetByIdReturnProduct(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *ReturnProductHandler) Update(c echo.Context) error {
	var returnproduct_ ReturnProduct

	err := c.Bind(&returnproduct_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForReturnProductValid(&returnproduct_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ReturnProductUsecase.UpdateReturnProduct(&returnproduct_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, returnproduct_)
}

func (ph *ReturnProductHandler) Store(c echo.Context) error {
	var returnproduct_ ReturnProduct

	err := c.Bind(&returnproduct_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForReturnProductValid(&returnproduct_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ReturnProductUsecase.StoreReturnProduct(&returnproduct_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, returnproduct_)
}

func (ph *ReturnProductHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.ReturnProductUsecase.DeleteReturnProduct(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForReturnProductValid(p *ReturnProduct) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}