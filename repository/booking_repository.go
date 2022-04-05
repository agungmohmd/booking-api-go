package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/agungmohmd/booking-api/repository/models"
)

type IBooking interface {
	SelectAll() (data []models.Booking, err error)
	FindOne(id string) (data models.Booking, err error)
	Add(body models.Booking) (res int, err error)
}

type BookingRepository struct {
	DB *sql.DB
}

func NewBookingRepository(db *sql.DB) IBooking {
	return &BookingRepository{DB: db}
}

func (model BookingRepository) SelectAll() (data []models.Booking, err error) {
	query := models.BookingSelectStatement
	rows, err := model.DB.Query(query)

	if err != nil {
		return data, err
	}
	defer rows.Close()
	for rows.Next() {
		d := models.Booking{}
		err = rows.Scan(&d.ID, &d.BookingPerson, &d.BookAt, &d.CreatedAt, &d.UpdatedAt)
		if err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, err
}

func (model BookingRepository) FindOne(id string) (data models.Booking, err error) {
	query := models.BookingSelectStatement + `where ` + models.BookingWhereID + `'` + id + `'`
	err = model.DB.QueryRow(query).Scan(&data.ID, &data.BookingPerson, &data.BookAt, &data.CreatedAt, &data.UpdatedAt)
	return data, err
}

func (model BookingRepository) Add(body models.Booking) (res int, err error) {
	x := []models.Booking{}
	querySelect := models.BookingSelectStatement + `where ` + models.BookingWhereDate + `'` + body.BookAt + `'::date`
	fmt.Println(querySelect)
	rows, err := model.DB.Query(querySelect)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		d := models.Booking{}
		err = rows.Scan(&d.ID, &d.BookingPerson, &d.BookAt, &d.CreatedAt, &d.UpdatedAt)
		x = append(x, d)
	}
	fmt.Println(x)
	if len(x) != 0 {
		//return true'
		res = 000
		err = errors.New("The date are already booked")
		return res, err
	}

	// return false
	queryInsert := `insert into bookings (booking_person, book_at, created_at) values($1, $2, $3) returning id`
	err = model.DB.QueryRow(queryInsert, body.BookingPerson, body.BookAt, body.CreatedAt).Scan(&res)
	return res, err
}
