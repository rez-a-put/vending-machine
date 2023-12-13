package model

type ReqBuyData struct {
	Input []float64 `json:"nominals"`
}

type ReqData struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
