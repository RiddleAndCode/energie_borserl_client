# Wien Energie - Energiebörserl client software

The *energieclient* package provides an easy to use interface to send consumption data to the logging device with ``SendConsumptionandGetIPDBTX`` function which will sign the payload with ECDSA, build the BigChainDb transaction and finally send the transaction Json as a string back to the client.

Logger device status can be querried via a simple ``CheckServerHealth`` function that also takes port number and IP as arguments and responds with the status of the server.

``SendConsumptionandGetIPDBTX("145.232 kWh","192.168.20.77","5555")``

``CheckServerHealth("192.168.20.77","5555")``


### Installation

Make sure your go environment is set and type the following on your terminal

``go get -u github.com/RiddleAndCode/energieclient``

### Usage

There are 3 options to send data and get the IPDB transaction JSON in return.

#### Option 1: CLI


There is CLI that demonstrates the usage of the *energieclient* package. Ip address and the port of the logging device is a global variable inside go source code. User can tip in any consumption data which will be forwarded to the logging device. Epoch time is taken automatically from the system time and added to the request.

##### Installation

``go get -u github.com/olekukonko/tablewriter``

##### Usage

To run the example simply navigate to cli folder

`cd $GOPATH/src/github.com/RiddleAndCode/energieclient/cli`

and run with the following arguments

``go run energieborserlcli.go -ip=192.168.xx.xx -port=xxx``

Default values for the IP and Port are as follows :

```
ip = 192.168.20.77
port = 5555

```

#### Option 2: Send a command from terminal

A JSON of the following type is expected by the server.

`{
    "consumption": "2452.2323 KWH",
    "epoch": "1593121481"
}`

The following command can be sent to directly communicate with the Logger device.

`  curl -X POST -H "Content-Type: application/json" -d @data.json http://192.168.20.77:5555/transaction `

Dont forget to put the JSON to a file named `data.json`.

#### Option 3: Using `SendConsumptionandGetIPDBTX` function from your own implementation

The ``SendConsumptionandGetIPDBTX(consumption string, epoch ,ip string, port string)`` function from ``energieclient`` package
can be called from another implementation. A simple example:

``response := energieclient.SendConsumptionandGetIPDBTX("145.232 kWh","1593121481","192.168.20.77","5555")``


### Example Response
#### Transaction Building
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
#### Health Checks
If everything is as expected on thhe server side you will get a message like below :
```
Signer Device is Healthy.

```
In case of an error, PKCS error from the lower layers is propogated back to the client.
```
pkcs11: 0x6: CKR_FUNCTION_FAILED

```
**Response Structure**

There are two fields in the transaction that hold the public keys:

 1. **base58 SECP256K1 pubkey**: 
 
 Inside the Secure Element, there is a SECP256K1 key which servers as the device identity. All the data such as comsumption and epoch is signed with this key. The  public counter part of this key is shown under this field in the transaction. It is base58 encoded and it also adheres to ASN.1 Structure with Der Encoding. This means, to have the actual bytes of the public key:
  * one has to first base58 decode this field : [online decoder](https://www.dcode.fr/base-58-cipher)
  * and then use a ASN.1 parser [ASN.1 parser](https://lapo.it/asn1js/)


  2. **public_key** :
  
  This public key is related to the IPDB network and represents the account name on the ledger. It is base58 encoded as well but is in **raw byte format** and **doesnt** require ASN.1 decoding.
