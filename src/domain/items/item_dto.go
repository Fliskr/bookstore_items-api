package items

import "fmt"

type Item struct {
	Id                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}

type Items []Item

func (i *Item) String() string {
	return fmt.Sprintf(`
Id: %s 
Description: %s
Seller %d
Title %s
Pictures %d
Video %s
Price %f
AvailableQuantity %d
SoldQuantity %d
Status %s`, i.Id, &i.Description, i.Seller, i.Title, len(i.Pictures), i.Video, i.Price, i.AvailableQuantity, i.SoldQuantity, i.Status)
}

func (d *Description) String() string {
	return fmt.Sprintf("\n  PlainText: \"%s\"\n  HTML: \"%s\"", d.PlainText, d.Html)
}
