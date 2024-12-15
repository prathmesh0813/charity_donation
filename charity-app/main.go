package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Charity struct {
	AssetType string `json:"assetType"`
	CharityId string `json:"charityID"`
	Amount    string `json:"amount"`
	Cause     string `json:"cause"`
}

type Donar struct {
	AssetType     string `json:"donar"`
	Name          string `json:"name"`
	DonarID       string `json:"donarID"`
	CharityID     string `json:"charityID"`
	Amount        string `json:"amount"`
	TransactionID string `json:"transactionID"`
	Status        string `json:"status"`
}

type CharityData struct {
	AssetType string `json:"AssetType"`
	CharityId string `json:"charityID"`
	Amount    string `json:"amount"`
	Cause     string `json:"cause"`
}

type DonarData struct {
	AssetType     string `json:"assetType"`
	Name          string `json:"name"`
	DonarID       string `json:"donarID"`
	CharityID     string `json:"charityID"`
	Amount        string `json:"amount"`
	TransactionID string `json:"transactionID"`
	Status        string `json:"status"`
}

type DonarHistory struct {
	Record    *DonarHistory `json:"record"`
	TxId      string        `json:"txId"`
	Timestamp string        `json:"timestamp"`
	IsDelete  bool          `json:"isDelete"`
}

func main() {
	router := gin.Default()

	var wg sync.WaitGroup
	wg.Add(1)
	go ChaincodeEventListener("charityorg", "charitychannel", "charity", &wg)

	//router.Static("/public", "./public")
	//router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		result := submitTxnFn("charityorg", "charitychannel", "charity", "CharityContract", "query", make(map[string][]byte), "GetAllCharities")

		var charity []CharityData

		if len(result) > 0 {
			// Unmarshal the JSON array string into the cars slice
			if err := json.Unmarshal([]byte(result), &charity); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.JSON(http.StatusOK, charity)
	})

	// router.GET("/charityorg", func(ctx *gin.Context) {
	// 	result := submitTxnFn("charityorg", "charitychannel", "charity", "CharityContract", "query", make(map[string][]byte), "GetAllCharities")

	// 	var charity []CharityData

	// 	if len(result) > 0 {
	// 		// Unmarshal the JSON array string into the cars slice
	// 		if err := json.Unmarshal([]byte(result), &charity); err != nil {
	// 			fmt.Println("Error:", err)
	// 			return
	// 		}
	// 	}

	// 	ctx.JSON(http.StatusOK, charity)
	// })

	router.POST("/api/charity", func(ctx *gin.Context) {
		var req Charity
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		fmt.Printf("car response %s", req)
		submitTxnFn("charityorg", "charitychannel", "charity", "CharityContract", "invoke", make(map[string][]byte), "CreateCharity", req.CharityId, req.Amount, req.Cause)

		ctx.JSON(http.StatusOK, req)
	})

	router.GET("/api/charity/:id", func(ctx *gin.Context) {
		charityId := ctx.Param("id")

		result := submitTxnFn("charityorg", "charitychannel", "charity", "CharityContract", "query", make(map[string][]byte), "ReadCharity", charityId)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	// router.GET("/api/order/match-car", func(ctx *gin.Context) {
	// 	carID := ctx.Query("carId")
	// 	result := submitTxnFn("charityorg", "charitychannel", "charity", "CharityContract", "query", make(map[string][]byte), "GetMatchingOrders", carID)

	// 	// fmt.Printf("result %s", result)

	// 	var orders []OrderData

	// 	if len(result) > 0 {
	// 		// Unmarshal the JSON array string into the orders slice
	// 		if err := json.Unmarshal([]byte(result), &orders); err != nil {
	// 			fmt.Println("Error:", err)
	// 			return
	// 		}
	// 	}

	// 	ctx.HTML(http.StatusOK, "matchOrder.html", gin.H{
	// 		"title": "Matching Orders", "orderList": orders, "carId": carID,
	// 	})
	// })

	// router.POST("/api/car/match-order", func(ctx *gin.Context) {
	// 	var req Match
	// 	if err := ctx.BindJSON(&req); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	// 		return
	// 	}

	// 	fmt.Printf("match  %s", req)
	// 	submitTxnFn("charityorg", "charitychannel", "charity", "CharityContract", "invoke", make(map[string][]byte), "MatchOrder", req.CarId, req.OrderId)

	// 	ctx.JSON(http.StatusOK, req)
	// })

	router.GET("/api/event", func(ctx *gin.Context) {
		result := getEvents()
		fmt.Println("result:", result)

		ctx.JSON(http.StatusOK, gin.H{"charityEvent": result})

	})

	// router.GET("/donar", func(ctx *gin.Context) {

	// 	ctx.HTML(http.StatusOK, "dealer.html", gin.H{
	// 		"title": "Dealer Dashboard",
	// 	})
	// })

	//Get all orders
	router.GET("/api/donar/all", func(ctx *gin.Context) {

		result := submitTxnFn("donar", "charitychannel", "charity", "DonarContract", "query", make(map[string][]byte), "GetAllDonars")

		var donars []DonarData

		if len(result) > 0 {
			// Unmarshal the JSON array string into the orders slice
			if err := json.Unmarshal([]byte(result), &donars); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.JSON(http.StatusOK, donars)
	})

	// router.POST("/api/donar", func(ctx *gin.Context) {
	// 	var req Donar
	// 	if err := ctx.BindJSON(&req); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	// 		return
	// 	}

	// 	fmt.Printf("donar  %s", req)

	// 	privateData := map[string][]byte{
	// 		"name":          []byte(req.Name),
	// 		"charityID":     []byte(req.CharityID),
	// 		"amount":        []byte(req.Amount),
	// 		"transactionID": []byte(req.TransactionID),
	// 	}
	// 	// privateData2 := map[string][]byte{
	// 	// 	"name":      []byte("Ram"),
	// 	// 	"charityID": []byte("c03"),
	// 	// 	"amount":    []byte("100"),
	// 	// 	"txnsID":    []byte("CCC"),
	// 	// }

	// 	fmt.Println(privateData)

	// 	submitTxnFn("donar", "charitychannel", "charity", "DonarContract", "private", privateData, "CreateDonar", req.DonarID)

	// 	//fmt.Println(res)
	// 	ctx.JSON(http.StatusOK, req)
	// })
	router.POST("/api/donar", func(ctx *gin.Context) {
		var req Donar
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		fmt.Printf("order  %s", req)

		privateData := map[string][]byte{
			"name":          []byte(req.Name),
			"charityID":     []byte(req.CharityID),
			"amount":        []byte(req.Amount),
			"transactionID": []byte(req.TransactionID),
		}
		// privateData2 := map[string][]byte{
		// 	"name":      []byte("Ram"),
		// 	"charityID": []byte("c03"),
		// 	"amount":    []byte("100"),
		// 	"txnsID":    []byte("CCC"),
		// }

		//fmt.Println(privateData)
		submitTxnFn("donar", "charitychannel", "charity", "DonarContract", "private", privateData, "CreateDonar", req.DonarID)

		ctx.JSON(http.StatusOK, req)
	})
	router.GET("/api/donar/:id", func(ctx *gin.Context) {
		donarId := ctx.Param("id")

		result := submitTxnFn("donar", "charitychannel", "charity", "DonarContract", "query", make(map[string][]byte), "ReadDonar", donarId)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.GET("/auditor", func(ctx *gin.Context) {
		result := submitTxnFn("auditor", "charitychannel", "charity", "CharityContract", "query", make(map[string][]byte), "GetAllCharities")

		var charity []CharityData

		if len(result) > 0 {
			// Unmarshal the JSON array string into the cars slice
			if err := json.Unmarshal([]byte(result), &charity); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.JSON(http.StatusOK, charity)
	})

	router.GET("/api/charity/history", func(ctx *gin.Context) {
		donarID := ctx.Query("donarID")
		result := submitTxnFn("auditor", "charitychannel", "charity", "CharityContract", "query", make(map[string][]byte), "GetDonarHistory", donarID)

		// fmt.Printf("result %s", result)

		var donar []DonarHistory

		if len(result) > 0 {
			// Unmarshal the JSON array string into the orders slice
			if err := json.Unmarshal([]byte(result), &donar); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.JSON(http.StatusOK, donar)

	})

	// router.POST("/api/car/register", func(ctx *gin.Context) {
	// 	var req Register
	// 	if err := ctx.BindJSON(&req); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	// 		return
	// 	}

	// 	fmt.Printf("car response %s", req)
	// 	submitTxnFn("mvd", "charitychannel", "charity", "CharityContract", "invoke", make(map[string][]byte), "RegisterCar", req.CarId, req.CarOwner, req.RegNumber)

	// 	ctx.JSON(http.StatusOK, req)
	// })

	router.Run("localhost:8083")
}
