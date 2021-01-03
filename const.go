package yookassa

//Payment statuses
//https://yookassa.ru/developers/payments/payment-process#payment-statuses
const (
	Pending           = "pending"
	WaitingForCapture = "waiting_for_capture"
	Succeeded         = "succeeded"
	Canceled          = "canceled"
)

//Currency codes
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

//Payment methods
//https://yookassa.ru/developers/payment-methods/overview#all
const (
	BankCard      = "bank_card"
	ApplePay      = "apple_pay"
	GooglePay     = "google_pay"
	YooMoney      = "yoo_money"
	Qiwi          = "qiwi"
	Webmoney      = "webmoney"
	Sberbank      = "sberbank"
	Alfabank      = "alfabank"
	TinkoffBank   = "tinkoff_bank"
	B2bSberbank   = "b2b_sberbank"
	MobileBalance = "mobile_balance"
	Cash          = "cash"
	Installments  = "installments"
	WeChat        = "wechat"
)

//Card types
//https://yookassa.ru/developers/api#payment_object_payment_method_bank_card_card_card_type
const (
	MasterCard      = "MasterCard"
	Visa            = "Visa"
	Mir             = "Mir"
	UnionPay        = "UnionPay"
	JCB             = "JCB"
	AmericanExpress = "AmericanExpress"
	DinersClub      = "DinersClub"
	Unknown         = "Unknown"
)

//Сountry codes
//https://www.iso.org/obp/ui/#iso:pub:PUB500001:en
const (
	AD = "AD"
	AE = "AE"
	AF = "AF"
	AG = "AG"
	AI = "AI"
	AL = "AL"
	AM = "AM"
	AO = "AO"
	AQ = "AQ"
	AR = "AR"
	AS = "AS"
	AT = "AT"
	AU = "AU"
	AW = "AW"
	AX = "AX"
	AZ = "AZ"
	BA = "BA"
	BB = "BB"
	BD = "BD"
	BE = "BE"
	BF = "BF"
	BG = "BG"
	BH = "BH"
	BI = "BI"
	BJ = "BJ"
	BL = "BL"
	BM = "BM"
	BN = "BN"
	BO = "BO"
	BQ = "BQ"
	BR = "BR"
	BS = "BS"
	BT = "BT"
	BV = "BV"
	BW = "BW"
	BY = "BY"
	BZ = "BZ"
	CA = "CA"
	CC = "CC"
	CD = "CD"
	CF = "CF"
	CG = "CG"
	CH = "CH"
	CI = "CI"
	CK = "CK"
	CL = "CL"
	CM = "CM"
	CN = "CN"
	CO = "CO"
	CR = "CR"
	CU = "CU"
	CV = "CV"
	CW = "CW"
	CX = "CX"
	CY = "CY"
	CZ = "CZ"
	DE = "DE"
	DJ = "DJ"
	DK = "DK"
	DM = "DM"
	DO = "DO"
	DZ = "DZ"
	EC = "EC"
	EE = "EE"
	EG = "EG"
	EH = "EH"
	ER = "ER"
	ES = "ES"
	ET = "ET"
	FI = "FI"
	FJ = "FJ"
	FK = "FK"
	FM = "FM"
	FO = "FO"
	FR = "FR"
	GA = "GA"
	GB = "GB"
	GD = "GD"
	GE = "GE"
	GF = "GF"
	GG = "GG"
	GH = "GH"
	GI = "GI"
	GL = "GL"
	GM = "GM"
	GN = "GN"
	GP = "GP"
	GQ = "GQ"
	GR = "GR"
	GS = "GS"
	GT = "GT"
	GU = "GU"
	GW = "GW"
	GY = "GY"
	HK = "HK"
	HM = "HM"
	HN = "HN"
	HR = "HR"
	HT = "HT"
	HU = "HU"
	ID = "ID"
	IE = "IE"
	IL = "IL"
	IM = "IM"
	IN = "IN"
	IO = "IO"
	IQ = "IQ"
	IR = "IR"
	IS = "IS"
	IT = "IT"
	JE = "JE"
	JM = "JM"
	JO = "JO"
	JP = "JP"
	KE = "KE"
	KG = "KG"
	KH = "KH"
	KI = "KI"
	KM = "KM"
	KN = "KN"
	KP = "KP"
	KR = "KR"
	KW = "KW"
	KY = "KY"
	KZ = "KZ"
	LA = "LA"
	LB = "LB"
	LC = "LC"
	LI = "LI"
	LK = "LK"
	LR = "LR"
	LS = "LS"
	LT = "LT"
	LU = "LU"
	LV = "LV"
	LY = "LY"
	MA = "MA"
	MC = "MC"
	MD = "MD"
	ME = "ME"
	MF = "MF"
	MG = "MG"
	MH = "MH"
	MK = "MK"
	ML = "ML"
	MM = "MM"
	MN = "MN"
	MO = "MO"
	MP = "MP"
	MQ = "MQ"
	MR = "MR"
	MS = "MS"
	MT = "MT"
	MU = "MU"
	MV = "MV"
	MW = "MW"
	MX = "MX"
	MY = "MY"
	MZ = "MZ"
	NA = "NA"
	NC = "NC"
	NE = "NE"
	NF = "NF"
	NG = "NG"
	NI = "NI"
	NL = "NL"
	NO = "NO"
	NP = "NP"
	NR = "NR"
	NU = "NU"
	NZ = "NZ"
	OM = "OM"
	PA = "PA"
	PE = "PE"
	PF = "PF"
	PG = "PG"
	PH = "PH"
	PK = "PK"
	PL = "PL"
	PM = "PM"
	PN = "PN"
	PR = "PR"
	PS = "PS"
	PT = "PT"
	PW = "PW"
	PY = "PY"
	QA = "QA"
	RE = "RE"
	RO = "RO"
	RS = "RS"
	RU = "RU"
	RW = "RW"
	SA = "SA"
	SB = "SB"
	SC = "SC"
	SD = "SD"
	SE = "SE"
	SG = "SG"
	SH = "SH"
	SI = "SI"
	SJ = "SJ"
	SK = "SK"
	SL = "SL"
	SM = "SM"
	SN = "SN"
	SO = "SO"
	SR = "SR"
	SS = "SS"
	ST = "ST"
	SV = "SV"
	SX = "SX"
	SY = "SY"
	SZ = "SZ"
	TC = "TC"
	TD = "TD"
	TF = "TF"
	TG = "TG"
	TH = "TH"
	TJ = "TJ"
	TK = "TK"
	TL = "TL"
	TM = "TM"
	TN = "TN"
	TO = "TO"
	TR = "TR"
	TT = "TT"
	TV = "TV"
	TW = "TW"
	TZ = "TZ"
	UA = "UA"
	UG = "UG"
	UM = "UM"
	US = "US"
	UY = "UY"
	UZ = "UZ"
	VA = "VA"
	VC = "VC"
	VE = "VE"
	VG = "VG"
	VI = "VI"
	VN = "VN"
	VU = "VU"
	WF = "WF"
	WS = "WS"
	YE = "YE"
	YT = "YT"
	ZA = "ZA"
	ZM = "ZM"
	ZW = "ZW"
)

//VAT types
//https://yookassa.ru/developers/api?lang=bash#payment_object_payment_method_b2b_sberbank_vat_data_calculated_type
const (
	Untaxed    = "untaxed"
	Calculated = "calculated"
	Mixed      = "mixed"
)

//Rate types
//https://yookassa.ru/developers/api?lang=bash#payment_object_payment_method_b2b_sberbank_vat_data_calculated_rate
const (
	Rate7  = "7"
	Rate10 = "10"
	Rate18 = "18"
	Rate20 = "20"
)
