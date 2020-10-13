package energieborserlclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Status string
	Body   string
	Header http.Header
}

func SendConsumptionandGetIPDBTX(consumption string, epoch string, ipAddress string, port string) (*Response, error) {
	url := "http://" + ipAddress + ":" + port + "/transaction"
	var jsonStr = `{"consumption":` + `"` + consumption + `"` + "," + " " + `"epoch":` + `"` + epoch + `"` + "}"
	var jsonByteArray = []byte(jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByteArray))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := Response{Status: resp.Status, Body: string(body), Header: resp.Header}
	return &r, err

}

func CheckServerHealth(ipAddress string, port string) (*Response, error) {
	url := "http://" + ipAddress + ":" + port + "/health"
	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := Response{Status: resp.Status, Body: string(body), Header: resp.Header}
	return &r, err

}
