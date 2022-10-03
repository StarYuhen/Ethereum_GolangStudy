package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

// 查询交易信息
func main() {
	// 先获取区块的信息事务
	client, _ := ethclient.Dial("https://cloudflare-eth.com")
	// * 自动获取块查询区块事务信息
	// HeadBlock, _ := client.BlockByNumber(context.Background(), nil)
	// NewBlockNumber, _ := client.BlockByNumber(context.Background(), HeadBlock.Number())
	// for _, tx := range NewBlockNumber.Transactions() {
	// 	// // 使用tx获取发送方地址 EIP155S
	// 	// ChanID, _ := client.NetworkID(context.Background())
	// 	// // AsMessage 获取信息 新版AsMessage 缺少一个类型Signer https://github.com/ethereum/go-ethereum/issues/23890
	// 	// if message, err := tx.AsMessage(types.LatestSignerForChainID(ChanID), ChanID); err != nil {
	// 	// 	// logrus.Error(message.From().Hash())
	// 	// } else {
	// 	// // 地址当前信息 0x009BD41feaA4FF711a41b05ff69fc65f14d16f6D
	// 	// 	logrus.Info(message.From().String())
	// 	// }
	//
	// 	// 获取事务是否成功
	// 	receipt, _ := client.TransactionReceipt(context.Background(), tx.Hash())
	// 	logrus.Info("Transaction Receipt Bool:", receipt.Status)
	// }

	// // 不自动获取块，指定获取块事务信息
	// blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	// count, _ := client.TransactionCount(context.Background(), blockHash)
	// // 遍历获取所有事务
	// for index := uint(0); index < count; index++ {
	// 	tx, _ := client.TransactionInBlock(context.Background(), blockHash, index)
	// 	logrus.Info("Block address:", tx.Hash().String())
	// }

	// 使用指定哈希获取事务
	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, is, _ := client.TransactionByHash(context.Background(), txHash)
	logrus.Info(tx.Hash().String(), ",is ", is)
}
