package models

type ProductShelve struct {
	ID        int  `json:"id"`
	ProductID int  `json:"product_id"`
	ShelveID  int  `json:"shelve_id"`
	IsPrimary bool `json:"is_primary"`
}
