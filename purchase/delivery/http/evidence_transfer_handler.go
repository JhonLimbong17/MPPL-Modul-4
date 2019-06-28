package http

import (
	. "MPPL-Modul-4-master/models/purchase"
	"MPPL-Modul-4-master/purchase"
	. "MPPL-Modul-4-master/purchase/delivery/utils"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"net/http"
	"strconv"
)

const maxUploadSize = 3 * 1024 * 1024 // 3 MB
const uploadPath = "./files"

type EvidenceTransferHandler struct {
	EvidenceTransferUsecase purchase.UseCase
}

func NewEvidenceTransferHandler(e *echo.Echo, evidencetransferusecase purchase.UseCase)  {
	handler := &EvidenceTransferHandler{EvidenceTransferUsecase: evidencetransferusecase}

	e.GET("/evidence_transfer", handler.FetchEvidenceTransfer)
	e.GET("/evidence_transfer/:id", handler.GetByIdEvidenceTransfer)
	e.POST("/evidence_transfer", handler.Store)
	e.PUT("/evidence_transfer/:id", handler.Update)
	e.DELETE("/evidence_transfer/:id", handler.Delete)
	//e.POST("/evidence-transfer/UploadImage", handler.uploadImage())
	//e.POST("evidence-transfer/uploadImage", handler.uploadImage())
	e.POST("/evidence-transfer/uploadImage", echo.WrapHandler(handler.uploadImage()))

	router := mux.NewRouter()

	//router.HandleFunc("/evidence_transfer", handler.FetchEvidenceTransferMux())

	router.HandleFunc("/evidence_transfer/uploadImage", handler.uploadImage())

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	//log.Fatal(http.ListenAndServe(":8080", router))
}

func (ph *EvidenceTransferHandler) FetchEvidenceTransfer(c echo.Context) error {
	listEl, err := ph.EvidenceTransferUsecase.FetchEvidenceTransfer()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

//func (ph *EvidenceTransferHandler)FetchEvidenceTransferMux() http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		listEl, err := ph.EvidenceTransferUsecase.FetchEvidenceTransfer()
//
//		if err != nil {
//			renderError(w, "INVALID_REQUEST_TYPE", http.StatusBadRequest)
//			return
//		}
//
//	})
//}

func (ph *EvidenceTransferHandler) GetByIdEvidenceTransfer(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.EvidenceTransferUsecase.GetByIdEvidenceTransfer(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *EvidenceTransferHandler) Store(c echo.Context) error {
	var evidenceTranser_ EvidenceTransfer

	err := c.Bind(&evidenceTranser_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForEvidenceTransferValid(&evidenceTranser_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.EvidenceTransferUsecase.StoreEvidenceTransfer(&evidenceTranser_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, evidenceTranser_)
}

func (ph *EvidenceTransferHandler)uploadImage() http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		// parse and validate file and post parameters
		//fileType := r.PostFormValue("type")
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		filetype := http.DetectContentType(fileBytes)
		switch filetype {
		case "image/jpeg", "image/jpg":
			break
		case "image/gif", "image/png":
			break
		default:
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}

		filename, err := ph.EvidenceTransferUsecase.UploadEvidenceTransfer(fileBytes, filetype)

		if err != nil {
			renderError(w, err.Error(), http.StatusInternalServerError)
		}else {
			w.Write([]byte(filename))
		}
	})
}

func (ph *EvidenceTransferHandler) Update(c echo.Context) error {
	var evidenceTransfer_ EvidenceTransfer
	
	err := c.Bind(&evidenceTransfer_)
	
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	
	if ok, err := isRequestForEvidenceTransferValid(&evidenceTransfer_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	err = ph.EvidenceTransferUsecase.UpdateEvidenceTransfer(&evidenceTransfer_)
	
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	
	return c.JSON(http.StatusCreated, evidenceTransfer_)
}

func (ph *EvidenceTransferHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.EvidenceTransferUsecase.DeleteEvidenceTransfer(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForEvidenceTransferValid(p *EvidenceTransfer) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}