package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
)

// 生成新钱包
func main() {
	privateKey, _ := crypto.GenerateKey()
	// 转化为字节
	privateKeyByte := crypto.FromECDSA(privateKey)
	// 转化为16进制 私钥
	logrus.Info("ethereum private :", hexutil.Encode(privateKeyByte))
	// 转化为16进制
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	publicKeyByte := crypto.FromECDSAPub(publicKeyECDSA)
	logrus.Info("ethereum public :", hexutil.Encode(publicKeyByte))
	// 转化为公共地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	logrus.Info("ethereum public address :", address)
	/*
		hash := sha3.NewLegacyKeccak256()
		hash.Write(publicKeyBytes[1:])
		fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
	*/
}
