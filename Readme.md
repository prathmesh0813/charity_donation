

### Project Name: CHARITY DONATION AND FUND TRACKING SYSTEM


### Create Project folder `Charity_donation`. In that create the folders `Chaincode` and `Charity-network`.

##### ***Open a command terminal with in Charity-network folder, let's call this terminal as host terminal

```
cd Charitye-network/

```

### Run Network using startCharityNetwork.sh script

##### Build `startCharityNetwork.sh` script file

```
chmod +x startCharityNetwork.sh
```

```
./startCharityNetwork.sh
```




### To stop the network using script file

##### Build `stopCharityNetwork.sh` script file
```
chmod +x stopCharityNetwork.sh
```
```
./stopCharityNetwork.sh
```

### Now create the contracts for chartiy organisation and donar contract

#### Create `charity-org-contract.go` and `donar-contract.go` file and write smart contract in these two files.

# Client


## Create the client folder

```bash
cd ..
```

```bash
mkdir Client
```

### Build the client application

```bash
cd Client

go mod init client

Create and build profile.go, connect.go, client.go, main.go

go get google.golang.org/grpc@v1.67.1

go mod tidy

go run .

```

**#Create a Events folder with in the Charity_donation directory**

```

mkdir Events

```

**#Build the events code**

```
cd Events

go mod init events

Create & Build profile.go, connect.go, events.go, main.go

go mod tidy



```

### Set a charity creation event in the CreateCharity function in charity-org-contract.go 

```
type EventData struct{
	Type string
	Model string
}
```


```
			addCharityEventData := EventData{
				Type:      "Charity creation",
				CharityId: charityID,
			}
			eventDataByte, _ := json.Marshal(addCharityEventData)
			ctx.GetStub().SetEvent("CreateCharity", eventDataByte)
```




### To Run Block Event



Open a terminal in the events folder & consider this terminal as to listening block events.

`go run .`

Note: For checking newly created block do a transaction using client application.(open a terminal from Client folder and execute `go run .`)

### To Run Chaincode Event

open another terminal from Events folder and consider it as to listening chaincode events.

`go run .`


Note: Do a create car transaction using client application.( change CreateCar transaction arguments in main.go then execute `go run .` in the client terminal)

### To Run Private Block Event

Open a new terminal from Events folder and consider it to listening private blockevent.

`go run .`

Note: Submit a CreateDonar transaction using client application.(Edit main.go for CreateDonar transaction and execute `go run .`in the client terminal)


**#To stop the Automobile network**

```
./stopAutomobileNetwork.sh

```
