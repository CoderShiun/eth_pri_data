package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

func checkTxbyTxHash(tx *types.Transaction) {
	// get the data back by txhash
	ozone := Ozone{}
	fmt.Println("tx data = ", tx.Data())

	err := json.Unmarshal(tx.Data(), &ozone)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx data = ", ozone)
}
