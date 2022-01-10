package prescriptions

import (
	"net/http"

	"Hospital-Management-System/business/prescriptions"
	"Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/prescriptions/request"
	"Hospital-Management-System/controllers/prescriptions/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PrescriptionController struct {
	prescriptionService prescriptions.Service
}

func NewControllerPrescription(serv prescriptions.Service) *PrescriptionController {
	return &PrescriptionController{
		prescriptionService: serv,
	}
}

func (ctrl *PrescriptionController) AddPrescription(c echo.Context) error {

	registerReq := request.Prescriptions{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.prescriptionService.AddPrescription(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *PrescriptionController) AllPrescription(c echo.Context) error {

	result, err := ctrl.prescriptionService.AllPrescription()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromPrescriptionListDomain(result))

}

func (ctrl *PrescriptionController) Update(c echo.Context) error {

	updateReq := request.Prescriptions{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	getData, _ := ctrl.prescriptionService.PrescriptionByID(id)
	result, err := ctrl.prescriptionService.Update(id, updateReq.ToDomain())
	result.ID = getData.ID

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdatePrescription(result))

}

func (ctrl *PrescriptionController) PrescriptionByID(c echo.Context) error {

	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.prescriptionService.PrescriptionByID(itemID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllPrescription(result))
}

func (ctrl *PrescriptionController) PrescriptionByIDSessionBooking(c echo.Context) error {

	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.prescriptionService.PrescriptionByIDSessionBooking(itemID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromPrescriptionListDomain(result))
}

func (ctrl *PrescriptionController) Delete(c echo.Context) error {

	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.prescriptionService.Delete(deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)
}
