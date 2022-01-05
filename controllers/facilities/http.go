package facilities

import (
	"net/http"

	"Hospital-Management-System/business/facilities"
	"Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/facilities/request"
	response "Hospital-Management-System/controllers/facilities/res"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FaciltyController struct {
	faciltyService facilities.Service
}

func NewControllerFacilty(serv facilities.Service) *FaciltyController {
	return &FaciltyController{
		faciltyService: serv,
	}
}

func (ctrl *FaciltyController) AddFacilty(c echo.Context) error {

	registerReq := request.Facilities{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.faciltyService.AddFacilty(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}
func (ctrl *FaciltyController) AllFacilty(c echo.Context) error {

	result, err := ctrl.faciltyService.AllFacilty()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromFaciltyListDomain(result))

}

func (ctrl *FaciltyController) Update(c echo.Context) error {

	updateReq := request.Facilities{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	getData, _ := ctrl.faciltyService.FacByID(id)
	result, err := ctrl.faciltyService.Update(id, updateReq.ToDomain())
	result.ID = getData.ID

	result.Name = getData.Name

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdateDoctor(result))

}
func (ctrl *FaciltyController) FacByID(c echo.Context) error {

	facID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.faciltyService.FacByID(facID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllFacilty(result))
}
func (ctrl *FaciltyController) Delete(c echo.Context) error {

	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.faciltyService.Delete(deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)

}
