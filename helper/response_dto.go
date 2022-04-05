package helper

import "github.com/agungmohmd/booking-api/usecase/viewmodel"

func ErrorResponse(message interface{}) viewmodel.ResponseErrorVm {
	err := []interface{}{message}
	res := viewmodel.ResponseErrorVm{Messages: err}

	return res
}

func SuccessResponse(data interface{}) viewmodel.ResponseSuccessVm {
	return viewmodel.ResponseSuccessVm{
		Data: data,
	}
}
