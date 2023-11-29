package comment

type CreatePayload struct {
	Message       string `json:"message" validate:"required"`
	AppointmentId uint   `json:"appointment_id" validate:"required"`
}

type UpdatePayload struct {
	Message string `json:"message" validate:"required"`
}

type Response struct {
	Id      uint    `json:"id"`
	Message string  `json:"message"`
	Creator Creator `json:"creator"`
}

type Creator struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
