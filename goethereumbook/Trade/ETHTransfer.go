package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
)

// 实现Eth 转账
func main() {
	// 此次测试使用本地
	client, _ := ethclient.Dial("http://127.0.0.1:8545")
	// 0x7c93ba33296e75fa18cb0ac13bab727273f0f213cd331ac0b16a7dd86e4a390d100Eth
	// 加载私钥-需要去除0x
	privateKey, _ := crypto.HexToECDSA("7c93ba33296e75fa18cb0ac13bab727273f0f213cd331ac0b16a7dd86e4a390d")
	// 获取公共地址
	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	PublicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	logrus.Info("PublicAddress:", PublicAddress.String())
	// 获取账户交易随机数
	Nonce, _ := client.PendingNonceAt(context.Background(), PublicAddress)
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	// gasPrice := big.NewInt(30000000000)      // in wei (30 gwei)
	// 通过市场获取燃气价格
	gasPricer, _ := client.SuggestGasPrice(context.Background())
	// 接收地址 0xE280029a7867BA5C9154434886c241775ea87e53
	ToAddress := common.HexToAddress("0xE280029a7867BA5C9154434886c241775ea87e53")
	// 智能合约通信
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    Nonce,
		GasPrice: gasPricer,
		Gas:      gasLimit,
		To:       &ToAddress,
		Value:    value,
		Data:     make([]byte, 0),
	})
	// 利用私钥进行签名
	ChanID, _ := client.NetworkID(context.Background())
	SignedTx, _ := types.SignTx(tx, types.LatestSignerForChainID(ChanID), privateKey)
	// 广播节点
	if err := client.SendTransaction(context.Background(), SignedTx); err != nil {
		logrus.Error("SendTransaction error:", err)
		return
	}
	logrus.Info("Transaction Success Address:", SignedTx.Hash().Hex())

	/*  交易成功
	time="2022-10-03T16:01:22+08:00" level=info msg="PublicAddress:0x002cFA2FD8423Cf2728961EF014AD2e9126118BE"
	time="2022-10-03T16:01:22+08:00" level=info msg="Transaction Success Address:0xb13eee5b299e81afb90a8a9252d264e23dd947667f88ffb5011f1b1f3c9aa6ec"
	ganache-cli 提示：
	  Gas usage: 21000
	  Block Number: 1
	  Block Time: Mon Oct 03 2022 16:01:22 GMT+0800 (中国标准时间)


	*/
}
