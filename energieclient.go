package energieclient
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)


func SendConsumptionandGetIPDBTX(consumption string, epoch string, ipAddress string, port string) {
	url := "http://" + ipAddress + ":" + port + "/transaction"
	var jsonStr = `{"consumption":` + `"` + consumption + `"` + "," + " " + `"epoch":` + `"` + epoch + `"` + "}"
	var jsonByteArray = []byte(jsonStr)
	fmt.Println(jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByteArray))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

