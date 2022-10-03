package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

// 查询区块
func main() {
	// 这里用测试网，本地节点没有区块
	client, _ := ethclient.Dial("https://cloudflare-eth.com")
	// 获取区块头信息
	HeadBlock, _ := client.BlockByNumber(context.Background(), nil)
	// 获取区块查询的头信息 2022-10-03 14:32 是 15665850
	logrus.Info("Query BlockHead:", HeadBlock.Number().String())
	// 获取完整区块信息 ,这里用自动查询的api
	NewBlockNumber, _ := client.BlockByNumber(context.Background(), HeadBlock.Number())
	// 只打印交易地址信息
	logrus.Info("Block Address:", NewBlockNumber.Hash().Hex())
	// 获取交易的数目
	QueryCount, _ := client.TransactionCount(context.Background(), NewBlockNumber.Hash())
	logrus.Info("Block TransactionCount:", QueryCount)
}
