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

type ProductExchangeHandler struct {
	ProductExchangeUsecase purchase.UseCase
}

func NewProductExchangeHandler(e *echo.Echo, productexchangeusecase purchase.UseCase)  {
	handler := &ProductExchangeHandler{ProductExchangeUsecase: productexchangeusecase}

	e.GET("/product_exchange", handler.FetchProductExchange)
	e.GET("/product_exchange/:id", handler.GetById)
	e.POST("/product_exchange", handler.Store)
	e.PUT("/product_exchange/:id", handler.Update)
	e.DELETE("/product_exchange/:id", handler.Delete)
}

func (ph *ProductExchangeHandler) FetchProductExchange(c echo.Context) error {
	listEl, err := ph.ProductExchangeUsecase.FetchProductExchange()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *ProductExchangeHandler) GetById(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.ProductExchangeUsecase.GetByIdProductExchange(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *ProductExchangeHandler) Update(c echo.Context) error {
	var productexchange_ ProductExchange

	err := c.Bind(&productexchange_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForProductExchangeValid(&productexchange_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ProductExchangeUsecase.UpdateProductExchange(&productexchange_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, productexchange_)
}

func (ph *ProductExchangeHandler) Store(c echo.Context) error {
	var productexchange_ ProductExchange

	err := c.Bind(&productexchange_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForProductExchangeValid(&productexchange_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ProductExchangeUsecase.StoreProductExchange(&productexchange_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, productexchange_)
}

func (ph *ProductExchangeHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.ProductExchangeUsecase.DeleteProductExchange(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForProductExchangeValid(p *ProductExchange) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}


