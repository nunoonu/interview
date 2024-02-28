package domain

import "time"

type ListAppointmentsRequest struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type ListInterviewsResponse struct {
	Appointments []Appointment `json:"appointments"`
	PageInfo     *PageInfo     `json:"page"`
}

type GetAppointmentRequest struct {
	AppointmentID string `uri:"appointmentId" binding:"required"`
}

type GetAppointmentResponse struct {
	Appointment Appointment `json:"appointment"`
	Comments    []Comment   `json:"comments"`
	Histories   []History   `json:"histories"`
}

type UpdateAppointmentRequest struct {
	Card
	UserID string
}

type UpdateAppointmentResponse struct {
}

type KeepAppointmentRequest struct {
	AppointmentID string `uri:"appointmentId" binding:"required"`
}

type KeepAppointmentResponse struct {
}

type CreateAppointmentCommentRequest struct {
	AppointmentID string `uri:"appointmentId" binding:"required"`
	Message       string `json:"message"`
	UserID        string
}

type CreateAppointmentCommentResponse struct {
	CommentID *string `json:"commentId"`
}

type UpdateAppointmentCommentRequest struct {
	AppointmentID string `uri:"appointmentId" binding:"required"`
	CommentID     string `uri:"commentId" binding:"required"`
	Message       string `json:"message"`
	UserID        string
}

type UpdateAppointmentCommentResponse struct {
}

type DeleteAppointmentCommentRequest struct {
	AppointmentID string `uri:"appointmentId" binding:"required"`
	CommentID     string `uri:"commentId" binding:"required"`
	UserID        string
}

type DeleteAppointmentCommentResponse struct {
}

type Appointment struct {
	Card
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
}

type Card struct {
	AppointmentID string `uri:"appointmentId" binding:"required" json:"appointmentId"`
	CardName      string `json:"cardName"`
	Status        string `json:"status"`
	Message       string `json:"message"`
}

type History struct {
	ID            string    `json:"historyId"`
	AppointmentID string    `json:"appointmentId"`
	CardName      string    `json:"cardName"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
}

type Histories []History

type Comment struct {
	ID            string    `json:"commentId"`
	AppointmentID string    `json:"appointmentId"`
	Message       string    `json:"message"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
}

type Comments []Comment

type PageInfo struct {
	Total      int64 `json:"total"`
	IsLastPage bool  `json:"isLastPage"`
}
