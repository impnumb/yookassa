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
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

//Recipient is
//https://yookassa.ru/developers/api#payment_object_recipient
type Recipient struct {
	AccountID string `json:"account_id"`
	GatewayID string `json:"gateway_id"`
}

//PaymentMethod is
//https://yookassa.ru/developers/api#payment_object_payment_method
type PaymentMethod struct {
	Type             string            `json:"type"`
	ID               string            `json:"id"`
	Saved            bool              `json:"saved"`
	Title            string            `json:"title,omitempty"`
	Login            string            `json:"login,omitempty"`
	Card             *Card             `json:"card,omitempty"`
	PayerBankDetails *PayerBankDetails `json:"payer_bank_details"`
	PaymentPurpose   string            `json:"payment_purpose"`
	VATData          *VATData          `json:"vat_data"`
	Phone            string            `json:"phone,omitempty"`
	AccountNumber    string            `json:"account_number,omitempty"`
}

//Card is
//https://yookassa.ru/developers/api#payment_object_payment_method_bank_card_card
type Card struct {
	First6        string `json:"first6,omitempty"`
	Last4         string `json:"last4"`
	ExpiryYear    string `json:"expiry_year"`
	ExpiryMonth   string `json:"expiry_month"`
	CardType      string `json:"card_type"`
	IssuerCountry string `json:"issuer_country,omitempty"`
	IssuerName    string `json:"issuer_name,omitempty"`
	Source        string `json:"source,omitempty"`
}

//PayerBankDetails is
//https://yookassa.ru/developers/api#payment_object_payment_method_b2b_sberbank_payer_bank_details
type PayerBankDetails struct {
	FullName   string `json:"full_name"`
	ShortName  string `json:"short_name"`
	Address    string `json:"address"`
	INN        string `json:"inn"`
	KPP        string `json:"kpp"`
	BankName   string `json:"bank_name"`
	BankBranch string `json:"bank_branch"`
	BankBIK    string `json:"bank_bik"`
	Account    string `json:"account"`
}

//VATData is
//https://yookassa.ru/developers/api?lang=bash#payment_object_payment_method_b2b_sberbank_vat_data
type VATData struct {
	Type   string  `json:"type"`
	Amount *Amount `json:"amount"`
	Rate   string  `json:"rate"`
}
