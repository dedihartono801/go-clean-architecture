package transaction

type CheckoutDto struct {
	Items []ItemDto `json:"items" validate:"required"`
}

type ItemDto struct {
	ID       string `json:"id" `
	Quantity int    `json:"quantity" `
}
