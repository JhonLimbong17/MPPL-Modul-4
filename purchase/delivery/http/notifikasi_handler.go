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

type NotifikasiHandler struct {
	NotifikasiUsecase purchase.UseCase
}

func NewNotifikasiHandler(e *echo.Echo, notifikasiusecase purchase.UseCase)  {
	handler := &NotifikasiHandler{NotifikasiUsecase: notifikasiusecase}

	e.GET("/notifikasi", handler.FetchNotifikasi)
	e.GET("/notifikasi/:id", handler.GetById)
	e.POST("/notifikasi", handler.Store)
	e.PUT("/notifikasi/:id", handler.Update)
	e.DELETE("/notifikasi/:id", handler.Delete)
}

func (ph *NotifikasiHandler) FetchNotifikasi(c echo.Context) error {
	listEl, err := ph.NotifikasiUsecase.FetchNotifikasi()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *NotifikasiHandler) GetById(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.NotifikasiUsecase.GetByIdNotifikasi(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *NotifikasiHandler) Update(c echo.Context) error {
	var notifikasi_ Notifikasi
	
	err := c.Bind(&notifikasi_)
	
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	
	if ok, err := isRequestForNotifikasiValid(&notifikasi_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	err = ph.NotifikasiUsecase.UpdateNotifikasi(&notifikasi_)
	
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	
	return c.JSON(http.StatusCreated, notifikasi_)
}

func (ph *NotifikasiHandler) Store(c echo.Context) error {
	var notifikasi_ Notifikasi

	err := c.Bind(&notifikasi_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForNotifikasiValid(&notifikasi_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.NotifikasiUsecase.StoreNotifikasi(&notifikasi_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, notifikasi_)
}

func (ph *NotifikasiHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.NotifikasiUsecase.DeleteNotikasi(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForNotifikasiValid(p *Notifikasi) (bool, error) {
	validae := validator.New()

	err := validae.Struct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}