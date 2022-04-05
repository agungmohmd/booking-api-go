package request

type BookingRequest struct {
	ID            int    `json:"id"`
	BookingPerson string `json:"booking_person"`
	BookAt        string `json:"book_at"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
