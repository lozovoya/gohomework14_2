package dto

type CardDTO struct {
	Id      int64  `json:"id"`
	Number  string `json:"number"`
	Balance int64  `json:"balance"`
	Issuer  string `json:"issuer"`
	Status  string `json:"status"`
}

type TransactionDTO struct {
	Id          int64 `json:"id"`
	Amount      int64 `json:"amount"`
	Category    int64 `json:"category"`
	Description int64 `json:"description"`
	Logo        int64 `json:"logo"`
}

type MessageDTO struct {
	Message string `json:"message"`
}

type MonMostDTO struct {
	CatId int64 `json:"cat_id"`
	Count int64 `json:"count"`
}
