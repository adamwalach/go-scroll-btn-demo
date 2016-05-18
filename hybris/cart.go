package hybris

import (
	"encoding/json"
	"fmt"
	"sort"

	log "github.com/Sirupsen/logrus"
)

// ByCode - ecommerce carts sorted
type ByCode []Cart

// Carts - ecommerce carts
type Carts struct {
	Carts []Cart `json:"carts"`
}

// Cart - ecommerce cart
type Cart struct {
	Type string `json:"type"`
	Code string `json:"code"`
	GUID string `json:"guid"`
}

func (c ByCode) Len() int           { return len(c) }
func (c ByCode) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByCode) Less(i, j int) bool { return c[i].Code < c[j].Code }

//FindCart list carts and find most recent one:
func FindCart(token string) (string, error) {

	body, err := doRequest("GET", apiAddress+"/rest/v2/electronics/users/"+accountEmail+"/carts", token)
	if err != nil {
		log.WithField("error", err).Error("Carts error")
		return "", err
	}
	var carts Carts
	if err = json.Unmarshal(body, &carts); err != nil {
		log.WithField("error", err).Error("Cart error - unmarshal")
		return "", err
	}

	sort.Sort(sort.Reverse(ByCode(carts.Carts)))
	return carts.Carts[0].Code, nil
}

//AddToCart adds predefined item to cart
func AddToCart() {
	fmt.Println("Add to cart!")
	body, err := doRequest("POST", apiAddress+"/authorizationserver/oauth/token?client_id=raspi&client_secret=secret&grant_type=password&username="+accountEmail+"&password="+accountPass, "")
	if err != nil {
		log.WithField("error", err).Error("Token error")
		return
	}

	fmt.Println(string(body))

	//-------------------------
	var token Token

	if err = json.Unmarshal(body, &token); err != nil {
		log.WithField("error", err).Error("Token error - unmarshal")
		return
	}
	fmt.Println("Token: ", token.AccessToken)

	//create cart:
	body, err = doRequest("POST", apiAddress+"/rest/v2/electronics/users/"+accountEmail+"/carts", token.AccessToken)
	if err != nil {
		log.WithField("error", err).Error("Carts error")
		return
	}
	var cart Cart
	if err = json.Unmarshal(body, &cart); err != nil {
		fmt.Println(string(body))
		log.WithField("error", err).Error("Cart error - unmarshal")
		return
	}

	fmt.Println("Cart code: ", cart.Code)

	//-------------------------
	body, err = doRequest("POST", apiAddress+"/rest/v2/electronics/users/"+accountEmail+"/carts/"+cart.Code+"/entries?code="+itemID+"&qty=1", token.AccessToken)
	if err != nil {
		log.WithField("error", err).Error("Add item error")
		return
	}
	fmt.Println(string(body))
	checkout(token, cart.Code)
}
