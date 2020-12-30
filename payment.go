package yookassa

//Payment statuses
//https://yookassa.ru/developers/payments/payment-process#payment-statuses
const (
	Pending           = "pending"
	WaitingForCapture = "waiting_for_capture"
	Succeeded         = "succeeded"
	Canceled          = "canceled"
)

//https://www.iso.org/iso-4217-currency-codes.html
const (
	AFN = "AFN"
	EUR = "EUR"
	ALL = "ALL"
	DZD = "DZD"
	USD = "USD"
	AOA = "AOA"
	XCD = "XCD"
	ARS = "ARS"
	AMD = "AMD"
	AWG = "AWG"
	AUD = "AUD"
	AZN = "AZN"
	BSD = "BSD"
	BHD = "BHD"
	BDT = "BDT"
	BBD = "BBD"
	BYN = "BYN"
	BZD = "BZD"
	XOF = "XOF"
	BMD = "BMD"
	INR = "INR"
	BTN = "BTN"
	BOB = "BOB"
	BOV = "BOV"
	BAM = "BAM"
	BWP = "BWP"
	NOK = "NOK"
	BRL = "BRL"
	BND = "BND"
	BGN = "BGN"
	BIF = "BIF"
	CVE = "CVE"
	KHR = "KHR"
	XAF = "XAF"
	CAD = "CAD"
	KYD = "KYD"
	CLP = "CLP"
	CLF = "CLF"
	CNY = "CNY"
	COP = "COP"
	COU = "COU"
	KMF = "KMF"
	CDF = "CDF"
	NZD = "NZD"
	CRC = "CRC"
	HRK = "HRK"
	CUP = "CUP"
	CUC = "CUC"
	ANG = "ANG"
	CZK = "CZK"
	DKK = "DKK"
	DJF = "DJF"
	DOP = "DOP"
	EGP = "EGP"
	SVC = "SVC"
	ERN = "ERN"
	SZL = "SZL"
	ETB = "ETB"
	FKP = "FKP"
	FJD = "FJD"
	XPF = "XPF"
	GMD = "GMD"
	GEL = "GEL"
	GHS = "GHS"
	GIP = "GIP"
	GTQ = "GTQ"
	GBP = "GBP"
	GNF = "GNF"
	GYD = "GYD"
	HTG = "HTG"
	HNL = "HNL"
	HKD = "HKD"
	HUF = "HUF"
	ISK = "ISK"
	IDR = "IDR"
	XDR = "XDR"
	IRR = "IRR"
	IQD = "IQD"
	ILS = "ILS"
	JMD = "JMD"
	JPY = "JPY"
	JOD = "JOD"
	KZT = "KZT"
	KES = "KES"
	KPW = "KPW"
	KRW = "KRW"
	KWD = "KWD"
	KGS = "KGS"
	LAK = "LAK"
	LBP = "LBP"
	LSL = "LSL"
	ZAR = "ZAR"
	LRD = "LRD"
	LYD = "LYD"
	CHF = "CHF"
	MOP = "MOP"
	MKD = "MKD"
	MGA = "MGA"
	MWK = "MWK"
	MYR = "MYR"
	MVR = "MVR"
	MRU = "MRU"
	MUR = "MUR"
	XUA = "XUA"
	MXN = "MXN"
	MXV = "MXV"
	MDL = "MDL"
	MNT = "MNT"
	MAD = "MAD"
	MZN = "MZN"
	MMK = "MMK"
	NAD = "NAD"
	NPR = "NPR"
	NIO = "NIO"
	NGN = "NGN"
	OMR = "OMR"
	PKR = "PKR"
	PAB = "PAB"
	PGK = "PGK"
	PYG = "PYG"
	PEN = "PEN"
	PHP = "PHP"
	PLN = "PLN"
	QAR = "QAR"
	RON = "RON"
	RUB = "RUB"
	RWF = "RWF"
	SHP = "SHP"
	WST = "WST"
	STN = "STN"
	SAR = "SAR"
	RSD = "RSD"
	SCR = "SCR"
	SLL = "SLL"
	SGD = "SGD"
	XSU = "XSU"
	SBD = "SBD"
	SOS = "SOS"
	SSP = "SSP"
	LKR = "LKR"
	SDG = "SDG"
	SRD = "SRD"
	SEK = "SEK"
	CHE = "CHE"
	CHW = "CHW"
	SYP = "SYP"
	TWD = "TWD"
	TJS = "TJS"
	TZS = "TZS"
	THB = "THB"
	TOP = "TOP"
	TTD = "TTD"
	TND = "TND"
	TRY = "TRY"
	TMT = "TMT"
	UGX = "UGX"
	UAH = "UAH"
	AED = "AED"
	USN = "USN"
	UYU = "UYU"
	UYI = "UYI"
	UYW = "UYW"
	UZS = "UZS"
	VUV = "VUV"
	VES = "VES"
	VND = "VND"
	YER = "YER"
	ZMW = "ZMW"
	ZWL = "ZWL"
	XBA = "XBA"
	XBB = "XBB"
	XBC = "XBC"
	XBD = "XBD"
	XTS = "XTS"
	XXX = "XXX"
	XAU = "XAU"
	XPD = "XPD"
	XPT = "XPT"
	XAG = "XAG"
)

//Payment is
//https://yookassa.ru/developers/api#payment_object
type Payment struct {
	ID            string        `json:"id"`
	Status        string        `json:"status"`
	Amount        Amount        `json:"amount"`
	IncomeAmount  Amount        `json:"income_amount,omitempty"`
	Description   string        `json:"description,omitempty"`
	Recipient     Recipient     `json:"recipient,omitempty"`
	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`
}

//Amount is
//https://yookassa.ru/developers/api#payment_object_amount
type Amount struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

//Recipient is
//https://yookassa.ru/developers/api#payment_object_recipient
type Recipient struct {
	AccountID int `json:"account_id"`
	GatewayID int `json:"gateway_id"`
}

//PaymentMethod is
//https://yookassa.ru/developers/api#payment_object_payment_method
type PaymentMethod struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Saved bool   `json:"saved"`
	Title string `json:"title,omitempty"`
	Login string `json:"login,omitempty"`
}
