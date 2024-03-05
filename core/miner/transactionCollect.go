package miner

import (
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
