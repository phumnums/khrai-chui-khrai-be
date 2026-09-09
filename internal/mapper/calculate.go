package mapper

import (
	"fmt"
	"khrai-chui-khrai/internal/domain"
)

func CalculateResponseMapper(order domain.OrderItems, amountToPay float64) domain.CalculateResponse {
	result := domain.CalculateResponse{
		Name:           order.Name,
		ItemName:       order.ItemName,
		PurchasedTotal: order.PurchasedAmount,
		AmountToPay:    fmt.Sprintf("%.2f", amountToPay),
	}
	return result
}
