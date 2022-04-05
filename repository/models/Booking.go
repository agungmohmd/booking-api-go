package models

type Booking struct {
	ID            int    `json:"id"`
	BookingPerson string `json:"booking_person"`
	BookAt        string `json:"book_at"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

var (
	BookingSelectStatement = `select id, booking_person, book_at, created_at, coalesce(to_char(updated_at,'YYYYMMDD HH24:MI'), '') updated_at from bookings w `
	BookingOrderBy         = ` w.id `
	BookingWhereID         = ` w.id = `
	BookingWhereDate       = ` w.book_at::date = `
)
