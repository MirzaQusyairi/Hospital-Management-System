package patients

import (
	"net/http"

	"Hospital-Management-System/business/patients"
	"Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/patients/request"
	response "Hospital-Management-System/controllers/patients/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PatientController struct {
	patientService patients.Service
}

func NewControllerPatient(serv patients.Service) *PatientController {
	return &PatientController{
		patientService: serv,
	}
}

func (ctrl *PatientController) Register(c echo.Context) error {

	registerReq := request.Patients{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.patientService.Register(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}
func (ctrl *PatientController) AllPatient(c echo.Context) error {

	result, err := ctrl.patientService.AllPatient()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromPatientListDomain(result))

}

func (ctrl *PatientController) Update(c echo.Context) error {

	updateReq := request.Patients{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	getData, _ := ctrl.patientService.PatientByID(id)
	result, err := ctrl.patientService.Update(id, updateReq.ToDomain())
	result.ID = getData.ID

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdatePatient(result))

}
func (ctrl *PatientController) PatientByID(c echo.Context) error {

	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.patientService.PatientByID(itemID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllPatient(result))
}
func (ctrl *PatientController) PatientByRM(c echo.Context) error {

	noRM, _ := strconv.Atoi(c.Param("no_rm"))

	result, err := ctrl.patientService.PatientByID(noRM)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllPatient(result))
}
func (ctrl *PatientController) Delete(c echo.Context) error {

	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.patientService.Delete(deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)

}
