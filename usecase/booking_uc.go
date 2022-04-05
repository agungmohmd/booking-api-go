package usecase

import (
	"time"

	"github.com/agungmohmd/booking-api/repository"
	"github.com/agungmohmd/booking-api/repository/models"
	request "github.com/agungmohmd/booking-api/server/requests"
	"github.com/agungmohmd/booking-api/usecase/viewmodel"
)

type BookingUC struct {
	*ContractUC
}

func (uc BookingUC) BuildBody(data *models.Booking, res *viewmodel.BookingVM) {
	res.ID = data.ID
	res.BookingPerson = data.BookingPerson
	res.BookAt = data.BookAt
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt
}

func (uc BookingUC) SelectAll() (res []viewmodel.BookingVM, err error) {
	repo := repository.NewBookingRepository(uc.DB)
	data, err := repo.SelectAll()
	if err != nil {
		return res, err
	}
	for _, r := range data {
		temp := viewmodel.BookingVM{}
		uc.BuildBody(&r, &temp)
		res = append(res, temp)
	}
	return res, err
}

func (uc BookingUC) FindOne(id string) (res viewmodel.BookingVM, err error) {
	repo := repository.NewBookingRepository(uc.DB)
	data, err := repo.FindOne(id)
	if err != nil {
		return res, err
	}
	uc.BuildBody(&data, &res)
	return res, err
}

func (uc BookingUC) Add(req *request.BookingRequest) (res models.Booking, err error) {
	repo := repository.NewBookingRepository(uc.DB)
	now := time.Now().UTC()
	res = models.Booking{
		ID:            req.ID,
		BookingPerson: req.BookingPerson,
		BookAt:        req.BookAt,
		CreatedAt:     now.Format(time.RFC3339),
	}
	res.ID, err = repo.Add(res)
	if err != nil {
		return res, err
	}
	return res, err
}
