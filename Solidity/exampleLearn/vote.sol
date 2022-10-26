// SPDX-License-Identifier: MIT
pragma solidity >=0.5 <=0.9;

// 声明一个投票类 abstract--用于标识不被部署
 contract Ballot {
    // 选民格式
    struct Voter {
        uint weight;  // 权重
        bool voted; // 是否投票
        address delegate; // 委托地址
        uint vote;// 投票的索引
    }

    // 提交票格式 提案
    struct Proposal {
        bytes32 name; // 投票名字
        uint voteCount; // 投票数总和
    }

    // 公开的地址变量
    address public chairperson;

    // 状态转换
    mapping(address => Voter) public voters;

    // 申明一个提交票的数组记录
    Proposal[] public proposal;

    /// 每个投票都创建一个新的提案
    /* 
    结构、数组或映射类型的所有变量的显式数据位置现在是强制性的。
    这也适用于函数参数和返回变量。例如，将 uint[] x = m_x 更改为 uint[] storage x = m_x，
    将函数 f(uint[][] x) 更改为函数 f(uint[][] memory x) 其中 memory 是数据位置，可能相应地被 storage 或 calldata 替换。请注意，外部函数需要具有 calldata 数据位置的参数。
    */
     constructor(bytes32[] memory proposalName)  {
        chairperson = msg.sender;
        voters[chairperson].weight = 1;
        // 每一个新提交的提案，都需要添加进一个新的提交票数记录数组
        for (uint i = 0; i < proposalName.length; i++) {
            proposal.push(Proposal({
            name : proposalName[i],
            voteCount : 1
            }));
        }
    }

    // 创建一个投票方法 
    // 授权用户进行投票
//     function givRightToVote(address voter) public {
//         // require 检测是否错误，现在已经是最新版的了 https://qa.1r1g.cn/ethereum/ask/5499371/ try catch
//         require(
           
//             "Only chairperson can give right to vote."
//         );

//         try msg.sender=chairperson{
//             require(!);
//         }catch{
//    "Only chairperson can give right to vote."
//         }


//         voter[Voter].weight=1;
//     }


}