package main

import (
	store "Ethereum_GolangStudy/goethereumbook/SmartContract/test"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
)

// 部署一个简易的智能合约
func main() {
	client, _ := ethclient.Dial("http://127.0.0.1:8545")
	// 设定私钥和燃气费
	private, _ := crypto.HexToECDSA("bd380efe66a9ff36d66d4ab20b76e79de62ecc745b3f31362d421bee33dc0446")
	// publicKey := private.Public()
	// public := publicKey.(*ecdsa.PublicKey)
	// NewKeyedTransactor在新版变成了 NewKeyedTransactorWithChainID
	ChanID, _ := client.NetworkID(context.Background())
	auto, _ := bind.NewKeyedTransactorWithChainID(private, ChanID)
	auto.Nonce = big.NewInt(300)
	auto.Value = big.NewInt(0)
	auto.GasLimit = 300
	auto.GasPrice, _ = client.SuggestGasPrice(context.Background())
	// 加载智能合约
	address, tx, instance, _ := store.DeployStore(auto, client, "1.0")
	logrus.Info("solc address :", address)
	logrus.Info("expen:", tx, "instance:", instance)
}
