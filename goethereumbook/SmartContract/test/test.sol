// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

// node.js 可以用这个 部署智能合约 https://kauri.io/#communities/Getting%20started%20with%20dapp%20development/truffle--smart-contract-compilation-and-deploymen/#solc-compiler

// 因为我是Golang开发，所以需要安装Solidity编译合约为abi，Abigen则是编译为Go文件

// solc --abi --bin test.sol  /  solc --abi --bin -o ./  test.sol
// abigen --abi=Store.abi --pkg=store --out=Store.go  生成Golang文件
// 结合合约 abigen --bin=Store.bin --abi=Store.abi --pkg=store --out=Store.go

contract Store {
    event ItemSet(bytes32 key, bytes32 value);

    string public version;
    mapping(bytes32 => bytes32) public items;

    constructor(string memory _version) {
        version = _version;
    }

    function setItem(bytes32 key, bytes32 value) external {
        items[key] = value;
        emit ItemSet(key, value);
    }
}