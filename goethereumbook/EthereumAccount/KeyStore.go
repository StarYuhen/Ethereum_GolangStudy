package main

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

// 生成本地密钥文件和读取密钥文件

func main() {
	// // 生成本地密钥文件
	// ks := keystore.NewKeyStore("./goethereumbook/EthereumAccount/KeyStore", keystore.StandardScryptN, keystore.StandardScryptP)
	// password := "StarYuhen"
	// account, _ := ks.NewAccount(password)
	// logrus.Info("address:", account.Address.Hex())

	// 读取本地密钥文件
	file := "./goethereumbook/EthereumAccount/KeyStore/UTC--2022-10-02T06-59-17.491272600Z--0b80d82d8ebe0587ead2da9c18eef4262b2f3200"
	ks := keystore.NewKeyStore("./goethereumbook/EthereumAccount/tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	// 读取本地文件，储存格式是json，也可以读取后用json库进行序列化之类的
	jsonByte, _ := ioutil.ReadFile(file)
	password := "StarYuhen"
	// 导入密钥
	account, _ := ks.Import(jsonByte, password, password)
	logrus.Info("address:", account.Address.Hex())
}
