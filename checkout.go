package yookassa

//Endpoint is API Endpoint
const Endpoint = "https://api.yookassa.ru/v3/"

//Checkout is checkout struct
type Checkout struct {
	ShopID        int
	SecurityToken string
	OAuthToken    string
}

//NewCheckout create Checkout
func NewCheckout() {

}
