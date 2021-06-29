package userin

type FinishInput struct {
	Paid          int64   `json:"paid"`
	PaymentMethod string  `json:"method"`
	TypeOfOrder   string  `json:"type"`
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
}


