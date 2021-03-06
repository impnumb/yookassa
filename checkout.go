package yookassa

//Endpoint is API Endpoint
const Endpoint = "https://api.yookassa.ru/v3/"

//Checkout is checkout struct
type Checkout struct {
	ShopID        int
	SecurityToken string
	OAuthToken    string
}

//New create Checkout
func New(shopID int, securityToken, oauthToken string) *Checkout {
	return &Checkout{
		ShopID:        shopID,
		SecurityToken: securityToken,
		OAuthToken:    oauthToken,
	}
}
