package yookassa

//Payment is
//https://yookassa.ru/developers/api#payment_object
type Payment struct {
	ID            string         `json:"id"`
	Status        string         `json:"status"`
	Amount        *Amount        `json:"amount"`
	IncomeAmount  *Amount        `json:"income_amount,omitempty"`
	Description   string         `json:"description,omitempty"`
	Recipient     *Recipient     `json:"recipient,omitempty"`
	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`
}

//Amount is
//https://yookassa.ru/developers/api#payment_object_amount
type Amount struct {
	Value    float64 `json:"value,string"`
	Currency string  `json:"currency"`
}

//Recipient is
//https://yookassa.ru/developers/api#payment_object_recipient
type Recipient struct {
	AccountID int `json:"account_id,string"`
	GatewayID int `json:"gateway_id,string"`
}

//PaymentMethod is
//https://yookassa.ru/developers/api#payment_object_payment_method
type PaymentMethod struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Saved bool   `json:"saved"`
	Title string `json:"title,omitempty"`
	Login string `json:"login,omitempty"`
	Card  *Card  `json:"card,omitempty"`
}

//Card is
//https://yookassa.ru/developers/api#payment_object_payment_method_bank_card_card
type Card struct {
	First6        uint32 `json:"first6,string,omitempty"`
	Last4         uint16 `json:"last4,string"`
	ExpiryYear    uint16 `json:"expiry_year,string"`
	ExpiryMonth   uint8  `json:"expiry_month,string"`
	CardType      string `json:"card_type"`
	IssuerCountry string `json:"issuer_country,omitempty"`
}
