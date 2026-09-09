package domain

type OrderItems struct {
	Name            string `json:"name"`             // ชื่อคน
	ItemName        string `json:"item_name"`        // ออเดอร์
	PurchasedAmount int64  `json:"purchased_amount"` // ยอดแต่ละรายการที่ซื้อ
}

type CalculateRequest struct {
	TotalAmount       int64        `json:"total_amount"`       // ยอดรวม
	GovernmentSupport int64        `json:"government_support"` // รัฐช่วย
	PaidAmount        int64        `json:"paid_amount"`        // จ่ายจริง
	OrderItems        []OrderItems `json:"order_items"`        // รายการแต่ละคน
}

type CalculateResponse struct {
	Name           string `json:"name"`
	ItemName       string `json:"item_name"`
	PurchasedTotal int64  `json:"purchased_total"`
	AmountToPay    string `json:"amount_to_pay"`
}
