# Wien Energie - Energiebörserl client software

The *energieclient* package provides an easy to use interface to send consumption data to the logging device which in turn will sign it and build the BigChainDb transaction and finally send the transaction Json as a string back to the client. 

There is only one function to be used which takes the port, ip and consumption data in simple string format as input.

``SendConsumptionandGetIPDBTX("145.232 kWh","192.168.0.1","5555")``

### Installation

Make sure your go environment is set and type the following on your terminal

``go get -u github.com/RiddleAndCode/energieclient`` 

### Usage 

There are 3 options to send data and get the IPDB transaction JSON in return.

#### Option 1: CLI


There is CLI that demonstrates the usage of the *energieclient* package. It prompts for the ip address and the port of the logging device. There after, the user can tip in any consumption data which will be forwarded to the logging device.

To run the example simply navigate to cli folder 

`cd $GOPATH/src/github.com/RiddleAndCode/energieclient/cli`

and run

``go run energieborserlcli.go``

#### Option 2: Send a command from terminal

A JSON of the following type is expected by the server.

`{
    "consumption": "2452.2323 KWH"
}`

The following command can be sent to directly communicate with the Logger device.

`  curl -X POST -H "Content-Type: application/json" -d @data.json http://<IP OF THE SIGNER DEVICE>:5555/transaction `

Dont forget to put the JSON to a file named `data.json`. 

#### Option 3: Using `SendConsumptionandGetIPDBTX` function from your own implementation

The ``SendConsumptionandGetIPDBTX(consumption string ,ip string, port string)`` function from ``energieclient`` package
can be called from another implementation. A simple example:

``response := energieclient.SendConsumptionandGetIPDBTX("145.232 kWh","192.168.0.1","5555")`` 


### Example Response

```
{
   "asset":{
      "data":{
         "base58 SECP256K1 pubkey":"PZ8Tyr4Nx8MHsRAGMpZmZ6TWY63dXWSD1526Lch5MWAsTpyoYGQd2exiC7kPMp8c5hcVtdsiwBcEdhfW3AUeWpqu7g5pXkqsGdPRE5j72nKrPQJD1MqaL7n5",
         "base58 signature":"AN1rKs4Mw88qDq8r2Nszy23J3HnDtc1kaDJ3hX4c3qKT5UphVtcY5AHtitJk9NeDgD12obZ7jW8VezKwodn6iXoN5wwe6p8Hi",
         "consumption":"2452.2323 KWH",
         "epoch":"1592571926"
      }
   },
   "id":"005e1a3a07b5ac3f60831d110d583a04c2388183996bfde0b34728d086c8b443",
   "inputs":[
      {
         "fulfillment":"pGSAIPESsaLVnTLX7lwN9Z_jmo4qkp210RmJFj1kb14MhSiegUBU43yKw60hq71ncQVWg0OV62C5nkIecSB-jA6mRLdyf_2oGZMExCTZ5XzTNOhdQs54GkDrfJjhFoIQiJlR6zsL",
         "fulfills":null,
         "owners_before":[
            "HE3niH5XY3tDk6BFyEdAx6zGmsd4Wcpji17W8XRyX2B7"
         ]
      }
   ],
   "metadata":{
      "Wien Energie":"PV Panel"
   },
   "operation":"CREATE",
   "outputs":[
      {
         "amount":"1",
         "condition":{
            "details":{
               "public_key":"HE3niH5XY3tDk6BFyEdAx6zGmsd4Wcpji17W8XRyX2B7",
               "type":"ed25519-sha-256"
            },
            "uri":"ni:///sha-256;8RKxotWdMtfuXA31n-OajiqSnbXRGYkWPWRvXgyFKJ4?cost=131072&fpt=ed25519-sha-256"
         },
         "public_keys":[
            "HE3niH5XY3tDk6BFyEdAx6zGmsd4Wcpji17W8XRyX2B7"
         ]
      }
   ],
   "version":"2.0"
}
```
