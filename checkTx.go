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

	fmt.Println("Stationscode: ", ozone.Stationscode)
	fmt.Println("Stationsname: ", ozone.Stationsname)
	fmt.Println("Tagesmaxima: ", ozone.Tagesmaxima)
	fmt.Println("Aktuellster Messtag im Jahr:", ozone.Aktuellster_Messtag_im_Jahr)
	fmt.Println("Erster Messtag im Jahr: ", ozone.Erster_Messtag_im_Jahr)
	fmt.Print("Jan:", ozone.Jan, ", ")
	fmt.Print("Feb:", ozone.Feb, ", ")
	fmt.Print("Mar:", ozone.Mar, ", ")
	fmt.Print("Apr:", ozone.Apr, ", ")
	fmt.Print("Mai:", ozone.Mai, ", ")
	fmt.Println("Jun:", ozone.Jun, ", ")
	fmt.Print("Jul:", ozone.Jul, ", ")
	fmt.Print("Aug:", ozone.Aug, ", ")
	fmt.Print("Sep:", ozone.Sep, ", ")
	fmt.Print("Okt:", ozone.Okt, ", ")
	fmt.Print("Nov:", ozone.Nov, ", ")
	fmt.Print("Dez:", ozone.Dez)
}
