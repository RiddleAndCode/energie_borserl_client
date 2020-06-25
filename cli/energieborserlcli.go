package main
import (
   "bufio"
   "fmt"
   ebc "github.com/RiddleAndCode/energieclient"
   "os"
   "strconv"
	"time"
)
var ipAddress string = "192.168.20.77"
var port string = "5555"

func main() {
	fmt.Print("Enter IP Address: ") //Print function is used to display output in same line
	scanner := bufio.NewScanner(os.Stdin)
	url := "http://" + ipAddress + ":" + port + "/transaction"
	fmt.Println("URL:>", url)
	for true {
		fmt.Print("\nEnter Consumption: ") //Print function is used to display output in same line
		var consumption string
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		consumption = scanner.Text()
		ebc.SendConsumptionandGetIPDBTX(consumption, getEpochTime(), ipAddress, port)
	}
}

func getEpochTime() string {
	return strconv.FormatInt((time.Now().Unix()), 10)
}



