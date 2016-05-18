package hybris

import log "github.com/Sirupsen/logrus"

func checkout(token Token, cartCode string) {
	body, err := doRequest("GET", apiAddress+"/rest/v2/electronics/users/"+accountEmail+"/addresses", token.AccessToken)
	if err != nil {
		log.WithField("error", err).Error("Carts error")
		return
	}
	log.Print(string(body))

	body, err = doRequest("PUT", apiAddress+"/rest/v2/electronics/users/"+accountEmail+"/carts/"+cartCode+"/addresses/delivery?addressId="+addressID, token.AccessToken)
	if err != nil {
		log.Print(string(body))
		log.WithField("error", err).Error("Delivery address error")
		return
	}

	_, err = doRequest("PUT", apiAddress+"/rest/v2/electronics/users/"+accountEmail+"/carts/"+cartCode+"/deliverymode?deliveryModeId=premium-gross", token.AccessToken)
	if err != nil {
		log.WithField("error", err).Error("Delivery mode error")
		return
	}

	_, err = doRequest("PUT", apiAddress+"/rest/v2/electronics/users/"+accountEmail+"/carts/"+cartCode+"/paymentdetails?paymentDetailsId="+paymentID, token.AccessToken)
	if err != nil {
		log.WithField("error", err).Error("Payment details error")
		return
	}

	_, err = doRequest("POST", apiAddress+"/rest/v2/electronics/users/"+accountEmail+"/orders?cartId="+cartCode+"&securityCode=123", token.AccessToken)
	if err != nil {
		log.WithField("error", err).Error("Order error")
		log.Print(string(body))
		return
	}
}
