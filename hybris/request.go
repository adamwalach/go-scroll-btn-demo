package hybris

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

// Token - oauth2 token
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

var (
	addressID    = "8796125855767"
	paymentID    = "8796093055018"
	apiAddress   = "https://10.8.0.30:9002"
	accountEmail = "adam.walach@gmail.com"
	accountPass  = "test123"
	itemID       = "1934398"
)

func doRequest(rType string, url string, token string) ([]byte, error) {
	log.WithField("url", url).Debug("Request")
	req, err := http.NewRequest(rType, url, nil)
	if err != nil {
		return nil, err
	}

	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	tr := &http.Transport{}
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, errors.New("Unable to read response body: " + err.Error())
	}

	if resp.StatusCode != 200 &&
		resp.StatusCode != 201 {
		return body, errors.New("Received status code: " + strconv.Itoa(resp.StatusCode))
	}
	return body, nil
}
