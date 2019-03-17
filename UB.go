package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

type Ozone struct {
	Stationscode                string `json:"stationscode"`
	Stationsname                string `json:"stationsname"`
	Tagesmaxima                 string `json:"tagesmaxima"`
	Erster_Messtag_im_Jahr      string `json:"erster_messtag_im_jahr"`
	Aktuellster_Messtag_im_Jahr string `json:"aktuellster_messtag_im_jahr"`
	Jan                         string `json:"jan"`
	Feb                         string `json:"feb"`
	Mar                         string `json:"mar"`
	Apr                         string `json:"apr"`
	Mai                         string `json:"mai"`
	Jun                         string `json:"jun"`
	Jul                         string `json:"jul"`
	Aug                         string `json:"aug"`
	Sep                         string `json:"sep"`
	Okt                         string `json:"okt"`
	Nov                         string `json:"nov"`
	Dez                         string `json:"dez"`
}

type PM25 struct {
	Stationscode                string `json:"stationscode"`
	Stationsname                string `json:"stationsname"`
	Tagesmittelwerte            string `json:"tagesmittelwerte"`
	Messmethode                 string `json:"messmethode"`
	Erster_Messtag_im_Jahr      string `json:"erster_messtag_im_jahr"`
	Aktuellster_Messtag_im_Jahr string `json:"aktuellster_messtag_im_jahr"`
	Jan                         string `json:"jan"`
	Feb                         string `json:"feb"`
	Mar                         string `json:"mar"`
	Apr                         string `json:"apr"`
	Mai                         string `json:"mai"`
	Jun                         string `json:"jun"`
	Jul                         string `json:"jul"`
	Aug                         string `json:"aug"`
	Sep                         string `json:"sep"`
	Okt                         string `json:"okt"`
	Nov                         string `json:"nov"`
	Dez                         string `json:"dez"`
}

func sendOzoneTx() {
	fromKeystore, err := ioutil.ReadFile("/home/shiun/Ethereum/Pri_Data/UB/Ozone/keystore/UTC--2019-03-10T22-32-50.757383048Z--29e6746b6639d422556b24697d1a2276b0642ba4")
	if err != nil {
		log.Fatal(err)
	}
	fromKey, err := keystore.DecryptKey(fromKeystore, "emsdata")
	privateKey := fromKey.PrivateKey
	publicKey := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	nonce, err := ubOzoneClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//设置我们将要转移的ETH数量。 但是我们必须将ETH转换为wei
	amount := big.NewInt(1000000000)

	//ETH转账的燃气应设上限为“21000”单位。
	gasLimit := uint64(210000)

	//燃气价格必须以wei为单位设定
	gasPrice := big.NewInt(30000000000) //30 wei

	/*
	//对燃气价格进行硬编码有时并不理想。
	// go-ethereum客户端提供SuggestGasPrice函数，用于根据'x'个先前块来获得平均燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	*/

	//send to the address
	toAddress := common.HexToAddress("0x5f36247e4f1e5160d6980c4828bafb57ae450d2d")

	// 认证信息组装
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = amount      // in wei
	auth.GasLimit = gasLimit // in units
	auth.GasPrice = gasPrice
	auth.From = fromAddress

	//data := readCSV()
	csvFile, _ := os.Open("/home/shiun/Documents/Masterarbeit/Data/UB_Concentrations_AirPollutants_Germany/Ozone/Ozone_Final.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		ozone := Ozone{}
		ozone.Stationscode = line[0]
		ozone.Stationsname = line[1]
		ozone.Tagesmaxima = line[2]
		ozone.Erster_Messtag_im_Jahr = line[3]
		ozone.Aktuellster_Messtag_im_Jahr = line[4]
		ozone.Jan = line[5]
		ozone.Feb = line[6]
		ozone.Mar = line[7]
		ozone.Apr = line[8]
		ozone.Mai = line[9]
		ozone.Jun = line[10]
		ozone.Jul = line[11]
		ozone.Aug = line[12]
		ozone.Sep = line[13]
		ozone.Okt = line[14]
		ozone.Nov = line[15]
		ozone.Dez = line[16]

		endata, err := json.Marshal(&ozone)
		if err != nil {
			log.Fatal(err)
		}

		tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, endata)
		nonce++

		signedTx, err := auth.Signer(types.HomesteadSigner{}, auth.From, tx)
		if err != nil {
			log.Fatal(err)
		}

		//调用“SendTransaction”来将已签名的事务广播到整个网络
		err = ubOzoneClient.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
		fmt.Println()
	}
	// 等待挖矿完成
	//bind.WaitMined(context.Background(),client,signedTx)
}
