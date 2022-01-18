package seschedules

import (
	"net/http"

	"Hospital-Management-System/business/seschedules"
	"Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/seschedules/request"
	"Hospital-Management-System/controllers/seschedules/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SessionScheduleController struct {
	sessionscheduleService seschedules.Service
}

func NewControllerSessionSchedule(serv seschedules.Service) *SessionScheduleController {
	return &SessionScheduleController{
		sessionscheduleService: serv,
	}
}

func (ctrl *SessionScheduleController) AddSessionSchedule(c echo.Context) error {

	registerReq := request.Seschedules{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.sessionscheduleService.AddSessionSchedule(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *SessionScheduleController) AllSessionSchedule(c echo.Context) error {

	result, err := ctrl.sessionscheduleService.AllSessionSchedule()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromSessionScheduleListDomain(result))

}

func (ctrl *SessionScheduleController) Update(c echo.Context) error {

	updateReq := request.Seschedules{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	getData, _ := ctrl.sessionscheduleService.SessionScheduleByID(id)
	result, err := ctrl.sessionscheduleService.Update(id, updateReq.ToDomain())
	result.ID = getData.ID

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdateSessionSchedule(result))

}

func (ctrl *SessionScheduleController) SessionScheduleByID(c echo.Context) error {

	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.sessionscheduleService.SessionScheduleByID(itemID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllSessionSchedule(result))
}

func (ctrl *SessionScheduleController) SessionScheduleByIDSchedule(c echo.Context) error {

	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.sessionscheduleService.SessionScheduleByIDSchedule(itemID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromSessionScheduleListDomain(result))
}

func (ctrl *SessionScheduleController) Delete(c echo.Context) error {

	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.sessionscheduleService.Delete(deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)
}
