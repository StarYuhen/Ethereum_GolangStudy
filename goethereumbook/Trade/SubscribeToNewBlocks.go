package main

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

// 订阅新区块
func main() {
	// 这个地址原本免费，现在需要注册获取token，请自己更改
	client, _ := ethclient.Dial("wss://eth-mainnet.ws.alchemyapi.io/ws/")
	// 创建通道
	heads := make(chan *types.Header)
	// 订阅返回的对象
	sub, _ := client.SubscribeNewHead(context.Background(), heads)
	// 循环监听
	for {
		select {
		case err := <-sub.Err():
			// 监听区块失败的信息
			logrus.Error("SubScribe Block Error:", err)
		case Heads := <-heads:
			logrus.Info("Heads Address:", Heads.Hash().Hex())
			// 监听区块头信息
			block, _ := client.BlockByHash(context.Background(), Heads.Hash())
			logrus.Info("Block Address:", block.Hash().String())
		}
	}
}
