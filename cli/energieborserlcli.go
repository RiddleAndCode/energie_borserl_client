package main
import (
   "bufio"
   "fmt"
   ebc "github.com/RiddleAndCode/energieclient"
   "os"
)
func main() {
   fmt.Print("Enter IP Address: ") //Print function is used to display output in same line
   scanner := bufio.NewScanner(os.Stdin)
   scanner.Scan() // use `for scanner.Scan()` to keep reading
   ipAddress := scanner.Text()
   fmt.Print("Enter Port: ")
   scanner.Scan() // use `for scanner.Scan()` to keep reading
   port := scanner.Text()
   url := "http://" + ipAddress + ":" + port + "/transaction"
   fmt.Println("URL:>", url)
   for true {
      fmt.Print("\nEnter Consumption: ") //Print function is used to display output in same line
      var consumption string
      scanner.Scan() // use `for scanner.Scan()` to keep reading
      consumption = scanner.Text()
      ebc.SendConsumptionandGetIPDBTX(consumption,ipAddress,port)
   }
}


