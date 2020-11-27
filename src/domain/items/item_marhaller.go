package items

import (
	"encoding/json"
)

type PublicItem struct {
	Id                string      `json:"id"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	Status            string      `json:"status"`
}

func (i *Item) MarshalJSON() ([]byte, error) {
	return json.Marshal(PublicItem{
		Id:                i.Id,
		Title:             i.Title,
		Description:       i.Description,
		Price:             i.Price,
		AvailableQuantity: i.AvailableQuantity,
		Status:            i.Status,
	})
}
