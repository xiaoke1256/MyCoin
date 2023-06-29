package miner

import (
	"time"

	"github.com/shopspring/decimal"

	"xiaoke1256.com/mycoin/model"
)

func mine() {

	//挖创世区块
	//构造block
	newBlock := model.CoreBlock{}

	//构造交易
	t1 := model.CoreTransaction{}
	t1.Version = "1.0"
	t1.InputCounter = 0
	t1.OutputCounter = 5

	outputs := []model.CoreOutput{}
	for i := 0; i < 5; i++ {
		output := model.CoreOutput{}
		amt, err := decimal.NewFromString("50000000000000000000")
		if err != nil {
			panic(err)
		}
		output.Amt = amt
		output.LockScriptSize = uint32(len("ssss"))
		output.LockScript = "ssss"
		outputs[i] = output
	}
	t1.Outputs = outputs
	t1.LockTime = time.Now()

	newBlock.TransactionCounter = 1
	ts := []model.CoreTransaction{t1}
	newBlock.Transactions = ts

	//构造head
	head := model.CoreBlockheader{}
	head.Version = "1.0"
	head.ParentHeadHash = []byte{'G', 'E', 'N', 'E', 'S', 'I', 'S'} //创世区块的父区块是个默认值
	//head.TransactionsMerkleRoot =
	head.Timestamp = time.Now()
	head.Target = []byte{0xFF, 0xFF, 0x00, 0x00}

}
