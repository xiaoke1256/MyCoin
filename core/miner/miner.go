package miner

import (
	"time"

	"github.com/shopspring/decimal"

	"xiaoke1256.com/mycoin/model"

	"math/rand"

	"xiaoke1256.com/mycoin/crypt"

	"xiaoke1256.com/mycoin/db"

	"fmt"
)

func Mine() {

	//挖创世区块
	//构造block
	newBlock := model.CoreBlock{}

	//构造交易(这是个coinbase交易)
	t1 := model.CoreTransaction{}
	t1.Version = "1.0"
	t1.InputCounter = 0
	t1.OutputCounter = 5

	outputs := make([]model.CoreOutput, 5, 5)
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
	head.ParentHeadHash = crypt.DoubleSha256([]byte{'G', 'E', 'N', 'E', 'S', 'I', 'S'}) //创世区块的父区块是个默认值
	head.TransactionsMerkleRoot = newBlock.GetTransactionMerkleRoot()
	head.Timestamp = time.Now()
	head.Target = [4]byte{0x00, 0x00, 0xFF, 0xFF}

	newBlock.Blockheader = head

	//挖吖挖吖挖
	for true {
		rand.Seed(time.Now().UnixNano())
		randNum1 := rand.Intn(2 ^ 16)
		randNum2 := rand.Intn(2 ^ 16)
		randNum3 := rand.Intn(2 ^ 16)
		randNum4 := rand.Intn(2 ^ 16)
		var bytes4 [4]byte = [4]byte{byte(randNum1), byte(randNum2), byte(randNum3), byte(randNum4)}
		head.Nonce = bytes4
		//校验是否满足target
		hashedBytes := crypt.Md5(head.ToBytes())
		var hashedBytes4 [4]byte
		copy(hashedBytes4[:4], hashedBytes[0:4])
		fmt.Printf("hashedBytes4: %x", hashedBytes4)
		if compareBytes(hashedBytes4, head.Target) < 0 {
			//满足需求
			break
		}
	}
	//挖出来了就保存区块
	db.Insert("block", newBlock)

}

/*
两个byte数组比大小
*/
func compareBytes(bytes1 [4]byte, bytes2 [4]byte) int {
	for i := 0; i < 4; i++ {
		if bytes1[i] < bytes2[i] {
			fmt.Println("-1")
			return -1
		} else if bytes1[i] > bytes2[i] {
			fmt.Println("1")
			return 1
		} else {
			continue
		}
	}
	return 0
}
