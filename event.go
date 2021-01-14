package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"ethescan/ethtool"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	log.Println("监听交易开始......")
	monitor()
	pushMonitor()
}


// ETH 监听
func pushMonitor() {
	client, err := ethclient.Dial("ws://192.168.3.114:8546")
	if err != nil {
		log.Fatalln(err)
	}
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Println(err)
				continue
			}
			for _, tx := range block.Transactions() {
				if handle.RedisSIsMember("TransactionLog", strings.ToLower(tx.Hash().String())) {
					// 有交易hash.. 之后写自己的逻辑
					fmt.Println("you", tx.Hash().String())

				}
			}
		}
	}
}
