package model

import (
	"time"

	"strconv"

	"github.com/shopspring/decimal"
)

type CoreBlock struct {
	Blocksize          uint32
	Blockheader        CoreBlockheader
	TransactionCounter uint32            //交易计数器（本包中有几个交易）
	Transactions       []CoreTransaction //交易
}

type CoreBlockheader struct {
	/*4	版本	区块版本号，表示本区块遵守的验证规则*/
	Version string
	/*32	父区块头哈希值	前一区块的哈希值，使用SHA256(SHA256(父区块头))计算*/
	ParentHeadHash []byte
	/*32	Merkle根	该区块中交易的Merkle树根的哈希值，同样采用SHA256(SHA256())计算*/
	TransactionsMerkleRoot []byte
	/* 4	时间戳	该区块产生的近似时间，精确到秒的UNIX时间戳，必须严格大于前11个区块时间的中值，
	同时全节点也会拒绝那些超出自己2个小时时间戳的区块*/
	Timestamp time.Time
	/* 4	难度目标	该区块工作量证明算法的难度目标，已经使用特定算法编码*/
	Target []byte
	/* 4	Nonce	为了找到满足难度目标所设定的随机数，
	为了解决32位随机数在算力飞升的情况下不够用的问题，规定时间戳和coinbase交易信息均可更改，以此扩展nonce的位数*/
	Nonce []byte
}

type CoreTransaction struct {
	/*4	交易版本号	明确这笔交易参照的规则*/
	Version string
	/* 1-9	输入计数器	包含的交易输入数量*/
	InputCounter uint32
	/*不定	交易输入	一个或多个交易输入*/
	Inputs []CoreInput
	/* 1-9	输出计数器	包含的交易输出数量*/
	OutputCounter uint32
	/*不定	交易输出	一个或多个交易输出*/
	Outputs []CoreOutput
	/*4	锁定时间	一个区块号或UNIX时间戳*/
	LockTime time.Time
}

type CoreInput struct {
	/*32	交易哈希值	指向被花费的UTXO所在的交易的哈希*/
	TXHash []byte
	/*4	输出索引	被花费的UTXO的索引号，第一个是0*/
	UTXOIdx uint16
	/*1-9	解锁脚本大小	用字节表示的后面的解锁脚本长度*/
	UnlockScriptSize uint32
	/*不定	解锁脚本	满足UTXO解锁脚本条件的脚本 (一般情况下就是个密钥)*/
	UnlockScript string
	/*4	序列号	固定值0xFFFFFFFF （用于标识脚本结束）*/
}

type CoreOutput struct {
	/* 8	总量	用聪表示的比特币值*/
	Amt decimal.Decimal
	/* 1-9	锁定脚本大小	用字节表示的后面的锁定脚本长度（可以是个密码）*/
	LockScriptSize uint32
	/* 不定	锁定脚本	一个定义了支付输出所需条件的脚本 */
	LockScript string
}

/*
把一个交易转成byte数组
*/
func (t CoreTransaction) ToBytes() []byte {
	return []byte{}
}

/*
把一个交易转成 json
*/
func (t CoreTransaction) ToJson() string {
	var s = ""
	s += "{"
	s += "version:"
	s += "'" + t.Version + "'"
	s += ","
	s += "inputCounter:" + strconv.FormatUint(uint64(t.InputCounter), 10)
	s += ","
	s += "inputs:"
	s += "["
	for i, input := range t.Inputs {
		if i > 0 {
			s += ","
		}
		s += input.ToJson()
	}
	s += "]"
	s += ","
	s += "outputCounter:"+strconv.FormatUint(uint64(t.OnputCounter), 10)
	s += "}"
	return s
}

func (input CoreInput) ToJson() string {
	var s = "{"
	s += "TXHash:" + string(input.TXHash)
	s += ","
	s += "UTXOIdx:" + strconv.FormatUint(uint64(input.UTXOIdx), 10)
	s += ","
	s += "unlockScriptSize:" + strconv.FormatUint(uint64(input.UnlockScriptSize), 10)
	s += ","
	s += "unlockScript:" + input.UnlockScript
	s += "}"
	return s
}
