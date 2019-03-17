package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

var client = ubOzoneClient

func getBlock(blockNo int64) *types.Block {
	blockNumber := big.NewInt(blockNo)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	return block
}

func getLastBlockNo() uint64 {
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return block.Number().Uint64()
}

func checkBlock() {
	//调用客户端的HeadByNumber来返回有关一个区块的头信息。若传入nil，它将返回最新的区块头。
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current Block is: ", header.Number.String())

	//调用客户端的BlockByNumber方法来获得完整区块。
	// 您可以读取该区块的所有内容和元数据，例如，区块号，区块时间戳，区块摘要，区块难度以及交易列表等等。
	blockNumber := big.NewInt(51)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}



	fmt.Println("Block Number: ", block.Number().Uint64())
	fmt.Println("Timestamp:", block.Time().Uint64())
	fmt.Println("Difficulty: ", block.Difficulty().Uint64())
	fmt.Println("Block Hash: ", block.Hash().Hex())
	fmt.Println("How many tx: ", len(block.Transactions()))

	//调用Transaction只返回一个区块的交易数目
	txCount, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("How many tx: ", txCount)

	// get the data back by txhash
	checkTxbyTxHash(block.Transaction(common.HexToHash("0x7c029aea694ed0ac0ef6239b88b6dcb0382a308e1b8191e0375243f6492509f9")))
}

func checkAllBlock() {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current Block is: ", header.Number.String())
	fmt.Println()

	//调用客户端的BlockByNumber方法来获得完整区块。
	// 您可以读取该区块的所有内容和元数据，例如，区块号，区块时间戳，区块摘要，区块难度以及交易列表等等。
	lastBlockNumber := getLastBlockNo()
	var i int64
	for i = 0 ; i <= int64(lastBlockNumber); i++{
		block, err := client.BlockByNumber(context.Background(), big.NewInt(i))
		if err != nil {
			log.Fatal(err)
		}

		if len(block.Transactions()) != 0 {
			fmt.Println("Block Number: ", block.Number().Uint64())
			fmt.Println("Timestamp:", block.Time().Uint64())
			fmt.Println("Difficulty: ", block.Difficulty().Uint64())
			fmt.Println("Block Hash: ", block.Hash().Hex())
			fmt.Println("How many tx: ", len(block.Transactions()))
			fmt.Println()
		}
	}
}
