package main

import (
	"bufio"
	ebc "energieborserlclient"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

var ipAddress string
var port string
var scanner *bufio.Scanner

func main() {

	clearDisplay()
	parseCMDLine()
	displayURL(ipAddress, port)
	menu()
	scanner = bufio.NewScanner(os.Stdin)

	for true {
		scanner.Scan()
		selection := scanner.Text()
		switch strings.ToLower(selection) {
		case "t":
			transaction()
			clearDisplay()
			menu()
		case "h":
			health()
			clearDisplay()
			menu()
		case "q":
			os.Exit(0)
		default:
			clearDisplay()
			fmt.Println("Invalid Choice.")
			menu()

		}

	}
}

func getEpochTime() string {
	return strconv.FormatInt((time.Now().Unix()), 10)
}

func transaction() {
	var consumption string
	var resp *ebc.Response
	fmt.Print("\nEnter Consumption: ") //Print function is used to display output in same line
	scanner.Scan()                     // use `for scanner.Scan()` to keep reading
	consumption = scanner.Text()
	resp, err := ebc.SendConsumptionandGetIPDBTX(consumption, getEpochTime(), ipAddress, port)
	if err == nil {
		fmt.Print("\n", resp.Status, "\n\n", resp.Body, "\n")
	} else {
		fmt.Print("\nServer Unreachable " + err.Error())
	}
	fmt.Print("\nPress Enter to Continue ...")
	scanner.Scan() // use `for scanner.Scan()` to keep reading

}

func menu() {

	menuTable := tablewriter.NewWriter(os.Stdout)
	menuTable.SetHeader([]string{"Command", "Description"})
	data := [][]string{
		[]string{"\"t\"", "Transaction Building"},
		[]string{"\"h\"", "Server Health Check"},
		[]string{"\"q\"", "Terminate the Program"},
	}
	for _, v := range data {
		menuTable.Append(v)
	}
	menuTable.Render() // Send output
	fmt.Print("\n", "Selection : ")

}

func health() {
	var resp *ebc.Response
	var serverStatusMessage string
	var serverStatusCode string

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Status", "Header", "Time"})

	resp, err := ebc.CheckServerHealth(ipAddress, port)
	if err == nil {
		serverStatusMessage = resp.Body
		serverStatusCode = resp.Status
	} else {
		serverStatusMessage = err.Error()
		serverStatusCode = "-"
	}

	data := [][]string{
		[]string{serverStatusCode, serverStatusMessage, time.Now().Format(time.RFC850)},
	}
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
	fmt.Print("\nPress Any Button to Continue ...\n")
	scanner.Scan() // use `for scanner.Scan()` to keep reading

}

func clearDisplay() {

	fmt.Println("\033[2J")

}

func displayURL(ip string, port string) {

	url := "http://" + ipAddress + ":" + port + "/"

	data := [][]string{
		[]string{ip, port, url},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"IP", "PORT", "URL"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
	fmt.Println("")
}

func parseCMDLine() {

	flag.StringVar(&ipAddress, "ip", "192.168.20.77", "IP of the signer device")
	flag.StringVar(&port, "port", "5555", "TCP port of the signer device")
	flag.Parse()

	helptable := tablewriter.NewWriter(os.Stdout)
	helptable.SetHeader([]string{"USAGE"})
	data := [][]string{
		[]string{"./energie_borserl_cli -ip=192.168.xx.xx -port=5555"}}
	for _, v := range data {
		helptable.Append(v)
	}
	helptable.Render() // Send output
	fmt.Println("")

}
