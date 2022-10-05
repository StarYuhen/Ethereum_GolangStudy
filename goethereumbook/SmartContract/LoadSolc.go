package main

import (
	store "Ethereum_GolangStudy/goethereumbook/SmartContract/test"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

// 加载智能合约
func main() {
	client, _ := ethclient.Dial("http://127.0.0.1:8545")
	// 加载地址
	address := common.HexToAddress("0x4ad02236A1229029b013838FE0A5e44d349828c9")
	instance, _ := store.NewStore(address, client)
	logrus.Info(instance.StoreCaller)
	// 查询智能合约
	logrus.Info(instance.Version(nil))
}
