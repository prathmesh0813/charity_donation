package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	AssetType string `json:"assetType"`
	CharityId string `json:"charityID"`
	Amount    string `json:"amount"`
	Cause     string `json:"cause"`
}

type DonarData struct {
	AssetType     string `json:"donar"`
	Name          string `json:"name"`
	DonarID       string `json:"donarID"`
	CharityID     string `json:"charityID"`
	Amount        string `json:"amount"`
	TransactionID string `json:"transactionID"`
	Status        string `json:"status"`
}

type Match struct {
	OrderId string `json:"orderId"`
	CarId   string `json:"carId"`
}

type DonarHistory struct {
	Record    *DonarData `json:"record"`
	TxId      string     `json:"txId"`
	Timestamp string     `json:"timestamp"`
	IsDelete  bool       `json:"isDelete"`
}

type Register struct {
	CarId     string `json:"carId"`
	CarOwner  string `json:"carOwner"`
	RegNumber string `json:"regNumber"`
}

func main() {
	router := gin.Default()

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go ChaincodeEventListener("charityorg", "charitychannel", "charity", &wg)

	// router.Static("/public", "./public")
	// router.LoadHTMLGlob("templates/*")

	// router.GET("/", func(ctx *gin.Context) {
	// 	result := submitTxnFn("charityorg", "charitychannel", "charity", "CharityContract", "query", make(map[string][]byte), "GetAllCharities")

	// 	var charity []CharityData

	// 	if len(result) > 0 {
	// 		// Unmarshal the JSON array string into the cars slice
	// 		if err := json.Unmarshal([]byte(result), &charity); err != nil {
	// 			fmt.Println("Error:", err)
	// 			return
	// 		}
	// 	}

	// 	ctx.HTML(http.StatusOK, charity)
	// })

	router.GET("/charityorg", func(ctx *gin.Context) {
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

	router.POST("/api/charity", func(ctx *gin.Context) {
		var req Charity
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		fmt.Printf("charity response %s", req)
		submitTxnFn("charityorg", "charitychannel", "charity", "CharityContract", "invoke", make(map[string][]byte), "CreateCharity", req.CharityId, req.Amount, req.Cause)

		ctx.JSON(http.StatusOK, req)
	})

	router.GET("/api/charity/:id", func(ctx *gin.Context) {
		charityID := ctx.Param("id")

		result := submitTxnFn("charityorg", "charitychannel", "charity", "CharityContract", "query", make(map[string][]byte), "ReadCharity", charityID)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

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

	router.POST("/api/donar", func(ctx *gin.Context) {
		var req Donar
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		fmt.Printf("donar  %s", req)

		privateData := map[string][]byte{
			"donarID":       []byte(req.DonarID),
			"name":          []byte(req.Name),
			"charityId":     []byte(req.CharityID),
			"amount":        []byte(req.Amount),
			"transactionID": []byte(req.TransactionID),
			"status":        []byte(req.Status),
		}

		submitTxnFn("donar", "charitychannel", "charity", "DonarContract", "private", privateData, "CreateDonar", req.DonarID)

		ctx.JSON(http.StatusOK, req)
	})

	router.GET("/api/donar/:id", func(ctx *gin.Context) {
		donarId := ctx.Param("id")

		result := submitTxnFn("donar", "charitychannel", "charity", "DonarContract", "query", make(map[string][]byte), "ReadDonar", donarId)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	// 	router.GET("/mvd", func(ctx *gin.Context) {
	// 		result := submitTxnFn("mvd", "charitychannel", "charity", "CarContract", "query", make(map[string][]byte), "GetAllCars")

	// 		var cars []CarData

	// 		if len(result) > 0 {
	// 			// Unmarshal the JSON array string into the cars slice
	// 			if err := json.Unmarshal([]byte(result), &cars); err != nil {
	// 				fmt.Println("Error:", err)
	// 				return
	// 			}
	// 		}

	// 		ctx.HTML(http.StatusOK, "mvd.html", gin.H{
	// 			"title": "MVD Dashboard", "carList": cars,
	// 		})
	// 	})

	router.GET("/api/donar/history", func(ctx *gin.Context) {
		donarID := ctx.Query("donarID")
		result := submitTxnFn("charityorg", "charitychannel", "charity", "CharityContract", "query", make(map[string][]byte), "GetDonarHistory", donarID)

		// fmt.Printf("result %s", result)

		var donars []DonarHistory

		if len(result) > 0 {
			// Unmarshal the JSON array string into the orders slice
			if err := json.Unmarshal([]byte(result), &donars); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.JSON(http.StatusOK, donars)
	})

	router.Run("localhost:8081")
}
