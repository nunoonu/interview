package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nunoonu/interview/internal/core/domain"
	"github.com/nunoonu/interview/internal/core/ports"
	"net/http"
)

type AppointmentHandler struct {
	auc ports.AppointmentUsecase
}

func NewAppointmentHandler(auc ports.AppointmentUsecase) *AppointmentHandler {
	return &AppointmentHandler{auc: auc}
}

func (h *AppointmentHandler) ListAppointments(ctx *gin.Context) {

	var req domain.ListAppointmentsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.auc.ListAppointments(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *AppointmentHandler) GetAppointment(ctx *gin.Context) {

	var req domain.GetAppointmentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.auc.GetAppointment(ctx, &req)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *AppointmentHandler) UpdateAppointment(ctx *gin.Context) {

	var req domain.UpdateAppointmentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if v, ok := ctx.Get("rct"); ok {
		rct := v.(*RouteContext)
		req.UserID = rct.UserID
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get route context"})
		return
	}
	res, err := h.auc.UpdateAppointment(ctx, &req)
	if err != nil {
		if errors.Is(err, domain.ErrorConflict) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, res)
}

func (h *AppointmentHandler) CreateComment(ctx *gin.Context) {
	var req domain.CreateAppointmentCommentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if v, ok := ctx.Get("rct"); ok {
		rct := v.(*RouteContext)
		req.UserID = rct.UserID
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get route context"})
		return
	}

	res, err := h.auc.CreateAppointmentComment(ctx, &req)
	if err != nil {
		if errors.Is(err, domain.ErrorConflict) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *AppointmentHandler) UpdateComment(ctx *gin.Context) {
	var req domain.UpdateAppointmentCommentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if v, ok := ctx.Get("rct"); ok {
		rct := v.(*RouteContext)
		req.UserID = rct.UserID
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get route context"})
		return
	}

	res, err := h.auc.UpdateAppointmentComment(ctx, &req)
	if err != nil {
		if errors.Is(err, domain.ErrorConflict) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, res)
}

func (h *AppointmentHandler) DeleteComment(ctx *gin.Context) {
	var req domain.DeleteAppointmentCommentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if v, ok := ctx.Get("rct"); ok {
		rct := v.(*RouteContext)
		req.UserID = rct.UserID
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get route context"})
		return
	}

	res, err := h.auc.DeleteAppointmentComment(ctx, &req)
	if err != nil {
		if errors.Is(err, domain.ErrorConflict) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, res)
}

func (h *AppointmentHandler) KeepAppointment(ctx *gin.Context) {
	var req domain.KeepAppointmentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.auc.KeepAppointment(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, res)
}
