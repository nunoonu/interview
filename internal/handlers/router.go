package handlers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(appHdl *AppointmentHandler) *gin.Engine {
	app := gin.Default()
	app.Use(gin.Recovery())
	app.Use(VerifyJWT())
	router := app.Group("/v1")
	router.GET("/health-check", func(c *gin.Context) {
		c.Status(200)
	})
	router.GET("/appointments", appHdl.ListAppointments)
	router.GET("/appointments/:appointmentId", appHdl.GetAppointment)
	router.PUT("/appointments/:appointmentId", appHdl.UpdateAppointment)
	router.PUT("/appointments/:appointmentId/keep", appHdl.KeepAppointment)
	router.POST("/appointments/:appointmentId/comments", appHdl.CreateComment)
	router.PUT("/appointments/:appointmentId/comments/:commentId", appHdl.UpdateComment)
	router.DELETE("/appointments/:appointmentId/comments/:commentId", appHdl.DeleteComment)

	return app
}
