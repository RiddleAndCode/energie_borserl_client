package energieclient
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)


func SendConsumptionandGetIPDBTX(consumption string, ipAddress string, port string){
	url := "http://" + ipAddress + ":" + port + "/transaction"
	var temp_str = `{"consumption":` +`"`+ consumption + `"`+"}"
	var jsonStr = []byte(temp_str)
	fmt.Println(temp_str)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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
