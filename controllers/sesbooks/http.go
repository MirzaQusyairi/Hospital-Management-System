package sesbooks

import (
	"net/http"

	"Hospital-Management-System/business/sesbooks"
	"Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/sesbooks/request"
	"Hospital-Management-System/controllers/sesbooks/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SesbookController struct {
	sesbookService sesbooks.Service
}

func NewControllerSesbook(serv sesbooks.Service) *SesbookController {
	return &SesbookController{
		sesbookService: serv,
	}
}

func (ctrl *SesbookController) AddSessionBook(c echo.Context) error {

	registerReq := request.Sesbooks{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.sesbookService.AddSessionBook(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *SesbookController) AllSessionBook(c echo.Context) error {

	result, err := ctrl.sesbookService.AllSessionBook()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromSesbookListDomain(result))

}

func (ctrl *SesbookController) Update(c echo.Context) error {

	updateReq := request.Sesbooks{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	getData, _ := ctrl.sesbookService.SessionBookByID(id)
	result, err := ctrl.sesbookService.Update(id, updateReq.ToDomain())
	result.ID = getData.ID

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdateSesbook(result))

}

func (ctrl *SesbookController) SessionBookByID(c echo.Context) error {

	sesbookID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.sesbookService.SessionBookByID(sesbookID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllSesbook(result))
}

func (ctrl *SesbookController) Delete(c echo.Context) error {

	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.sesbookService.Delete(deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)
}
