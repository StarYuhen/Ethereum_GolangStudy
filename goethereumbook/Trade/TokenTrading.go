package main

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
	"math/big"
)

// 进行模拟代币交易
func main() {
	// 使用教程的测试网 https://goethereumbook.org/zh/transfer-tokens/
	client, _ := ethclient.Dial("https://rinkeby.infura.io")
	// 因为代币，以太坊交易值改为0 https://tokenfactory.surge.sh
	EthValue := big.NewInt(0)
	// 使用文章地址，防止出现其他问题
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	// 代币合约地址
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")
	// 函数方法切片传递
	transferFnSignature := []byte("transfer(address,uint256)")
	// 生成签名hash
	hash := sha3.New384()
	hash.Write(transferFnSignature)
	// 切片
	methID := hash.Sum(nil)[:4]
	// 填充地址
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	// 发送代币数目
	Number := new(big.Int)
	Number.SetString("1000000000000000000000", 10)
	// 代币量填充
	var data []byte
	data = append(data, methID...)
	data = append(data, paddedAddress...)
	data = append(data, common.LeftPadBytes(Number.Bytes(), 32)...)

	// 使用方法估算燃气费
	gasLime, _ := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	logrus.Info("Ethereum GasLime:", gasLime)
	// 进行交易事务
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    246,
		GasPrice: EthValue,
		Gas:      gasLime,
		To:       &tokenAddress,
		Value:    EthValue,
		Data:     make([]byte, 0),
	})
	// 这里的转账和eth的转账一样，就不多说了
	logrus.Info(tx)
}
