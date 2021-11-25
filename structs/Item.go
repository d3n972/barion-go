package structs

type IItem interface {
}

type Item struct {
	Name        string
	Description string
	Quantity    int32
	Unit        string
	UnitPrice   float32
	ItemTotal   float32
	SKU         string
}

func (I *Item) NewItem() Item {
	return Item{
		"",
		"",
		0,
		"",
		0.0,
		0.0,
		"",
	}
}
