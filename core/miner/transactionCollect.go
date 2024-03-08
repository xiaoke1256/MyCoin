package miner

import (
	"github.com/shopspring/decimal"
	"xiaoke1256.com/mycoin/db"
	"xiaoke1256.com/mycoin/model"
)

/**
 * 收集交易（即从本地读取交易）
 * n:查询多少条数据
 */
func CollectTransaction(n int64) []model.CoreTransaction {
	//TODO 过滤出没有被本矿机处理掉的，也没有出现在先前的区块中的
	var ts []model.CoreTransaction = db.Search("transaction", n, model.CoreTransaction{})
	return ts
}

// 以下数据用于测试
func InsertATransaction() {
	t := model.CoreTransaction{Version: "0.1", InputCounter: 1, OutputCounter: 1}
	input := model.CoreInput{TXHash: [32]byte{0x1E, 0x00, 0xFF, 0xFF,
		0x1E, 0x0F, 0xFF, 0xFF,
		0x1E, 0xEA, 0xFF, 0xFF,
		0x1E, 0x6E, 0xFF, 0xFF,
		0x1E, 0x00, 0xFF, 0xFF,
		0x1E, 0x00, 0xFF, 0xFF,
		0x1E, 0x00, 0xFF, 0xFF,
		0x1E, 0x00, 0xFF, 0xFF}, UTXOIdx: 1, UnlockScriptSize: 18, UnlockScript: "23221324342112321332"}
	t.Inputs = []model.CoreInput{input}
	output := model.CoreOutput{Amt: decimal.NewFromInt32(int32(2324435)), LockScriptSize: 18, LockScript: "3242311232213443542"}
	t.Outputs = []model.CoreOutput{output}
	db.Insert("transaction", t)
}
