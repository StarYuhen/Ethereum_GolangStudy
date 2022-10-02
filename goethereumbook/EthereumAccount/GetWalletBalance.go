package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

// 查询钱包地址
func main() {
	// 连接到本地测试节点
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		logrus.Error("client error:", err)
	}
	// (0) 0xE280029a7867BA5C9154434886c241775ea87e53 (100 ETH)
	account := common.HexToAddress("0xE280029a7867BA5C9154434886c241775ea87e53")
	// 获取区块好并返回余额地址
	balance, _ := client.BalanceAt(context.Background(), account, nil)
	logrus.Info("区块地址余额：", balance)
}
