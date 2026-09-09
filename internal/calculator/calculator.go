package calculator

import (
	"khrai-chui-khrai/internal/domain"
	"khrai-chui-khrai/internal/mapper"
	"math"
)

func CalculateService(request domain.CalculateRequest) []domain.CalculateResponse {

	var response []domain.CalculateResponse

	for _, o := range request.OrderItems {

		// calculate
		amountToPay := float64(o.PurchasedAmount) * float64(request.PaidAmount) / float64(request.TotalAmount)
		amountToPay = math.Round(amountToPay*100) / 100

		// response
		response = append(response, mapper.CalculateResponseMapper(o, amountToPay))
	}

	return response
}
