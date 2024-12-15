package main

import "fmt"

func main(){
	// result := submitTxnFn(
	// 	"charityorg",
	// 	"charitychannel",
	// 	"charity",
	// 	"CharityContract",
	// 	"invoke",
	// 	make(map[string][]byte),
	// 	"CreateCharity",
	// 	"C69",
	// 	"100000",
	// 	"Orphan",
	// )

	privateData := map[string][]byte{
		"name":       []byte("Ram"),
		"charityID":      []byte("c03"),
		"amount":      []byte("100"),
		"txnsID": []byte("CCC"),
	}

	result := submitTxnFn("donar", "charitychannel", "charity", "DonarContract", "private", privateData, "CreateDonar", "D01")

	//result := submitTxnFn("donar", "charitychannel", "charity", "DonarContract", "query", make(map[string][]byte), "ReadDonar", "D01")

	//result := submitTxnFn("charityorg", "charitychannel", "charity", "CharityContract", "query", make(map[string][]byte), "GetAllCharities")

	// result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "OrderContract", "query", make(map[string][]byte), "GetAllOrders")

	// result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "GetMatchingOrders", "Car-06")

	// result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "invoke", make(map[string][]byte), "MatchOrder", "Car-06", "ORD-03")

	// result := submitTxnFn("mvd", "autochannel", "KBA-Automobile", "CarContract", "invoke", make(map[string][]byte), "RegisterCar", "Car-06", "Dani", "KL-01-CD-01")




	 	
	//result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "ReadCar", "Car-06")




	fmt.Println(result)

}