# Wien Energie - Energiebörserl client software

The *energieclient* package provides an easy to use interface to send consumption data to the logging device which in turn will sign it and build the BigChainDb transaction and finally send the transaction Json as a string back to the client. 

There is only one function to be used which takes the port, ip and consumption data in simple string format as input.

``SendConsumptionandGetIPDBTX("145.232 kWh",192.168.0.1,5555)``

### Installation

Make sure your go environment is set and type the following on your terminal

``go get -u github.com/RiddleAndCode/energieclient`` 

### Usage Example

There is CLI that demonstrates the usage of the *energieclient* package. It prompts for the ip address and the port of the logging device. There after, the user can tip in any consumption data which will be forwarded to the logging device.

To run the example simply navigate to cli folder 

`cd $GOPATH/src/github.com/RiddleAndCode/energieclient/cli`

and run

``go run energieborserlcli.go``

