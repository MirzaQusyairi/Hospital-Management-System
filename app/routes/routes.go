package routes

import (
	middlewareApp "Hospital-Management-System/app/middlewares"
	"Hospital-Management-System/business"
	controller "Hospital-Management-System/controllers"
	"Hospital-Management-System/controllers/admins"
	"Hospital-Management-System/controllers/doctors"
	"Hospital-Management-System/controllers/facilities"
	"Hospital-Management-System/controllers/nurses"
	"Hospital-Management-System/controllers/patients"
	"Hospital-Management-System/controllers/prescriptions"
	"Hospital-Management-System/controllers/schedules"
	"Hospital-Management-System/controllers/sesbooks"
	"Hospital-Management-System/controllers/seschedules"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware             middleware.JWTConfig
	PatientController         patients.PatientController
	FaciltyController         facilities.FaciltyController
	DoctorController          doctors.DoctorController
	AdminController           admins.AdminController
	ScheduleController        schedules.ScheduleController
	SessionScheduleController seschedules.SessionScheduleController
	PrescriptionController    prescriptions.PrescriptionController
	SesbookController         sesbooks.SesbookController

	nurses.NurseController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// Admins
	e.POST("/api/v1/admins/login", cl.AdminController.Login)
	e.POST("/api/v1/admins/register", cl.AdminController.Register)

	e.POST("/api/v1/doctors/add/prescription", cl.PrescriptionController.AddPrescription, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationDoctor())
	e.PUT("/api/v1/doctors/update/prescription/:id", cl.PrescriptionController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationDoctor())
	e.DELETE("/api/v1/doctors/delete/prescription/:id", cl.PrescriptionController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationDoctor())
	e.GET("/api/v1/doctors/list/prescription", cl.PrescriptionController.AllPrescription, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationDoctorandNurse())
	e.GET("/api/v1/doctors/prescription/:id", cl.PrescriptionController.PrescriptionByID, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationDoctorandNurse())
	e.GET("/api/v1/doctors/list/prescription/:id", cl.PrescriptionController.PrescriptionByIDSessionBooking, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationDoctorandNurse())

	e.POST("/api/v1/admins/add/schedule", cl.ScheduleController.AddSchedule, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.PUT("/api/v1/admins/update/schedule/:id", cl.ScheduleController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.DELETE("/api/v1/admins/delete/schedule/:id", cl.ScheduleController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.GET("/api/v1/admins/list/schedule", cl.ScheduleController.AllSchedule, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAll())
	e.GET("/api/v1/admins/schedule/:id", cl.ScheduleController.ScheduleByID, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())

	e.POST("/api/v1/admins/add/sessionschedule", cl.SessionScheduleController.AddSessionSchedule, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.PUT("/api/v1/admins/update/sessionschedule/:id", cl.SessionScheduleController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.DELETE("/api/v1/admins/delete/sessionschedule/:id", cl.SessionScheduleController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.GET("/api/v1/admins/list/sessionschedule", cl.SessionScheduleController.AllSessionSchedule, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAll())
	e.GET("/api/v1/admins/sessionschedule/:id", cl.SessionScheduleController.SessionScheduleByID, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())

	e.POST("/api/v1/admins/add/facilty", cl.FaciltyController.AddFacilty, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.PUT("/api/v1/admins/update/facilty/:id", cl.FaciltyController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.DELETE("/api/v1/admins/delete/facilty/:id", cl.FaciltyController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.GET("/api/v1/admins/list/facilty", cl.FaciltyController.AllFacilty, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAll())
	e.GET("/api/v1/admins/facilty/:id", cl.FaciltyController.FacByID, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())

	e.POST("/api/v1/admins/add/patient", cl.PatientController.Register, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.PUT("/api/v1/admins/update/patient/:id", cl.PatientController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.DELETE("/api/v1/admins/delete/patient/:id", cl.PatientController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.GET("/api/v1/admins/patient/", cl.PatientController.PatientByRM, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.GET("/api/v1/admins/list/patient", cl.PatientController.AllPatient, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAll())
	e.GET("/api/v1/admins/patient/:id", cl.PatientController.PatientByID, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAll())

	e.POST("/api/v1/admins/add/doctor", cl.DoctorController.Register, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.PUT("/api/v1/admins/update/doctor/:id", cl.DoctorController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.DELETE("/api/v1/admins/delete/doctor/:id", cl.DoctorController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.GET("/api/v1/admins/list/doctor", cl.DoctorController.AllDoctor, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAll())
	e.GET("/api/v1/admins/doctor/:id", cl.DoctorController.DoctorByID, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdminandDoctor())

	e.POST("/api/v1/admins/add/nurse", cl.NurseController.Register, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.PUT("/api/v1/admins/update/nurse/:id", cl.NurseController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.DELETE("/api/v1/admins/delete/nurse/:id", cl.NurseController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.GET("/api/v1/admins/list/nurse", cl.NurseController.AllNurse, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.GET("/api/v1/admins/nurse/:id", cl.NurseController.NurseByID, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdminandNurse())

	e.POST("/api/v1/admins/add/sessionbook", cl.SesbookController.AddSessionBook, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.PUT("/api/v1/admins/update/sessionbook/:id", cl.SesbookController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.DELETE("/api/v1/admins/delete/sessionbook/:id", cl.SesbookController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	e.GET("/api/v1/admins/list/sessionbook", cl.SesbookController.AllSessionBook, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAll())
	e.GET("/api/v1/admins/sessionbook/:id", cl.SesbookController.SessionBookByID, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	// Doctors
	e.POST("/api/v1/doctors/login", cl.DoctorController.Login)

	// Nurse
	e.POST("/api/v1/nurses/login", cl.NurseController.Login)
}

func RoleValidationAll() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "admin" || claims.Role == "doctor" || claims.Role == "nurse" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}

func RoleValidationAdmin() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "admin" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}

func RoleValidationDoctor() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "doctor" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}

func RoleValidationNurse() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "nurse" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}

func RoleValidationDoctorandNurse() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "doctor" || claims.Role == "nurse" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}

func RoleValidationAdminandDoctor() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "admin" || claims.Role == "doctor" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}

func RoleValidationAdminandNurse() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "admin" || claims.Role == "nurse" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}
