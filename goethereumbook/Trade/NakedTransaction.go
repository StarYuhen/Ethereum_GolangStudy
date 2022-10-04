package main

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/sirupsen/logrus"
)

// 发送原始交易
func main() {
	// 这个原始url也需要最新的了
	client, _ := ethclient.Dial("https://rinkeby.infura.io")
	// 构建原始交易（裸交易）和之前的eth货币交易差不多，就不管了。
	rawTx := "f86d8202b28477359400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ca05924bde7ef10aa88db9c66dd4f5fb16b46dff2319b9968be983118b57bb50562a001b24b31010004f13d9a26b320845257a6cfc2bf819a3d55e3fc86263c5f0772"
	RawTX, _ := hex.DecodeString(rawTx)
	// 构建原始事务
	tx := new(types.Transaction)
	rlp.DecodeBytes(RawTX, &tx)
	// 广播说一哈
	if err := client.SendTransaction(context.Background(), tx); err != nil {
		logrus.Error("SendTransaction Error:", err)
		return
	}

	logrus.Info(tx.Hash().Hex())
}
