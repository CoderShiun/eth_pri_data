package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func connUBOzoneClient() *ethclient.Client {
	client, err := ethclient.Dial("/home/shiun/Ethereum/Pri_Data/UB/Ozone/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}

	_ = client
	return client
}