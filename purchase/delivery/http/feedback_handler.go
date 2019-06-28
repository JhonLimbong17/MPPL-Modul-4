package http

import (
	. "MPPL-Modul-4-master/models/purchase"
	"MPPL-Modul-4-master/purchase"
	. "MPPL-Modul-4-master/purchase/delivery/utils"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type FeedbackHandler struct {
	FeedbackUsecase purchase.UseCase

}

func NewFeedbackHandler(e *echo.Echo, feedbackusecase purchase.UseCase)  {
	handler := &FeedbackHandler{FeedbackUsecase: feedbackusecase}

	e.GET("/feedback", handler.FetchBrand)
	e.GET("/feedback/:id", handler.GetById)
	e.POST("/feedback", handler.Store)
	e.PUT("feedback/:id", handler.Update)
	e.DELETE("/feedback/:id", handler.Delete)
}

func (ph *FeedbackHandler) FetchBrand(c echo.Context) error {
	listEl, err := ph.FeedbackUsecase.FetchFeedback()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *FeedbackHandler) GetById(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.FeedbackUsecase.GetByIdFeedback(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *FeedbackHandler) Update(c echo.Context) error {
	var feedback_ Feedback

	err := c.Bind(&feedback_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForFeedbackValid(&feedback_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.FeedbackUsecase.UpdateFeedback(&feedback_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, feedback_)
}

func (ph *FeedbackHandler) Store(c echo.Context) error {
	var feedback_ Feedback

	err := c.Bind(&feedback_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForFeedbackValid(&feedback_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.FeedbackUsecase.StoreFeedback(&feedback_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, feedback_)
}

func (ph *FeedbackHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.FeedbackUsecase.DeleteFeedback(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForFeedbackValid(p *Feedback) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}