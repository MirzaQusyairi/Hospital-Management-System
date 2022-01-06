package schedules

import (
	"net/http"

	"Hospital-Management-System/business/schedules"
	"Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/schedules/request"
	"Hospital-Management-System/controllers/schedules/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ScheduleController struct {
	scheduleService schedules.Service
}

func NewControllerSchedule(serv schedules.Service) *ScheduleController {
	return &ScheduleController{
		scheduleService: serv,
	}
}

func (ctrl *ScheduleController) AddSchedule(c echo.Context) error {

	registerReq := request.Schedules{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.scheduleService.AddSchedule(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *ScheduleController) AllSchedule(c echo.Context) error {

	result, err := ctrl.scheduleService.AllSchedule()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromScheduleListDomain(result))

}

func (ctrl *ScheduleController) Update(c echo.Context) error {

	updateReq := request.Schedules{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	getData, _ := ctrl.scheduleService.ScheduleByID(id)
	result, err := ctrl.scheduleService.Update(id, updateReq.ToDomain())
	result.ID = getData.ID

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdateSchedule(result))

}

func (ctrl *ScheduleController) ScheduleByID(c echo.Context) error {

	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.scheduleService.ScheduleByID(itemID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllSchedule(result))
}

func (ctrl *ScheduleController) Delete(c echo.Context) error {

	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.scheduleService.Delete(deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)
}
